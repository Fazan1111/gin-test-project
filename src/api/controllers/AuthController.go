package controllers

import (
	"fmt"
	authDto "learnGin/src/api/dto/auth"
	"learnGin/src/api/models"
	"learnGin/src/api/services"
	customerError "learnGin/src/common/customError"
	respons "learnGin/src/common/response"
	"learnGin/src/common/util"
	customJWT "learnGin/src/libs/jwt"
	mailOTP "learnGin/src/libs/mail"
	redisLib "learnGin/src/libs/redis"
	vonageAPI "learnGin/src/libs/vonage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth
// @Summary Register
// @Description Register new account
// @Schemes
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   register  body authDto.RegisterDto  true  "Register new account"
// @Success 201 {string} string "ok"
// @Router /api/auth/register [post]
func Registor(c *gin.Context) {
	var user models.User
	var input authDto.RegisterDto
	// Bind the request body to the RequestBody struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, customerError.ResponseError(customerError.INVALID_FIELD, err.Error()))
	}

	// // Check duplicate user
	// existUser := services.FindUserByName(c, user.Name)
	// if existUser != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User already exist"})
	// 	return
	// }

	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	now := time.Now().UTC()

	user.Name = input.Name
	user.Email = input.Email
	user.Password = hashedPassword
	user.CreatedAt = now
	user.UpdatedAt = now
	user.IsDeleted = false
	newUser := services.CreateUser(c, user)
	fmt.Println("new userId", newUser.Id.Hex())
	tk := customJWT.SignJWT(c, newUser.Id.Hex(), user.Name, user.Email)

	resp := authDto.RegistorResp{
		Id:           newUser.Id,
		Name:         user.Name,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		AccessToken:  tk.AccessToken,
		RefreshToken: tk.AccessToken,
	}

	c.JSON(http.StatusCreated, respons.ResponseSuccess(resp))
}

// Auth
// @Summary Login
// @Description Login to app
// @Schemes
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   register  body authDto.LoginDto  true  "Login"
// @Success 201 {object} authDto.RegistorResp
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var input authDto.LoginDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, customerError.ResponseError(customerError.INVALID_FIELD, err.Error()))
	}

	if input.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password cannot empty"})
		return
	}

	// Find user by email
	authUser := services.FindUserByMail(c, input.Email)

	// compare password
	isValid := util.ComparePassword(authUser.Password, input.Password)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid password"})
		return
	}

	tk := customJWT.SignJWT(c, authUser.Id.Hex(), authUser.Name, authUser.Email)

	resp := authDto.RegistorResp{
		Id:           authUser.Id,
		Name:         authUser.Name,
		Email:        authUser.Email,
		AccessToken:  tk.AccessToken,
		RefreshToken: tk.AccessToken,
	}

	c.JSON(http.StatusCreated, respons.ResponseSuccess(resp))
}

func VerifyMailOTP(c *gin.Context) {
	var input authDto.VerifyMailOTP

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, customerError.ResponseError(customerError.INVALID_FIELD, err.Error()))
		return
	}

	otp, err := mailOTP.SendMailOTP(input.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messasge": "OTP error"})
	}
	resp := map[string]interface{}{
		"otpCode": otp,
	}

	redisLib.SetVal("dd", otp)

	rdVal := redisLib.GetVal("dd")
	fmt.Println("redis value", rdVal)
	c.JSON(http.StatusOK, respons.ResponseSuccess(resp))
}

func VerifyPhoneOTP(c *gin.Context) {
	var input authDto.VerifyPhoneOTP
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, customerError.ResponseError(customerError.INVALID_FIELD, err.Error()))
		return
	}

	otp := vonageAPI.SendSMS(input.PhoneNumber)
	resp := map[string]interface{}{
		"otpCode": otp,
	}
	c.JSON(http.StatusOK, respons.ResponseSuccess(resp))
}
