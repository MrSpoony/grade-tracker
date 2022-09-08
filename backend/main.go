package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func main() {
	db, err := sql.Open("mysql", "root:thisisaverysecuremysqlpassword@/gradetracker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT id, username, email FROM tabUser; -- sql")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	fmt.Println(users)
}
