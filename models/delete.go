package models

import "apiREST/db"

// Recebe apenas o Id que será excluido

func Delete(id int64) (int64, error) {
	// Abrindo a conexão com o banco
	conn, err := db.OpenConnection()
	// Valida se tem o erro ou não
	if err != nil {
		return 0, err // não definimos no () então passamos aqui
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM albums WHERE id=$1`, id)
	// Validar se deu erro
	if err != nil {
		return 0, err
	}

	// Se não deu erro, retorna o numero de linhas afetadas pelo delete
	return res.RowsAffected()
}
