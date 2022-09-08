package restauth

func (h *Handler) Handle() {
	h.srv.Router.HandleFunc("/api/login", h.login)
	h.srv.Router.HandleFunc("/api/signup", h.signup)
	h.srv.Router.HandleFunc("/api/refresh", h.refresh)
}
