package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./aquarium.db")
	// database, _ := sql.Open("sqlite3", "./nraboy.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS fish (name TEXT, species TEXT, tank_number INTEGER)")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("INSERT INTO fish (name, species, tank_number) VALUES (?, ?, ?)", "Sammy", "shark", 1)
	if err != nil {
		fmt.Println(err)
		return

	}
	rows, err := db.Query("SELECT name, species, tank_number FROM fish")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	var name string
	var species string
	var tankNumber int
	for rows.Next() {
		err = rows.Scan(&name, &species, &tankNumber)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name: %s, Species: %s, Tank Number: %d\n", name, species, tankNumber)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
