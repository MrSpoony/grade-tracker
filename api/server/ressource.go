package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MrSpoony/grade-tracker/api/db"
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
	s.Router.Use(recoverPanic, Headers)
	return http.ListenAndServe(":6969", s.Router)
}

func Headers(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Backend always sends JSON
		w.Header().Set("Content-Type", "application/json")
		// No CORS problems
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Origin, Set-Cookie, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Requested-With, Authorization, enctype")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization, Set-Cookie")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// // Security
		// w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// w.Header().Add("Expiw", "0")
		// w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
		// w.Header().Add("Pragma", "no-cache")
		// w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		// w.Header().Add("X-Xss-Protection", "1; mode=block")
		// w.Header().Add("X-Content-Type-Options", "nosniff")
		//
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
