package product

import "github.com/gorilla/mux"

func (h handler) Register(r *mux.Router) {
	r.HandleFunc("/", h.List).Methods("GET")
	r.HandleFunc("/products", h.List).Methods("GET")
	r.HandleFunc("/products/new", h.AddForm).Methods("GET")
	r.HandleFunc("/products/new", h.Add).Methods("POST")
	r.HandleFunc("/products/{id}", h.Edit).Methods("GET")
	r.HandleFunc("/products/{id}", h.Update).Methods("POST")
	r.HandleFunc("/products/{id}", h.Delete).Methods("DELETE")
}
