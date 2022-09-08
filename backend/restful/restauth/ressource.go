package restauth

import (
	"github.com/MrSpoony/grade-tracker/backend/server"
)

// Handler is a restful interface for business partner related endpoints.
type Handler struct {
	srv *server.Server
}

// NewHandler creates a new handler for auth
func NewHandler(srv *server.Server) (h *Handler) {
	h = &Handler{srv: srv}
	return
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
