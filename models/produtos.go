package models

import (
	"fmt"
	"site_loja/db"
)

type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.Connection()

	query, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		fmt.Println(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = query.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			fmt.Println(err.Error())
		}

		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func AdicionaProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.Connection()

	query, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		fmt.Println(err.Error())
	}

	query.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.Connection()

	query, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		fmt.Println(err.Error())
	}

	query.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.Connection()

	query, err := db.Query("select * from produtos where id = $1", id)
	if err != nil {
		fmt.Println(err.Error())
	}

	atualizar := Produto{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = query.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			fmt.Println(err.Error())
		}

		atualizar.ID = id
		atualizar.Nome = nome
		atualizar.Descricao = descricao
		atualizar.Preco = preco
		atualizar.Quantidade = quantidade
	}
	defer db.Close()
	return atualizar
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.Connection()

	query, err := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")
	if err != nil {
		fmt.Println(err.Error())
	}

	query.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
