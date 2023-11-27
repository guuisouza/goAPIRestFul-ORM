package handlers

import (
	"apiREST/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	// Vamos converter os parametros (id) da URL para INT através do strconv.Atoi
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		// Recebe 3 parâmetros (response writer, mensagem de erro e o statusCode)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Se não houve nenhum erro podemos fazer o get dos models
	album, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Encode em json do album encontrado no banco
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album) // Faz um encode do map pra um json de resposta
}
