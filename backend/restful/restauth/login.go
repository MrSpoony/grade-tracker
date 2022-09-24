package restauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/MrSpoony/grade-tracker/backend/cookie"
	"github.com/MrSpoony/grade-tracker/backend/logic/user"
)

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var creds credentials
	// Parse JSON input
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\"}", "Could not decode JSON")
		return
	}

	// Get correct user
	usr, err := user.GetUserByUsername(h.DB, creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}

	// Hash password
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(creds.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("{\"message\": \"Wrong password\"}"))
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &cookie.Claims{
		Data: cookie.Data{
			User: *usr,
		},
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(cookie.JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	out, err := json.Marshal(cookie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(out)
}
