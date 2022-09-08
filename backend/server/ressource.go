package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MrSpoony/grade-tracker/backend/db"
)

type Server struct {
	DB     *db.DB
	Router *mux.Router
}

func New(db *db.DB, router *mux.Router) *Server {
	return &Server{db, router}
}

func (s *Server) Run() error {
	defer s.DB.Close()
	return http.ListenAndServe(":6969", s.Router)
}
