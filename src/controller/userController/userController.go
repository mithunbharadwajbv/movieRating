package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanjeev/src/domain/userRepo"
	"github.com/sanjeev/src/service/userService"
)

func AddUser(c *gin.Context) {

	var user userRepo.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userService.AddUser(user)
	c.JSON(http.StatusOK, "added user succesfully")
}

func GetUserByName(c *gin.Context) {

	name := c.Params.ByName("name")
	user, err := userService.GetUserByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
