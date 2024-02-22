# Catálogo de Itens API

Bem-vindo ao Catálogo de Itens API! Este projeto fornece uma API para gerenciar um catálogo de itens gerais. Os itens podem ser de qualquer tipo, como peças de carro, produtos de uma loja, livros em uma biblioteca, entre outros. A API segue o padrão JSON para representar os itens, com os seguintes campos:

- `nome`: O nome do item.
- `valor`: O valor do item.
- `id`: Um identificador único do item (UUID).
- `tipo`: O tipo do item.

## Recursos Disponíveis

### Listar Itens

Endpoint: `/itens`
Método: GET

Retorna uma lista de todos os itens no catálogo.

Exemplo de chamada com cURL:

```bash
curl -X GET http://localhost:3000/itens
```

### Obter Item por ID

Endpoint: `/itens/{id}`
Método: GET

- Retorna um item específico com base no ID fornecido.

Exemplo de chamada com cURL:

```bash
curl -X GET http://localhost:3000/itens/123e4567-e89b-12d3-a456-426614174000
```


### Adicionar Novo Item

Endpoint: `/itens`
Método: POST

- Adiciona um novo item ao catálogo.

Exemplo de chamada com cURL:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"nome":"Novo Item","valor":99.99,"id":"123e4567-e89b-12d3-a456-426614174000","tipo":"livro"}' http://localhost:3000/itens
```

### Atualizar Item

Endpoint: `/itens/{id}`
Método: PUT

- Atualiza um item existente com base no ID fornecido.

Exemplo de chamada com cURL:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"nome":"Item Atualizado","valor":149.99,"tipo":"livro"}' http://localhost:3000/itens/123e4567-e89b-12d3-a456-426614174000
```

### Remover Item

Endpoint: `/itens/{id}`
Método: DELETE

- Remove um item do catálogo com base no ID fornecido.

Exemplo de chamada com cURL:

```bash
curl -X DELETE http://localhost:3000/itens/123e4567-e89b-12d3-a456-426614174000
```

## Executando Localmente

## Para executar este projeto localmente, siga estas etapas:

1. Clone este repositório.
2. Certifique-se de ter o [Golang](https://go.dev/dl) (1.21.6) instalado em sua máquina.
3. Instale as dependências do projeto executando `go mod tidy`.
4. Inicie o servidor executando `go run main.go`.
5. A API estará disponível em `http://localhost:8000`.

## Contribuindo

Se você quiser contribuir com este projeto, fique à vontade para enviar pull requests ou abrir issues. Toda contribuição é bem-vinda!

## Licença

Este projeto está licenciado sob a [MIT License](https://opensource.org/licenses/MIT).
