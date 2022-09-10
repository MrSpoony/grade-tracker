package restauth

import "net/http"

func (h *Handler) Handle() {
	h.Router.Handle("/api/login", http.HandlerFunc(h.login)).Methods("POST")
	h.Router.Handle("/api/signup", http.HandlerFunc(h.signup)).Methods("POST")
	h.Router.Handle("/api/refresh", http.HandlerFunc(h.refresh))
}
