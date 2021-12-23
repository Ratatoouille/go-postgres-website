package server

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ratatoouille/product"
	"github.com/Ratatoouille/product/handler"
	"github.com/Ratatoouille/product/repository"
	"github.com/Ratatoouille/product/usecase"
	"github.com/jackc/pgx/v4"

	"github.com/gorilla/mux"
)

type App struct {
	httpServer *http.Server

	productUC product.UseCase
}

func NewApp() *App {
	db := initDB()

	productRepo := repository.NewProductRepository(db)

	return &App{
		productUC: usecase.NewProductUseCase(productRepo),
	}
}

func initDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:mypass@localhost:5432/product")
	if err != nil {
		log.Panicln(err)
		os.Exit(1)
	}

	return conn
}

func (a *App) Run(port string) error {
	r := mux.NewRouter()

	handler.RegisterHTTPEndpoints(
		r, a.productUC,
		template.Must(template.ParseGlob("./web/*")),
	)

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
