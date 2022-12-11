package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"site_loja/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // Retorna os templates .html apontados no arquivo templates

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Erro na conversão do preço:", err.Error())
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			fmt.Println("Erro na conversão da quantidade:", err.Error())
		}

		models.AdicionaProduto(nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeletaProduto(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.EditaProduto(id)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Erro na conversão do id:", err.Error())
		}

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Erro na conversão do preço:", err.Error())
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			fmt.Println("Erro na conversão da quantidade:", err.Error())
		}

		models.AtualizaProduto(idConv, nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", 301)
}
