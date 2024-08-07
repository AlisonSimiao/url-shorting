![imagem do link](https://www.elegantthemes.com/blog/wp-content/uploads/2015/02/custom-trackable-short-url-feature.png)

# Encurtador de links

## init

- instalar golang
- go mod tidy
- go run main.go

status code | descrição
--- | ---
200 | ok
201 | created
204 | no response
400 | erro request
401 | no authorized
403 | forbidden
404 | no found
409 | conflict
422 | umprocesss
429 | rate limit

## Acesso

#### URL

- <https://vagas-api-7jko.onrender.com>

## Frontend

### Tecnologias

- TypeScript
  - Next.js
  - Tailwind CSS

### Telas

- Login
- Cadastro
- Redirect
- Main Page
- Perfil

## Backend

### Tecnologias

- Golang
  - Gin-gonic
  - Golang-jwt
  - Google-Uuid
  - GoValidator
  - GoDotenv
  - Crypto
  - Postgres
  - Gorm
- Postgres

### Routers

Metodo | rota | descrição
---  | --- | ---
POST | /users/singin | Se logar na rede
POST |  /users/singup | Se cadastrar
GET | /users/:username | Pesquisar usuário
PATCH | /users/:username | Alterar usuário

### Tabelas

- users
