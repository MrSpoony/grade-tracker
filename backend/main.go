package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/MrSpoony/grade-tracker/backend/db"
	"github.com/MrSpoony/grade-tracker/backend/restful/restauth"

	//"github.com/MrSpoony/grade-tracker/backend/restful/restclass"
	"github.com/MrSpoony/grade-tracker/backend/restful/restsubject"
	"github.com/MrSpoony/grade-tracker/backend/server"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func main() {
	mysqlDB, err := NewDB()
	if err != nil {
		panic(err.Error())
	}
	r := mux.NewRouter().StrictSlash(true)
	db := db.New(mysqlDB)
	srv := server.New(db, r)

	restauth.NewHandler(srv)
	restsubject.NewHandler(srv)

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
