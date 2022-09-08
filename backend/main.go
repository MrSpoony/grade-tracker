package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/MrSpoony/grade-tracker/backend/db"
	"github.com/MrSpoony/grade-tracker/backend/restful/restauth"
	"github.com/MrSpoony/grade-tracker/backend/server"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func main() {

	// rows, err := db.Query("SELECT id, username, email FROM tabUser; -- sql")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer rows.Close()
	// users := make([]User, 0)
	// for rows.Next() i
	// 	var user User
	// 	if err = rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
	// 		panic(err.Error())
	// 	}
	// 	users = append(users, user)
	// }
	// fmt.Println(users)

	mysqlDB, err := NewDB()
	if err != nil {
		panic(err.Error())
	}
	r := mux.NewRouter()
	db := db.New(mysqlDB)
	srv := server.New(db, r)

	auth := restauth.NewHandler(srv)
	auth.Handle()

	if err = srv.Run(); err != nil {
		panic(err.Error())
	}
}

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:thisisaverysecuremysqlpassword@/gradetracker")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}