package main

import "database/sql"
import (
	_ "fortl.net/ch3/dbdriver/postgres"
	"log"
)

func main() {
	conn, err := sql.Open("aaa", "mydb")
	if err != nil {
		log.Println(err)
		return
	}

	println(conn)
}
