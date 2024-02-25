![imagem do link](https://www.elegantthemes.com/blog/wp-content/uploads/2015/02/custom-trackable-short-url-feature.png)

# Encurtador de links

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
POST | /users/singin | SE logar na rede
POST |  /users/singup | SE cadastrar
POST | /link | criar link
GET | /link | buscar todos link
GET | /link/:hash | buscar um link
PATCH | /link/:hash | atualizar um link
DELETE | /link/:hash | deletar um link
PATCH | /link/:hash/click | atualizar a quantidade de clicks 

### Tabelas

- users
- qrcodes
- clicks
- links

### Schema

![imagem do schema](./assets/schema.png)
