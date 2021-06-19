package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //import implicito esse arquivo nao vai usar, mas o database/sql vai
)

//go get github.com/go-sql-driver/mysql

func main() {
	//urlConexao := "usuario:senha@/banco"
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close() //antes de terminar a função main a conexão é fechada

	//a variável foi reaproveitada
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")

	if(erro != nil){
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
