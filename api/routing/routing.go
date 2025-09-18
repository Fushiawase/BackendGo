package routing

import (
	"BackendGo/core/document"
	"BackendGo/core/errors"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type server struct {
	docRepo document.Repository
}

func NewServer(docRepo document.Repository) server {
	return server{docRepo}
}

type decodedDocument struct {
	Name        string
	Description string
}

func (s *server) SetUpRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /documents", s.handleCreate)
	mux.HandleFunc("GET /documents/{Id}", s.handleFindById)
	mux.HandleFunc("DELETE /documents/{Id}", s.handleDelete)
	return mux
}

func (s *server) handleCreate(w http.ResponseWriter, r *http.Request) {
	var decodedDoc decodedDocument
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&decodedDoc); err != nil {
		http.Error(w, "Malformed Json", 400)
		return
	}
	doc := &document.Document{Name: decodedDoc.Name, Description: decodedDoc.Description}
	if err := s.docRepo.Create(doc); err != nil {
		http.Error(w, "Database Error", 500)
		return
	}
	fmt.Fprintf(w, "Successfully created document of Id : %d", doc.Id)
}

func (s *server) handleFindById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("Id")
	id, err := strconv.Atoi(idString)
	if err != nil || id < 0 {
		http.Error(w, "Id must be an integer and >= 0", 400)
		return
	}
	d, err := s.docRepo.FindByID(id)
	if err != nil {
		if rnfErr, ok := err.(errors.RecordNotFoundErr); ok {
			http.Error(w, rnfErr.Error(), 404)
			return
		}
		http.Error(w, "Database Error", 500)
		return
	}
	fmt.Fprintf(w, "Found document of Id: %d with Name: %s and Description: %s", d.Id, d.Name, d.Description)
}

func (s *server) handleDelete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("Id")
	id, err := strconv.Atoi(idString)
	if err != nil || id < 0 {
		http.Error(w, "Id must be an integer and >= 0", 400)
		return
	}
	if err := s.docRepo.Delete(id); err != nil {
		if rnfErr, ok := err.(errors.RecordNotFoundErr); ok {
			http.Error(w, rnfErr.Error(), 404)
			return
		}
		http.Error(w, "Database Error", 500)
		return
	}
	fmt.Fprintf(w, "Successfully deleted document of Id: %d", id)
}
