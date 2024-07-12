# DATCOM-TD backend

## Instruções

Instale as dependências e use o migrate.go:

```bash
$ go get .
$ go run migrate/migrate.go
```

Inicie o servidor:

```bash
$ go run .
```

O servidor estará esperando por conexões em 127.0.0.1:8000.

## REST API

### Editais endpoints

Para conseguir a lista de documentos da página de editais, envie um GET request para /api/editais/:

```bash
$ curl http://127.0.0.1:8000/api/editais/ | jq '.'
```
