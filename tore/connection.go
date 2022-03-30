package main

import(
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	
)

//getConnection to bd
func getConnection() *sql.DB{
	dsn:="postgres://postgres:12345@localhost/postgres?sslmode=disable"
	db,err:=sql.Open("postgres",dsn)
	if err!=nil{
		log.Fatal(err)
	}
	return db
}
//clear regista

//clear regista
