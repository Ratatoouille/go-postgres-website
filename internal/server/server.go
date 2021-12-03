package server

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ratatoouille/handler/product"

	"github.com/gorilla/mux"
)

type App struct {
	httpServer *http.Server
	db         *sql.DB
}

func NewApp() *App {
	db := initDB()

	return &App{db: db}
}

func initDB() *sql.DB {
	connStr := "user=postgres password=mypass dbname=product sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panicln(err)
	}

	return db
}

func (a *App) CloseDb() {
	a.db.Close()
}

func (a *App) Run(port string) error {
	handlers := product.NewHandler(a.db, template.Must(template.ParseGlob("../web/*")))

	r := mux.NewRouter()

	handlers.Register(r)

	a.httpServer = &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		err := a.httpServer.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
