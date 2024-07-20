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

```bash
curl -s -L -X POST -H "Content-Type: application/json" \
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
| 500    | Internal Error | failed hashing the password |
| 500    | Internal Error | failed creating the record |
| 201    | Created        | success |

#### /api/token

Obtém o token de autentificação de um usuário.

```bash
curl -s -L -X POST -H "Content-Type: application/json" \
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
| 200    | OK             | success |

#### /api/user/\<username\>

Obtém as informações públicas de algum usuário.

```bash
curl -s -L http://localhost:8000/api/user/patrick | jq '.'
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
| 200    | OK             | success |

## TODO

- [ ] Migrate the DB from SQLite to PostgreSQL (sync to async)
