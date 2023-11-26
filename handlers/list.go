package handlers

import (
	"apiREST/models"
	"encoding/json"
	"log"
	"net/http"
)

// Função que vai retornar uma lista de albums
func List(w http.ResponseWriter, r *http.Request) {
	albums, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums) /* Retorna um array e nao uma mensagem */
}
