package library

import (
	"Quiz3/config"
	_ "fmt"
	_ "log"
	"time"
)

type Category struct {
	Id 			int 		`json:"id"`
	Name 		string 		`json:"name"`
	CreatedAt 	time.Time	`json:"created_at,omitempty"`
	UpdatedAt   time.Time	`json:"updated_at,omitempty"`
}

func GetAllKategori() []Category {
	var category = []Category{}

	sqlStatement := `SELECT * FROM categories`
	rows, err 	 :=  config.Db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	
	for rows.Next() {
		var c = Category{}

		err = rows.Scan(&c.Id, &c.Name, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			panic(err)
		}

		category = append(category, c)
	}

	return category
}

func CreateKategori(c Category) Category{
	var category = Category{}
	// log.Println(c.Name) //cek isi name

	sqlStatement := `
	INSERT INTO categories (name, created_at, updated_at)
	VALUES ($1, $2, $3)
	Returning *
	`

	config.Err = config.Db.QueryRow(sqlStatement, c.Name, time.Now(), time.Now()).
		Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
			
	if config.Err != nil {
		panic(config.Err)
	}

	return category
}

func UpdateKategori(c Category) {
	sqlStatement := `
	UPDATE categories
	SET name = $2, updated_at = $3
	WHERE id = $1;
	`

	if _, err := config.Db.Exec(sqlStatement, c.Id, c.Name, time.Now()); err != nil {
		panic(err)
	}
}

func DeleteKategori(c Category){
	sqlStatement := `
	DELETE FROM categories WHERE id = $1
	`

	if _, err := config.Db.Exec(sqlStatement, c.Id); err != nil {
		panic(err)
	}
}