package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Ratatoouille/model"
	"github.com/Ratatoouille/product"

	"github.com/gorilla/mux"
)

type Handler struct {
	useCase product.UseCase
	Tmpl    *template.Template
}

func NewHandler(useCase product.UseCase, tmpl *template.Template) *Handler {
	return &Handler{
		useCase: useCase,
		Tmpl:    tmpl,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.useCase.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.Tmpl.ExecuteTemplate(w, "index.html", struct {
		Products []*model.Product
	}{
		Products: products,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddForm(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		log.Println(err)
	}

	err = h.useCase.CreateProduct(&model.Product{
		Model:   r.FormValue("model"),
		Company: r.FormValue("company"),
		Price:   price,
	})
	if err != nil {
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	p := &model.Product{}

	err = h.useCase.EditProduct(p, id)
	if err != nil {
		return
	}

	err = h.Tmpl.ExecuteTemplate(w, "edit.html", p)
	if err != nil {
		http.Error(w, "template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		log.Println(err)
	}

	p := &model.Product{
		Id:      id,
		Model:   r.FormValue("model"),
		Company: r.FormValue("company"),
		Price:   price,
	}

	err = h.useCase.UpdateProduct(p)
	if err != nil {
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	err = h.useCase.DeleteProduct(id)
	if err != nil {
		return
	}

	http.Redirect(w, r, "/", http.StatusOK)
}
