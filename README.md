![imagem do link](https://www.elegantthemes.com/blog/wp-content/uploads/2015/02/custom-trackable-short-url-feature.png)

# Encurtador de links

## Frontend

## Backend

### Tecnologias

- Golang
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

### Tabelas

- users
  nome | tipo
  ---| ---
  name | string
  email | string
  password | string
  birth |

- clicks
- links
