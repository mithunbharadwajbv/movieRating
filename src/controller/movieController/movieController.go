package movieController

import (
	"github.com/sanjeev/src/domain/movieRepo"
	"github.com/sanjeev/src/domain/userRepo"
	"github.com/sanjeev/src/service/movieService"
	"github.com/sanjeev/src/service/userService"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMovie(c *gin.Context) {

	var movie movieRepo.Movie

	//get request body of mmovie
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//get logged in user by query param
	userName := c.Query("userName")

	//call movie service to add movies
	err = movieService.AddMovie(userName, movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "added movie succesfully")
}

func GetMovieByName(c *gin.Context) {
	name := c.Params.ByName("name")
	user, err := movieService.GetMovieByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func AddMovieRating(c *gin.Context) {
	//get logged in user
	userName := c.Query("userName")

	//get Rating
	var rating userRepo.Rating
	err := c.ShouldBindJSON(&rating)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//call userService to save movieRating
	err = userService.AddMovieRating(userName, rating)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "succesfully rated movie")
}
