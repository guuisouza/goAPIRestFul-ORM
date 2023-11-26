package models

import "apiREST/db"

// Função para atualizar que recebe 2 parametros
// Recebe o Id que será atualizado e o album que será atualizado

func Update(id int64, album Album) (int64, error) {
	//Abrindo a conexão com o banco
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err // não definimos no () então passamos aqui
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE albums SET title=$2, artist=$3, description=$4, price=$5
	WHERE id=$1`, id, album.Title, album.Artist, album.Description, album.Price)
	// Validar se deu erro
	if err != nil {
		return 0, err
	}

	// Se não deu erro, retorna o numero de linhas afetadas pelo update
	return res.RowsAffected()
}
