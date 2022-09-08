package restauth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/MrSpoony/grade-tracker/backend/logic/user"
)

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var user user.User
	// Parse JSON input
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}

	// Hash password
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	user.Password = string(pwd)

	// Store user
	err = h.srv.DB.StoreNewUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}
