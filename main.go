package main

import (
	"apiREST/configs"
	"apiREST/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load() /* Retorna um erro */
	// Valida o erro, se não carregar as configs - panic
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	// Declarando as rotas
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
