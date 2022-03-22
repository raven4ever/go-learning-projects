package main

var movies = []Movie{
	{
		ID:       "1",
		Title:    "The Shawshank Redemption",
		Year:     1994,
		Director: &Director{FirstName: "Frank", LastName: "Darabont"},
	},
	{
		ID:       "2",
		Title:    "The Godfather",
		Year:     1972,
		Director: &Director{FirstName: "Francis", LastName: "Ford Coppola"},
	},
}
