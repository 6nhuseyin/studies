package main

import "database/sql"

func main() {
	database, _ := sql.Open("sqlite3", "./nraboy.db")
}
