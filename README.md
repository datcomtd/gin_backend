# DATCOM-TD backend

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

```
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

```
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

```
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

#### /api/user/\<username\>/update

Atualiza as informações de um usuário.  
Use newpassword para atualizar a senha ao invés de password, que foi reservado para a autentificação por post request.

```bash
$ curl -s -L -X POST -H "Content-Type: application/json" \
  -d "{\"username\": \"patrick\", \"password\": \"patrick123\", \"email\": \"newemail@gmail.com\"}" \
  http://localhost:8000/api/user/patrick/update | jq '.'

```

```
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

## TODO

- [ ] Migrate the DB from SQLite to PostgreSQL (sync to async)
