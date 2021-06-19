package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

//Criar usuário insere um usuário no banco de dados
func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		rw.Write([]byte("Falha ao ler o corpo da requisição"))
		return //para a execução da função
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		rw.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	fmt.Println(usuario)

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	//Prepare statement (cria um comando de inserção, para evitar sql injection)

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?,?)")

	if erro != nil {
		rw.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		rw.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		rw.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	//STATUS CODES
	rw.WriteHeader(201)
	rw.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))

}

//Buscar usuários traz todos os usuários salvos no banco de dados
func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	//SELECT * FROM USUARIOS

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		rw.Write([]byte("Erro ao buscar os usuários"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			rw.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)

	}

	rw.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(rw).Encode(usuarios); erro != nil {
		rw.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

//BuscarUsuario traz um usuário especifico salvo no banco de dados
func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("só um usuario"))
}
