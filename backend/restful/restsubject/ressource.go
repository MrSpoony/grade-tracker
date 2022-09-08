package restsubject

import (
	"github.com/gorilla/mux"

	"github.com/MrSpoony/grade-tracker/backend/db"
	"github.com/MrSpoony/grade-tracker/backend/server"
)

// Handler is a restful interface for business partner related endpoints.
type Handler struct {
	DB     *db.DB
	Router *mux.Router
}

// NewHandler creates a new handler for auth
func NewHandler(srv *server.Server) (h *Handler) {
	h = &Handler{srv.DB, srv.Router}
	return
}
