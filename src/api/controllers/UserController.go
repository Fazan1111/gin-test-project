package controllers

import (
	"fmt"
	userDto "learnGin/src/api/dto/user"
	"learnGin/src/api/services"
	respons "learnGin/src/common/response"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Users
// @Summary Get users
// @Description Gest list user
// @Schemes
// @Tags User
// @Accept  json
// @Produce  json
// @Success 201 {string} string "ok"
// @Router /api/users [get]
func GetUsers(c *gin.Context) {
	users := []userDto.ListUserDto{}

	fmt.Println(users)
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 10
	}

	fmt.Println("request header", c.Request.Header.Get("userName"))

	page, errPage := strconv.Atoi(c.Query("page"))
	if errPage != nil {
		page = 1
	}

	data := services.ListUsers()
	fmt.Println("fetch user data", data)

	c.JSON(http.StatusOK, respons.Paginate(data, limit, page))
}
