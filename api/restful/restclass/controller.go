package restclass

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/MrSpoony/grade-tracker/api/logic/class"
)

func (h *Handler) Handle() {
	h.Router.Handle("/api/class", http.HandlerFunc(h.add)).Methods("POST")
	h.Router.Handle("/api/class", http.HandlerFunc(h.getAll)).Methods("GET")
	h.Router.Handle("/api/class/{id}", http.HandlerFunc(h.get)).Methods("GET")
	h.Router.Handle("/api/class/{id}", http.HandlerFunc(h.update)).Methods("PUT")
	h.Router.Handle("/api/class/{id}", http.HandlerFunc(h.delete)).Methods("DELETE")
}

func (h *Handler) add(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	classes, err := class.GetAllClasses(h.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	out, err := json.Marshal(classes)
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
	class, err := class.GetClassByID(h.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	out, err := json.Marshal(class)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"message\": \"%s\" }", err.Error())
		return
	}
	w.Write(out)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {

}
