package server

import (
	"fmt"
	"log"
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
	s.Router.Use(recoverPanic, setHeaders)
	return http.ListenAndServe(":6969", s.Router)
}

func setHeaders(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func recoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Print(fmt.Errorf("panic: %+v", r))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
