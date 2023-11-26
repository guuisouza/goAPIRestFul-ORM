package models

// Definindo a struct para carregar os dados pra dentro do nosso banco

type Album struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Artist      string  `json:"artist"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
