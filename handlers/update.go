package handlers

//chi
import (
	"apiREST/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	// Vamos converter os parametros (id) da URL para INT através do strconv.Atoi
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		// Recebe 3 parâmetros (response writer, mensagem de erro e o statusCode)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var album models.Album

	err = json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		// Recebe 3 parâmetros (response writer, mensagem de erro e o statusCode)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Se não houve nenhumerro podemos fazer o update
	rows, err := models.Update(int64(id), album)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Se atualizou mais de uma linha
	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "dados atualizados com sucesso",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp) // Faz um encode do map pra um json de resposta
}
