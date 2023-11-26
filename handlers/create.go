// Adicionará uma nova entrada no banco de dados
package handlers

import (
	"apiREST/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Função http
func Create(w http.ResponseWriter, r *http.Request) {
	// Faz o parse da request recebida
	var album models.Album

	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		// Recebe 3 parâmetros (response writer, mensagem de erro e o statusCode)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(album)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true, // Poderia ser feito com o status code da resposta também
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false, // Poderia ser feito com o status code da resposta também
			"Message": fmt.Sprintf("Album inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp) // Faz um encode do map pra um json de resposta
}
