package movieService

import (
	"fmt"

	"github.com/sanjeev/src/domain/movieRepo"
	"github.com/sanjeev/src/domain/userRepo"
)

func AddMovie(userName string, movie movieRepo.Movie) error {

	//only admin user can add movie
	user, err := userRepo.GetUserByName(userName)
	if err != nil {
		return err
	}
	if user.Type != "admin" {
		return fmt.Errorf("only admin can add movies to Mongo")
	}

	err = movieRepo.AddMovie(movie)
	if err != nil {
		return err
	}
	return nil
}

func GetMovieByName(name string) (*movieRepo.Movie, error) {
	movie, err := movieRepo.GetMovieByName(name)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
