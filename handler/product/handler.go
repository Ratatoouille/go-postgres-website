package product

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"postgres/model"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type handler struct {
	DB   *sql.DB
	Tmpl *template.Template
}

func NewHandler(db *sql.DB, tmpl *template.Template) *handler {
	return &handler{DB: db, Tmpl: tmpl}
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	products := []*model.Product{}

	rows, err := h.DB.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		product := &model.Product{}

		err = rows.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
		if err != nil {
			log.Println(err)
		}

		products = append(products, product)
	}
	rows.Close()

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

func (h *handler) AddForm(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) Add(w http.ResponseWriter, r *http.Request) {
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		log.Println(err)
	}

	_, err = h.DB.Exec(
		"INSERT INTO products (model, company, price) VALUES ($1, $2, $3)",
		r.FormValue("model"),
		r.FormValue("company"),
		price,
	)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *handler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	post := &model.Product{}

	row := h.DB.QueryRow("SELECT id, model, company, price  FROM products WHERE id = $1", id)

	err = row.Scan(&post.Id, &post.Model, &post.Company, &post.Price)
	if err != nil {
		log.Println("scan", err)
	}

	err = h.Tmpl.ExecuteTemplate(w, "edit.html", post)
	if err != nil {
		http.Error(w, "template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	// в целям упрощения примера пропущена валидация
	_, err = h.DB.Exec(
		"UPDATE products SET model = $1, company = $2, price = $3 WHERE id = $4",
		r.FormValue("model"),
		r.FormValue("company"),
		r.FormValue("price"),
		id,
	)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	_, err = h.DB.Exec(
		"DELETE FROM products WHERE id = $1",
		id,
	)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusOK)
}
