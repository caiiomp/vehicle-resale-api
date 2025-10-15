# API de Veículos

Este repositório contém a API para gestão de veículos, permitindo o cadastro, listagem e compra de veículos. A autenticação e autorização dos usuários (compradores) é realizada pelo serviço de autenticação JWT.

## Funcionalidades

- **Cadastro de veículos:** Permite o cadastro de veículos à venda (marca, modelo, ano, cor e preço).
- **Edição de veículos:** Permite a edição dos dados de veículos cadastrados.
- **Listagem de veículos à venda:** Exibe os veículos à venda, ordenados por preço (do mais barato ao mais caro).
- **Listagem de veículos vendidos:** Exibe os veículos vendidos, também ordenados por preço.
- **Compra de veículos:** Permite que usuários autenticados comprem veículos. A operação de compra requer que o comprador esteja autenticado (com um token JWT válido).

## Tecnologias Utilizadas

- **Go (Golang):** Para o desenvolvimento da API de veículos.
- **MongoDB:** Para o armazenamento dos dados de veículos.
- **Gin:** Framework web para o desenvolvimento da API.
- **JWT (JSON Web Tokens):** Para autenticação e autorização de usuários.
- **Docker Compose:** Para o setup do MongoDB via Docker.

## Como Rodar o Projeto Localmente

### 1. Pré-requisitos

Certifique-se de que você tem as seguintes dependências instaladas:

- **Go (Golang)** versão 1.18 ou superior
- **Git** para clonar o repositório
- **Docker** e **Docker Compose** (para MongoDB)

### 2. Configuração do MongoDB com Docker Compose

1. Clone o repositório:

    ```bash
    git clone git@github.com:caiiomp/vehicle-resale-api.git
    ```

2. Na raiz do projeto, inicie o MongoDB com o `docker`:

    ```bash
    docker compose up -d
    ```

    Isso irá iniciar o MongoDB em um contêiner.

### 3. Configuração da API de Veículos

1. Na raiz do projeto instale as dependências do Go:

    ```bash
    go mod tidy
    ```

2. Inicie o servidor da API:

    ```bash
    go run src/main.go
    ```

    A API de veículos estará disponível em `http://localhost:4001`.

### 4. Testando a API de Veículos

Use **Postman**, **Insomnia**, **cURL** ou qualquer outro cliente **HTTP** para testar os endpoints:

- `POST /vehicles` - Cadastrar um novo veículo (necessário token JWT de autenticação).
- `GET /vehicles?is_sold=false` - Listar todos os veículos à venda.
- `GET /vehicles?is_sold=true` - Listar todos os veículos vendidos.
- `GET /vehicles/:vehicle_id` - Buscar veículo por id.
- `PATCH /vehicles/:vehicle_id` - Editar um veículo existente (necessário token JWT de autenticação).
- `POST /vehicles/:vehicle_id/buy` - Comprar um veículo (necessário token JWT de autenticação).
- `GET /sales` - Listar todas as vendas.

### 5. Exemplo de Uso

Para realizar a compra de um veículo, o comprador deve fornecer um **token JWT válido** gerado pelo serviço de autenticação. O token deve ser incluído no cabeçalho da requisição:

```
Authorization: Bearer <token-jwt>
```

```json
// Exemplo de cadastro de veículo
{
    "brand": "Ford",
    "model": "Ka",
    "year": 2022,
    "color": "Preto",
    "price": 50000
}
```

## Documentação (Swagger)

Para acessar a documentação do serviço, acessar o seguinte endpoint: 
```
http://localhost:4000/swagger/index.html
```