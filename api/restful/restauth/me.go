package restauth

import (
	"encoding/json"
	"net/http"

	"github.com/MrSpoony/grade-tracker/api/cookie"
)

func (h *Handler) me(w http.ResponseWriter, r *http.Request) {
	claims, err := cookie.GetClaims(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data, err := json.Marshal(claims.Data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(data)

}
