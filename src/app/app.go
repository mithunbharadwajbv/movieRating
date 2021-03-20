package app

import (
	"github.com/sanjeev/src/controller/movieController"
	"github.com/sanjeev/src/controller/userController"

	"github.com/gin-gonic/gin"
	"github.com/sanjeev/src/config"
)

var (
	router           = gin.Default()
	applicattionPort = config.Conf.ApplicattionPort
)

// StartApplication starts gin application
func StartApplication() {

	router.POST("/AddMovie", movieController.AddMovie)
	router.GET("/GetMovieByName/:name", movieController.GetMovieByName)
	router.POST("/AddMovieRating", movieController.AddMovieRating)

	router.POST("/AddUser", userController.AddUser)
	router.GET("/GetUserByName/:name", userController.GetUserByName)

	router.Run(applicattionPort)
}
