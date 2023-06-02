package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

func main() {
	database, _ := sql.Open("sqlite3", "./test.db")
	modif, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, NAME VARCHAR)")
	modif.Exec()
	modif, _ = database.Prepare("INSERT INTO users(name) VALUES (?)")
	modif.Exec("test")
	rows, _ := database.Query("SELECT id, name From users")
	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(strconv.Itoa(id) + ": " + name)
	}
}
