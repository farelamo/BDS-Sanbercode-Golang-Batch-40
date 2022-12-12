package config

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

const (
	host 		= "localhost"
	port 		= 5432
	user 		= "postgres"
	password 	= "root"
	dbname   	= "quiz3" 
)

var (
	Db *sql.DB
	Err error
)

func init(){
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	Db, Err = sql.Open("postgres", psqlInfo)
	if Err != nil {
		panic(Err)
	}
	// defer Db.Close()

	Err = Db.Ping()
	if Err != nil {
		panic(Err)
	}

	fmt.Println("Successfully connected to database", dbname)
}