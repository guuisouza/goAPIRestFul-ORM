// Separando as operações por arquivo pra ficar mais fácil
// de ver o que está acontecendo em cada um deles, para uma maneira mais didática
package models

import "apiREST/db"

// Função que será publica para ser chamada nos handlers posteriormente
// Retornará um ID e um erro em caso de falha no insert
func Insert(album Album) (id int64, err error) {
	// Abrindo conexao com o banco
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// Sql que será executado
	// Fará o insert e irá retornar o id que ele gerou após a inserção
	sql := `INSERT INTO albums (title, artist, description, price) VALUES ($1, $2, $3, $4) RETURNING id`

	err = conn.QueryRow(sql, album.Title, album.Artist, album.Description, album.Price).Scan(&id)

	// Retorna o ID e o erro (se houver), o erro será tratado no handler
	return
}
