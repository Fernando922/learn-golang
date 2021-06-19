package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver de conexão com o MySql
)

//Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro  //retorna o valor zero como nil
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil

}
