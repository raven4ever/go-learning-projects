package data

import (
	"github.com/raven4ever/movies-crud-rest-api/pkg/model"
)

var Movies = []model.Movie{
	{
		ID:       "1",
		Title:    "The Shawshank Redemption",
		Year:     1994,
		Director: &model.Director{FirstName: "Frank", LastName: "Darabont"},
	},
	{
		ID:       "2",
		Title:    "The Godfather",
		Year:     1972,
		Director: &model.Director{FirstName: "Francis", LastName: "Ford Coppola"},
	},
}
