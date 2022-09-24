package restsubject

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/MrSpoony/grade-tracker/api/logic/subject"
)

func (h *Handler) Handle() {
	h.Router.Handle("/api/subject", http.HandlerFunc(h.add)).Methods("POST")
	h.Router.Handle("/api/subject", http.HandlerFunc(h.getAll)).Methods("GET")
	h.Router.Handle("/api/subject/{id}", http.HandlerFunc(h.get)).Methods("GET")
	h.Router.Handle("/api/subject/{id}", http.HandlerFunc(h.update)).Methods("PUT")
	h.Router.Handle("/api/subject/{id}", http.HandlerFunc(h.delete)).Methods("DELETE")
}

func (h *Handler) add(w http.ResponseWriter, r *http.Request) {
	var subjct *subject.Subject
	err := json.NewDecoder(r.Body).Decode(&subjct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := subject.CreateSubject(h.DB, *subjct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	subjct, err = subject.GetSubjectByID(h.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	out, err := json.Marshal(subjct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	w.Write(out)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	subjects, err := subject.GetAllSubjects(h.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	out, err := json.Marshal(subjects)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	w.Write(out)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\":\"Id has to be an integer\"}"))
		return
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

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	var subjct subject.Subject
	err := json.NewDecoder(r.Body).Decode(&subjct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	err = subject.UpdateSubject(h.DB, subjct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	out, err := json.Marshal(subjct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	w.Write(out)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\":\"Id has to be an integer\"}"))
		return
	}
	subjct, err := subject.GetSubjectByID(h.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	if subjct == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", "Subject does not exist")
	}
	err = subject.DeleteSubjectByID(h.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	out, err := json.Marshal(subjct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	w.Write(out)
}
