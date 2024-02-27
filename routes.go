package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
As rotas farão a validação antes das informações serem inseridas
no banco de dados
*/

type Item struct {
	ID    string
	Nome  string
	Valor uint16
	Tipo  string
}

func acessDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "{'erro':'Method Not Allowed','code':405}", http.StatusMethodNotAllowed)
	}
	const docs = `
	<!DOCTYPE html>
	<html lang="pt-br">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Instruções de Uso da API</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				line-height: 1.6;
				margin: 20px;
				background-color: #f8f9fa;
				color: #343a40;
			}
			h1 {
				font-size: 28px;
				border-bottom: 1px solid #e9ecef;
				padding-bottom: 10px;
				margin-bottom: 20px;
			}
			h2 {
				font-size: 24px;
				margin-top: 30px;
			}
			h3 {
				font-size: 20px;
				margin-top: 20px;
			}
			p, ul, ol {
				margin-bottom: 20px;
			}
			pre {
				background-color: #f3f4f6;
				border-radius: 5px;
				padding: 10px;
				overflow-x: auto;
			}
			code {
				background-color: #f3f4f6;
				padding: 2px 5px;
				border-radius: 3px;
				font-family: monospace;
			}
		</style>
	</head>
	<body>
	
	<h1>Instruções de Uso da API</h1>
	
	<h2>Introdução</h2>
	
	<p>Bem-vindo às instruções de uso da nossa API. Abaixo estão os detalhes de como você pode interagir com nossos serviços.</p>
	
	<h2>Autenticação</h2>
	
	<p>Todas as solicitações à API devem ser autenticadas usando uma chave de API. Para autenticar, inclua a chave de API no cabeçalho <code>Authorization</code> de todas as solicitações HTTP.</p>
	
	<pre><code>Authorization: Bearer SUA_CHAVE_DE_API</code></pre>
	
	<h2>Endpoints Disponíveis</h2>
	
	<h3>Obter Recursos</h3>
	
	<p>Para obter recursos específicos, faça uma solicitação GET para o seguinte endpoint:</p>
	
	<pre><code>GET /api/recursos/{id}</code></pre>
	
	<p>Substitua <code>{id}</code> pelo ID do recurso desejado.</p>
	
	<h3>Criar Recurso</h3>
	
	<p>Para criar um novo recurso, faça uma solicitação POST para o seguinte endpoint:</p>
	
	<pre><code>POST /api/recursos</code></pre>
	
	<p>Inclua os dados do recurso no corpo da solicitação.</p>
	
	<pre><code>{
		"nome": "Novo Recurso",
		"descricao": "Descrição do novo recurso"
	}</code></pre>
	
	<h2>Exemplos de Respostas</h2>
	
	<h3>Resposta de Sucesso</h3>
	
	<p>Se a solicitação for bem-sucedida, você receberá uma resposta com o código de status 200 e os dados do recurso solicitado.</p>
	
	<pre><code>{
		"id": 1,
		"nome": "Recurso 1",
		"descricao": "Descrição do Recurso 1"
	}</code></pre>
	
	<h3>Resposta de Erro</h3>
	
	<p>Se ocorrer um erro durante o processamento da solicitação, você receberá uma resposta com o código de status correspondente e uma mensagem de erro.</p>
	
	<pre><code>{
		"mensagem": "Recurso não encontrado"
	}</code></pre>
	
	</body>
	</html>
	
	`
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, docs)
		return
	}

}

func getItemByID(w http.ResponseWriter, r *http.Request) {

}

func deleteItemByID(w http.ResponseWriter, r *http.Request) {

}

func updateItemByID(w http.ResponseWriter, r *http.Request) {

}

func createItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"code":405}`, http.StatusMethodNotAllowed)
	}

	if r.Method == http.MethodPost {
		e, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println("erro ", err)
		}

		fmt.Println(e)
	}
}
