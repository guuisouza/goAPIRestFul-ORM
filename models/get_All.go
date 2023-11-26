package models

import "apiREST/db"

// Função GetAll  que retorna um slice [] de album ou um erro
// Busca todos os albums
func GetAll() (albums []Album, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	// Fecha a conexão assim que termina a execução do GetAll
	defer conn.Close()

	// Select de tudo que tem no banco de dados (tabela)
	rows, err := conn.Query(`SELECT * FROM albums`)
	if err != nil {
		return
	}

	for rows.Next() { // Percorre todos os itens retornados e faz o scann
		var album Album
		err = rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Description, &album.Price)
		if err != nil {
			continue
		}

		// Se não houver erro adiciona no fim da lista o elemento
		albums = append(albums, album)
	}

	return
}
