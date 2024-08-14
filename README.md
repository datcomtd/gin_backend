# DATCOM-TD backend

<img align="right" width="128px" src=".media/coffee.png">

Esse repositório contém uma implementação de uma REST API para o backend do site do DATCOM-TD. Foi escrito em [Go](https://go.dev/) utilizando o framework [Gin](https://github.com/gin-gonic/gin/).

Estatisticas:

| Linguagem | Arquivos | Linhas   | Blanks  | Comentários |
|:---------:|:--------:|:--------:|:-------:|:-----------:|
| Go        | 19       | 1218     | 222     | 217 |
| Bash      | 3        | 200      | 28      | 5 |
| Markdown  | 2        | 599      | 105     | 0 |
| **Total** | **24**   | **2017** | **355** | **222** |

As rotas (endpoints) implementadas estão listadas nas tabelas abaixo:

| USER Endpoint                        | Request | Token | Auth |
|:-------------------------------------|:-------:|:-----:|:----:|
| [/api/register](#register)           | POST    | x     | x |
| [/api/token](#token)                 | POST    | x     | o |
| /api/users                           | GET     | x     | x |
| [/api/user/\<username\>](#user-read) | GET     | x     | x |
| [/api/user/update](#user-update)     | POST    | x     | o |
| [/api/user/delete](#user-delete)     | POST    | x     | o |
| [/api/user/picture](#user-picture)   | POST    | o     | x |

| DOCUMENT Endpoint                      | Request | Token | Auth |
|:---------------------------------------|:-------:|:-----:|:----:|
| [/api/documents](#doc-read-all)                              | GET     | x     | x |
| [/api/document/by-id/\<id\>](#doc-read-id)                   | GET     | x     | x |
| [/api/document/by-category/\<category\>](#doc-read-category) | GET     | x     | x |
| [/api/document/upload/{\<key\>}](#doc-upload)                | POST    | o     | x |
| [/api/document/update](#doc-update)                          | POST    | o     | x |
| [/api/document/delete](#doc-delete)                          | POST    | o     | x |

| PRODUCT Endpoint                              | Request | Token | Auth |
|:----------------------------------------------|:-------:|:-----:|:----:|
| [/api/products](#product-read-all)            | GET     | x     | x    |
| [/api/product/by-id/\<id\>](#product-read-id) | GET     | x     | x    |
| [/api/product/create](#product-create)        | POST    | o     | x    |
| [/api/product/update](#product-update)        | POST    | o     | o    |
| [/api/product/delete](#product-delete)        | POST    | o     | x    |

## Instruções

Instale as dependências:

```bash
$ go get .
```

OPCIONAL: Para poder utilizar sua conta GMail com SMTP no backend, crie uma "senha de app" seguindo [esse passo-a-passo](https://canaltech.com.br/internet/como-usar-o-gmail-como-servidor-smtp/). Depois, modifique o env de acordo:

```bash
# OPCIONAL
$ vim initializers/env.go
var SenderEmail string = "alexandreboutrik@alunos.utfpr.edu.br"
var SenderPassword string = "<your app password>"
```

Configure o PostgreSQL:

```bash
$ sudo vim /etc/postgresql-16/postgresql.conf # ou equivalente
listen_addresses = 'localhost'
port = 4145

$ sudo vim /etc/postgresql-16/pg_hba.conf # caso exista esse arquivo
local all postgres trust
local all all md5
```

Inicie o servidor do postgres:

```bash
$ sudo systemctl start postgresql
```

Para resetar, criar um novo banco de dados:

```bash
$ ./reset.sh
```

Para iniciar o backend do DATCOM-TD:

```bash
$ go run .
```

O servidor estará esperando por conexões em 127.0.0.1:8000.

## Documentação da REST API

### USER endpoints

<h4 id="register">
:book:&nbsp;&nbsp;/api/register
</h4>

Registra um novo usuário no banco de dados.  
Somente o Presidente pode registrar novos usuários.

| Enum | Role/Função |
|:----:|:-----|
| 1    | Presidente |
| 2    | Vice de Computação |
| 3    | Vice de TSI |
| 4    | Secretário(a) |
| 5    | Tesoureiro(a) |
| 6    | Diretor(a) |
| ≥ 7  | Outros |

| Enum | Curso |
|:----:|:------|
| 1    | Computação |
| 2    | TSI |

| Field | Type | Required |
|:-----:|:----:|:--------:|
| admin-username | string | yes |
| admin-password | string | yes |
| username  | string | yes |
| email     | string | no  |
| password  | string | yes |
| role      | enum   | yes |
| course    | enum   | yes |

```bash
$ curl -s -L -X POST -H "Content-Type: application/json" \
  -d "{\"admin-username\": \"admin\", \"admin-password\": \"eb8fac70478d46e4c68c\", \"username\": \"patrick\", \"password\": \"patrick123\", \"role\": 1, \"course\": 1}" \
  http://localhost:8000/api/register | jq '.'
```

```json
{
  "user": {
    "CreatedAt": "2024-07-19T22:44:22.11276403-03:00",
    "UpdateAt": "0001-01-01T00:00:00Z",
    "Token_UpdatedAt": "2024-07-20T01:44:22.112720565Z",
    "Token": "VUVJsMvvLUSOicICknLsJpARmNnCXfAallxEeySjVksVsCadBDoGvQftSisiooXj",
    "Password": "$2a$14$r/3GZdaJFzCVyQAnNdhjm.Ya9IkR7pfwtRKgVzpc3661iQScPKnwS",
    "name": "patrick",
    "email": "",
    "role": 1,
    "course": 1
  }
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 400    | Bad Request    | required fields are not filled |
| 400    | Bad Request    | user is already registered |
| 400    | Bad Request    | invalid course |
| 401    | Unauthorized   | invalid admin-key |
| 500    | Internal Error | failed hashing the password |
| 500    | Internal Error | failed creating the record |
| 201    | Created        | - |

<h4 id="token">
:book:&nbsp;&nbsp;/api/token
</h4>

Obtém o token de autentificação de um usuário.

| Field | Type | Required |
|:-----:|:----:|:--------:|
| username | string | yes |
| password | string | yes |

```bash
$ curl -s -L -X POST -H "Content-Type: application/json" \
  -d "{\"username\": \"patrick\", \"password\": \"patrick123\"}" \
  http://localhost:8000/api/token | jq '.'
```

```json
{
  "token": "VUVJsMvvLUSOicICknLsJpARmNnCXfAallxEeySjVksVsCadBDoGvQftSisiooXj"
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 400    | Bad Request    | required fields are not filled |
| 401    | Unauthorized   | invalid username or password |
| 500    | Internal Error | failed updating the record |
| 200    | OK             | - |

<h4 id="user-read">
:book:&nbsp;&nbsp;/api/user/&lt;username&gt;
</h4>

Obtém as informações públicas de algum usuário.

```bash
$ curl -s -L http://localhost:8000/api/user/patrick | jq '.'
```

```json
{
  "user": {
    "name": "patrick",
    "email": "",
    "role": 1,
    "course": 1
  }
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 404    | Not Found      | user not found |
| 200    | OK             | - |

<h4 id="user-update">
:book:&nbsp;&nbsp;/api/user/update
</h4>

Atualiza as informações de um usuário.  
Use newpassword para atualizar a senha ao invés de password, que foi reservado para a autentificação por post request.

| Field | Type | Required |
|:-----:|:----:|:--------:|
| admin-username | string | no |
| admin-password | string | no |
| username | string | yes |
| password | string | yes |
| newpassword | string | no |
| email  | string | no |
| role   | enum   | no |
| course | enum   | no |

```bash
$ curl -s -L -X POST -H "Content-Type: application/json" \
  -d "{\"username\": \"patrick\", \"password\": \"patrick123\", \"email\": \"newemail@gmail.com\"}" \
  http://localhost:8000/api/user/update | jq '.'

```

```json
{
  "user": {
    "CreatedAt": "2024-07-20T02:32:21.950871021-03:00",
    "UpdateAt": "0001-01-01T00:00:00Z",
    "Token_UpdatedAt": "2024-07-20T05:32:21.95082028Z",
    "Token": "AMFXsITtGPFGSUaJvETnuSxwmRDHsTomsjBHJZKZYErYaagSeuEzfSpjeRdWSsGR",
    "Password": "$2a$14$KlOQnE10.KGjC2R0wsWaL.Vw75bLMIMztaDgKqvI7n4AJy4stmnT.",
    "name": "patrick",
    "email": "newemail@gmail.com",
    "role": 1,
    "course": 1
  }
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 400    | Bad Request    | required fields are not filled |
| 401    | Unauthorized   | invalid username or password |
| 500    | Internal Error | failed hashing the password |
| 500    | Internal Error | failed updating the record |
| 200    | OK             | - |

<h4 id="user-delete">
:book:&nbsp;&nbsp;/api/user/delete
</h4>

Deleta algum usuário do banco de dados.  
Não foi implementado nenhum tipo de soft delete, então a remoção é permanente.

| Field | Type | Required |
|:-----:|:----:|:--------:|
| admin-username | string | no |
| admin-password | string | no |
| username | string | yes |
| password | string | yes |

```bash
$ curl -s -L -X POST -H "Content-Type: application/json" \
  -d "{\"username\": \"patrick\", \"password\": \"patrick123\"}" \
  http://localhost:8000/api/user/delete | jq '.'
```

```json
{
  "message": "user deleted"
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 400    | Bad Request    | required fields are not filled |
| 401    | Unauthorized   | invalid username or password |
| 500    | Internal Error | failed deleting the record |
| 200    | OK             | user deleted |

<h4 id="user-picture">
:book:&nbsp;&nbsp;/api/user/picture
</h4>

Upload a profile picture.

```bash
$ curl -s -L -X POST \
  -H "Authorization: <token>" \
  -H "Content-Type: multipart/form-data" \
  -F "file=@<filename>" \
  http://localhost:8000/api/user/picture | jq '.'
```

```json
{
  "message": "picture uploaded"
}
```

| Code | Status          | Message |
|:----:|:---------------:|:--------|
| 400  | Bad Request     | invalid file |
| 400  | Bad Request     | invalid extension |
| 401  | Unauthorized    | invalid token |
| 500  | Internal Error  | failed uploading the file |
| 201  | Created         | - |

### DOCUMENT endpoints

<h4 id="doc-read-all">
:book:&nbsp;&nbsp;/api/documents
</h4>

Obtém a lista de todos os documentos.

```bash
$ curl -s -L http://localhost:8000/api/documents | jq '.'
```

```json
{
  "document": [
    {
      "CreatedAt": "2024-07-20T11:55:30.107290365-03:00",
      "UpdateAt": "0001-01-01T00:00:00Z",
      "id": 1,
      "Key": "nZKoUVWsWxHJqQillbJVjeZvmfdUQZvM",
      "title": "Simple Text File",
      "description": "",
      "source": "Tester",
      "category": "edital",
      "created-by": "patrick",
      "last-updated-by": "patrick"
    }
  ]
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 200    | OK             | -       |

<h4 id="doc-read-id">
:book:&nbsp;&nbsp;/api/document/by-id/&lt;id&gt;
</h4>

Obtém as informações de um documento.

```bash
$ curl -s -L http://localhost:8000/api/document/by-id/1 | jq '.'
```

```json
{
  "document": {
      "CreatedAt": "2024-07-20T11:55:30.107290365-03:00",
      "UpdateAt": "0001-01-01T00:00:00Z",
      "id": 1,
      "Key": "nZKoUVWsWxHJqQillbJVjeZvmfdUQZvM",
      "title": "Simple Text File",
      "description": "",
      "source": "Tester",
      "category": "edital",
      "created-by": "patrick",
      "last-updated-by": "patrick"
    }
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 404    | Not Found      | document not found |
| 200    | OK             | - |

<h4 id="doc-read-category">
:book:&nbsp;&nbsp;/api/document/by-category/&lt;category&gt;
</h4>

Obtém uma lista de documentos de uma categoria.

```bash
$ curl -s -L http://localhost:8000/api/document/by-category/edital | jq '.'
```

```json
{
  "document": {
      "CreatedAt": "2024-07-20T11:55:30.107290365-03:00",
      "UpdateAt": "0001-01-01T00:00:00Z",
      "id": 1,
      "Key": "nZKoUVWsWxHJqQillbJVjeZvmfdUQZvM",
      "title": "Simple Text File",
      "description": "",
      "source": "Tester",
      "category": "edital",
      "created-by": "patrick",
      "last-updated-by": "patrick"
    }
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 200    | OK             | -       |

<h4 id="doc-upload">
:book:&nbsp;&nbsp;/api/document/upload e /api/document/upload/&lt;key&gt;
</h4>

Faz upload de algum documento.

O upload é feito em duas etapas:
1. Envio de metadados do arquivo (titulo, descrição, orgão e categoria) por um post request. Será gerado uma chave key.
2. Envio do arquivo para /api/document/upload/\<key\>.  

Os documentos são salvos em ./media/\<key\>\_\<filename\>.

| Field | Type | Required |
|:-----:|:----:|:--------:|
| title       | string | yes |
| source      | string | yes |
| category    | string | yes |
| filename    | string | no |
| description | string | no |

```bash
# Envio de metadados e geração da chave key
$ curl -s -L -X POST \
  -H "Authorization: \<token\>" \
  -H "Content-Type: application/json" \
  -d "{\"title\": \"O Senhor dos Anéis\", \"source\": \"J. R. R. Tolkien\", \"category\": \"fictional book\"}" \
  http://localhost:8000/api/document/upload | jq '.'

$ curl -s -L -X POST \
  -H "Authorization: \<token\>" \
  -H "Content-Type: multipart/form-data" \
  -F "file=@senhor_dos_aneis.pdf" \
  http://localhost:8000/api/document/upload/\<key\> | jq '.'
```

```json
{
  "key": "kGVnythePZEOGjKHRIVdkzimYIWFHHQC"
}
```

```json
{
  "document": {
    "CreatedAt": "2024-07-20T12:11:28.770287042-03:00",
    "UpdateAt": "0001-01-01T00:00:00Z",
    "id": 1,
    "Key": "kGVnythePZEOGjKHRIVdkzimYIWFHHQC",
    "filename": "senhor_dos_aneis.pdf",
    "title": "O Senhor dos Anéis",
    "description": "",
    "source": "J. R. R. Tolkien",
    "category": "fictional book",
    "created-by": "patrick",
    "last-updated-by": "patrick"
  }
}
```

```bash
$ ls ./media
kGVnythePZEOGjKHRIVdkzimYIWFHHQC_senhor_dos_aneis.pdf
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 400    | Bad Request    | required fields are not filled |
| 400    | Bad Request    | file already exists |
| 400    | Bad Request    | invalid document |
| 401    | Unauthorized   | invalid token |
| 401    | Unauthorized   | invalid key |
| 403    | Forbidden      | user does not have permission |
| 500    | Internal Error | failed creating the record |
| 500    | Internal Error | failed updating the record |
| 500    | Internal Error | failed saving the document |
| 200    | OK             | (for the key step) |
| 201    | Created        | (document uploaded) |

<h4 id="doc-update">
:book:&nbsp;&nbsp;/api/document/update
</h4>

Atualiza os metadados de um documento.  
Para modificar o documento em si, remova ele e faça upload de um novo documento. O motivo pelo qual optei por não implementar um update do documento em si pode ser encontrado [aqui](https://philsturgeon.com/http-rest-api-file-uploads/).

| Field | Type | Required |
|:-----:|:----:|:--------:|
| id          | integer | yes |
| filename    | string  | no  |
| title       | string  | no  |
| description | string  | no  |
| source      | string  | no  |
| category    | string  | no  |

```bash
$ curl -s -L -X POST \
    -H "Authorization: \<token\>" \
    -H "Content-Type: application/json" \
    -d "{\"id\": 1, \"filename\": \"novo_nome.pdf\"}" \
    http://localhost:8000/api/document/update | jq '.'
```

```json
{
  "document": {
    "CreatedAt": "2024-07-20T12:11:28.770287042-03:00",
    "UpdateAt": "0001-01-01T00:00:00Z",
    "id": 1,
    "Key": "kGVnythePZEOGjKHRIVdkzimYIWFHHQC",
    "filename": "novo_nome.pdf",
    "title": "O Senhor dos Anéis",
    "description": "",
    "source": "J. R. R. Tolkien",
    "category": "fictional book",
    "created-by": "patrick",
    "last-updated-by": "patrick"
  }
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 400    | Bad Request    | required fields are not filled |
| 401    | Unauthorized   | invalid token |
| 403    | Forbidden      | user does not have permission |
| 404    | Not Found      | document not found |
| 500    | Internal Error | failed renaming the file |
| 500    | Internal Error | failed updating the record |
| 200    | OK             | - |

<h4 id="doc-delete">
:book:&nbsp;&nbsp;/api/document/delete
</h4>

Deleta um documento.  
Somente o criador (que fez upload) daquele documento tem permissão para removê-lo.

| Field | Type | Required |
|:-----:|:----:|:--------:|
| admin-username | string | no |
| admin-password | string | no |
| id    | integer | yes |

```bash
$ curl -s -L -X POST \
    -H "Authorization: \<token\>" \
    -H "Content-Type: application/json" \
    -d "{\"id\": 1}" \
    http://localhost:8000/api/document/delete | jq '.'
```

```json
{
  "message": "document deleted"
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 400    | Bad Request    | required fields are not filled |
| 401    | Unauthorized   | invalid token |
| 403    | Forbidden      | user is not the document's creator |
| 404    | Not Found      | document not found |
| 500    | Internal Error | failed deleting the document |
| 500    | Internal Error | failed deleting the record |
| 200    | OK             | document deleted |

### PRODUCT endpoints

<h4 id="product-read-all">
:book:&nbsp;&nbsp;/api/products
</h4>

```bash
$ curl -s -L http://localhost:8000/api/products | jq '.'
```

```json
{
  "product": [
    {
      "CreatedAt": "2024-08-14T19:04:08.910267+01:00",
      "UpdateAt": "0000-12-31T23:58:45-00:01",
      "id": 1,
      "count": 0,
      "photos": null,
      "title": "camiseta DATCOM",
      "description": "",
      "in-stock": true,
      "created-by": "patrick",
      "last-updated-by": "patrick"
    }
  ]
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 200    | OK             | - |

<h4 id="product-read-id">
:book:&nbsp;&nbsp;/api/product/by-id/&lt;id&gt;
</h4>

```bash
$ curl -s -L http://localhost:8000/api/product/by-id/1 | jq '.'
```

```json
{
  "product": {
      "CreatedAt": "2024-08-14T19:04:08.910267+01:00",
      "UpdateAt": "0000-12-31T23:58:45-00:01",
      "id": 1,
      "count": 0,
      "photos": null,
      "title": "camiseta DATCOM",
      "description": "",
      "in-stock": true,
      "created-by": "patrick",
      "last-updated-by": "patrick"
    }
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 404    | Not Found      | product not found |
| 200    | OK             | - |

<h4 id="product-create">
:book:&nbsp;&nbsp;/api/product/create
</h4>

Cria um produto da lojinha.

| Field | Type | Required |
|:-----:|:----:|:--------:|
| title       | string  | yes |
| description | string  | no |

```bash
$ curl -s -L -X POST \
  -H "Authorization: <token>" \
  -H "Content-Type: application/json" \
  -d "{\"title\": \"Camiseta DATCOM-TD\", \"description\": \"...\"}" \
  http://localhost:8000/api/product/create | jq '.'
```

```json
{
  "message": "product created"
}
```

| Código | Status         | Message |
|:------:|:--------------:|:--------|
| 400    | Bad Request    | required fields are not filled |
| 400    | Bad Request    | product already exists |
| 401    | Unauthorized   | invalid token |
| 403    | Forbidden      | user does not have permission |
| 500    | Internal Error | failed creating the record |
| 201    | Created        | - |

<h4 id="product-update">
:book:&nbsp;&nbsp;/api/product/update
</h4>

| Código | Status          | Message |
|:------:|:---------------:|:-------:|
| 501    | Not Implemented | not implemented yet |

<h4 id="product-delete">
:book:&nbsp;&nbsp;/api/product/delete
</h4>

| Código | Status          | Message |
|:------:|:---------------:|:-------:|
| 501    | Not Implemented | not implemented yet |

## TODO

- [x] Migrate the DB from SQLite to PostgreSQL (sync to async)
