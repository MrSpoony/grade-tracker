package restsubject

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/MrSpoony/grade-tracker/backend/logic/subject"
)

func (h *Handler) Handle() {
	h.Router.Handle("/api/subject", http.HandlerFunc(h.addSubject)).Methods("POST")
	h.Router.Handle("/api/subject/{id}", http.HandlerFunc(h.getSubject)).Methods("GET")
	h.Router.Handle("/api/subject/{id}", http.HandlerFunc(h.updateSubject)).Methods("PUT")
	h.Router.Handle("/api/subject/{id}", http.HandlerFunc(h.deleteSubject)).Methods("DELETE")
}

func (h *Handler) addSubject(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\":\"Id has to be an integer\"}"))
	}
	subject, err := subject.GetSubjectByID(h.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	out, err := json.Marshal(subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	w.Write(out)
}

func (h *Handler) updateSubject(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteSubject(w http.ResponseWriter, r *http.Request) {

}
