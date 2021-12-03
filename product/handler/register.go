package handler

import (
	"html/template"

	"github.com/Ratatoouille/product"
	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoints(r *mux.Router, u product.UseCase, tmpl *template.Template) {
	h := NewHandler(u, tmpl)

	r.HandleFunc("/", h.List).Methods("GET")
	r.HandleFunc("/products", h.List).Methods("GET")
	r.HandleFunc("/products/new", h.AddForm).Methods("GET")
	r.HandleFunc("/products/new", h.Add).Methods("POST")
	r.HandleFunc("/products/{id}", h.Edit).Methods("GET")
	r.HandleFunc("/products/{id}", h.Update).Methods("POST")
	r.HandleFunc("/products/{id}", h.Delete).Methods("DELETE")
}
