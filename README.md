# DATCOM-TD backend

<img align="right" width="128px" src=".media/coffee.png">

Esse repositório contém uma implementação de uma REST API para o backend do site do DATCOM-TD. Foi escrito em [Go](https://go.dev/) utilizando o framework [Gin](https://github.com/gin-gonic/gin/).

Estatisticas:

| Linguagem | Arquivos | Linhas | Blanks | Comentários |
|:---------:|:--------:|:------:|:------:|:-----------:|
| Go        | 18       | 1023   | 193    | 192 |
| Bash      | 3        | 160    | 20     | 5 |
| Markdown  | 1        | 453    | 83     | 0 |
| Text      | 1        | 1      | 0      | 0 |
| **Total** | **23** | **1637** | **296** | **197** |

As rotas (endpoints) implementadas estão listadas nas tabelas abaixo:

| USER Endpoint           | Request | Token | Auth |
|:------------------------|:-------:|:-----:|:----:|
| /api/register           | POST    | x     | x |
| /api/token              | POST    | x     | o |
| /api/users              | GET     | x     | x |
| /api/user/\<username\>  | GET     | x     | x |
| /api/user/update        | POST    | x     | o |
| /api/user/delete        | POST    | x     | o |

| DOCUMENT Endpoint                      | Request | Token | Auth |
|:---------------------------------------|:-------:|:-----:|:----:|
| /api/documents                         | GET     | x     | x |
| /api/document/by-id/\<id\>             | GET     | x     | x |
| /api/document/by-category/\<category\> | GET     | x     | x |
| /api/document/upload/{\<key\>}         | POST    | o     | x |
| /api/document/update                   | POST    | o     | x |
| /api/document/delete                   | POST    | o     | x |

## Instruções

Instale as dependências e use o migrate.go:

```bash
$ go get .
$ go run migrate/migrate.go
```

Para resetar o banco de dados:

```bash
$ ./reset.sh
```

Inicie o servidor:

```bash
$ go run .
```

O servidor estará esperando por conexões em 127.0.0.1:8000.

## REST API

### USER endpoints

#### /api/register

Registra um novo usuário no banco de dados.

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

```bash
$ curl -s -L -X POST -H "Content-Type: application/json" \
  -d "{\"username\": \"patrick\", \"password\": \"patrick123\", \"role\": 1, \"course\": 1}" \
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
| 500    | Internal Error | failed hashing the password |
| 500    | Internal Error | failed creating the record |
| 201    | Created        | - |

#### /api/token

Obtém o token de autentificação de um usuário.

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

#### /api/user/\<username\>

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

#### /api/user/update

Atualiza as informações de um usuário.  
Use newpassword para atualizar a senha ao invés de password, que foi reservado para a autentificação por post request.

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

#### /api/user/delete

Deleta algum usuário do banco de dados.  
Não foi implementado nenhum tipo de soft delete, então a remoção é permanente.

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

### DOCUMENT endpoints

#### /api/documents

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

#### /api/document/by-id/\<id\>

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

#### /api/document/by-category/\<category\>

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

#### /api/document/upload e /api/document/upload/\<key\>

Faz upload de algum documento.

O upload é feito em duas etapas:
1. Envio de metadados do arquivo (titulo, descrição, orgão e categoria) por um post request. Será gerado uma chave key.
2. Envio do arquivo para /api/document/upload/\<key\>.  

Os documentos são salvos em ./media/\<key\>\_\<filename\>.

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

#### /api/document/update

Atualiza os metadados de um documento.  
Para modificar o documento em si, remova ele e faça upload de um novo documento. O motivo pelo qual optei por não implementar um update do documento em si pode ser encontrado [aqui](https://philsturgeon.com/http-rest-api-file-uploads/).

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

#### /api/document/delete

Deleta um documento.  
Somente o criador (que fez upload) daquele documento que tem permissão para removê-lo.

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

## TODO

- [ ] Migrate the DB from SQLite to PostgreSQL (sync to async)
