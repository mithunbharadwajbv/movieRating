package userService

import (
	"github.com/sanjeev/src/domain/movieRepo"
	"github.com/sanjeev/src/domain/userRepo"
)

func AddUser(user userRepo.User) error {
	err := userRepo.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByName(name string) (*userRepo.User, error) {
	user, err := userRepo.GetUserByName(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddMovieRating(userName string, rating userRepo.Rating) error {

	//get logged in user
	user, err := userRepo.GetUserByName(userName)
	if err != nil {
		return err
	}

	//add rating to user object and save
	user.Rating = append(user.Rating, rating)
	err = userRepo.UpdateUser(*user)
	if err != nil {
		return err
	}

	//add comment to overall movie rating too
	{
		movie, err := movieRepo.GetMovieByName(rating.MovieName)
		if err != nil {
			return err
		}
		movie.Comments = append(movie.Comments, rating.Comment)
		movie.Rating = ((float64(movie.NoOfPeopleRated) * movie.Rating) + float64(rating.Rating)) / float64(movie.NoOfPeopleRated+1)
		movie.NoOfPeopleRated = movie.NoOfPeopleRated + 1
		err = movieRepo.UpdateMovie(*movie)
		if err != nil {
			return err
		}
	}

	return nil
}
