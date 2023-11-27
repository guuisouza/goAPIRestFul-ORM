// Arquivo responsável pela conexão com o banco de dados
package db

import (
	"apiREST/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// func que abre uma conexão com o banco de dados
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB() // pega as configurações do banco

	//String de conexão com o banco postgres
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	//Abrindo conexao com o banco
	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}
	err = conn.Ping() // Verificar se a conexão está estabelecida

	return conn, err

}
