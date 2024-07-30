package customJWT

import (
	"fmt"
	customerError "learnGin/src/common/customError"
	envconfig "learnGin/src/common/envConfig"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	SignJWT(c *gin.Context, _id string, name string, email string) SignJwtResp
	VerifyJwt(c *gin.Context, jwtStr string) VerifyJwtResp
}

func SignJWT(c *gin.Context, _id string, name string, email string) SignJwtResp {

	tkExipred, err := strconv.Atoi(envconfig.GetEnv("ACCESS_TOKEN_EXPIRED"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	rfExpired, err := strconv.Atoi(envconfig.GetEnv("REQUEST_TOKEN_EXPIRED"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":   _id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Second * time.Duration(tkExipred)).Unix(),
	})

	rfTk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":   _id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Second * time.Duration(rfExpired)).Unix(),
	})

	jwtSecret := envconfig.GetEnv("JWT_SECRET")

	tkStr, err := tk.SignedString([]byte(jwtSecret))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Sign token error"})
	}

	rfTkStr, err := rfTk.SignedString([]byte(jwtSecret))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Sign refresh token error"})
	}

	resp := SignJwtResp{
		AccessToken:  tkStr,
		RefreshToken: rfTkStr,
	}

	return resp
}

func VerifyJwt(c *gin.Context, jwtStr string) VerifyJwtResp {
	token, _ := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(envconfig.GetEnv("JWT_SECRET")), nil
	})

	var verified VerifyJwtResp

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Chec k the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, customerError.ResponseError(
				customerError.UNAUTHORIZE, "Unauthorized",
			))
		}

		_id := claims["_id"].(string)
		verified = VerifyJwtResp{
			Id:    _id,
			Name:  claims["name"].(string),
			Email: claims["email"].(string),
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, customerError.ResponseError(
			customerError.UNAUTHORIZE, "Unauthorized",
		))
	}

	return verified
}
