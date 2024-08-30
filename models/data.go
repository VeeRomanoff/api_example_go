package models

var DataBase []Book

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthYear int    `json:"birth_year"`
}

func init() {
	book := Book{
		ID:    1,
		Title: "Star Wars",
		Author: Author{
			FirstName: "Lucas",
			LastName:  "Films",
			BirthYear: 1968,
		},
		YearPublished: 2000,
	}
	book2 := Book{
		ID:    2,
		Title: "Lord of Rings",
		Author: Author{
			FirstName: "J.R",
			LastName:  "Tolklin",
			BirthYear: 1822,
		},
		YearPublished: 2000,
	}
	DataBase = append(DataBase, book, book2)
}
