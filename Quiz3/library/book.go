package library

import (
	"Quiz3/config"
	"time"
)

type Book struct {
	Id 			int 		`json:"id"`
	Title 		string 		`json:"title"`
	Description	string		`json:"desc"`
	ImageUrl	string		`json:"image_url"`
	ReleaseYear	int			`json:"release_year"`
	Price		string		`json:"price"`
	TotalPage	int			`json:"total_page"`
	Thickness	string  	`json:"thickness"`
	CreatedAt 	time.Time	`json:"created_at,omitempty"`
	UpdatedAt   time.Time	`json:"updated_at,omitempty"`
}

func GetAllBuku() []Book {
	var books = []Book{}

	sqlStatement := `SELECT * FROM books`
	rows, err    := config.Db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var b = Book{}

		err = rows.Scan(&b.Id, &b.Title, &b.Description, &b.ImageUrl, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CreatedAt, &b.UpdatedAt)
		if err != nil {
			panic(err)
		}

		books = append(books, b)
	}

	return books
}

func CreateBuku(b Book) Book {
	var book = Book{}

	sqlStatement := `
	INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	Returning *
	`

	config.Err = config.Db.QueryRow(
		sqlStatement, b.Title, b.Description, b.ImageUrl, 
		b.ReleaseYear, b.Price, b.TotalPage, b.Thickness,
		time.Now(), time.Now(),
	).Scan(
		&book.Id, &book.Title, &book.Description, &book.ImageUrl,
		&book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness,
		&book.CreatedAt, &book.UpdatedAt,
	)

	if config.Err != nil {
		panic(config.Err)
	}

	return book
}

func UpdateBuku(b Book) {
	sqlStatement := `
	UPDATE books
	SET title = $2, description = $3, image_url = $4, release_year = $5, price = $6, total_page = $7, thickness = $8, updated_at = $9
	WHERE id = $1
	`

	if _, err := config.Db.Exec(
		sqlStatement, b.Id, b.Title, b.Description, b.ImageUrl, 
		b.ReleaseYear, b.Price, b.TotalPage, b.Thickness, time.Now(),
	); err != nil {
		panic(err)
	}
}

func DeleteBuku(b Book){
	sqlStatement := `
	DELETE FROM books where id = $1
	`
	
	if _, err := config.Db.Exec(sqlStatement, b.Id); err != nil {
		panic(err)
	}
}