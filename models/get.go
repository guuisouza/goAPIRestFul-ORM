package models

import "apiREST/db"

// Função get que espera um ID e retorna um album ou um erro
// Busca somente 1 dos albums cadastrados no banco de dados
func Get(id int64) (album Album, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	// Fecha a conexão quando a função get for finalizada
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM albums WHERE id=$1`, id)

	// Fará o scan em um ponteiro para cada um dos atributos de albums
	err = row.Scan(&album.ID, &album.Title, &album.Artist, &album.Description, &album.Price)

	return

}
