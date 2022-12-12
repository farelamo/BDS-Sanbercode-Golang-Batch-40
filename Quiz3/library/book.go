package library

import (
	"Quiz3/config"
)

type Book struct {
	Id 			int 	`json:"id"`
	Title 		string 	`json:"title"`
	Description	string	`json:"desc"`
	ImageUrl	string	`json:"image_url"`
	ReleaseYear	int		`json:"release_year"`
	Price		string	`json:"price"`
	TotalPage	int		`json:"total_page"`
}

func CreateBuku(b Book) Book {
	var book = Book{}

	sqlStatement := `
	INSERT INTO books (title, description, image_url, release_year, price, total_page)
	VALUES ($1, $2, $3, $4, $5, $6)
	Returning *
	`

	config.Err = config.Db.QueryRow(
		sqlStatement, b.Title, b.Description, b.ImageUrl, 
		b.ReleaseYear, b.Price, b.TotalPage, 
	).Scan(
		&book.Id, &book.Title, &book.Description, &book.ImageUrl,
		&book.ReleaseYear, &book.Price, &book.TotalPage,
	)

	if config.Err != nil {
		panic(config.Err)
	}

	return book
}