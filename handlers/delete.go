package handlers

//chi
import (
	"apiREST/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	// Vamos converter os parametros (id) da URL para INT através do strconv.Atoi
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		// Recebe 3 parâmetros (response writer, mensagem de erro e o statusCode)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Se não houve nenhumerro podemos fazer o delete
	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Se atualizou mais de uma linha
	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "dados removidos com sucesso",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp) // Faz um encode do map pra um json de resposta
}
