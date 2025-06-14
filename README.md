# DATCOM backend

<img align="right" width="128px" src=".media/coffee.png" />

![Version Badge](https://img.shields.io/badge/version-v1.0.2--beta-blue)

Esse repositório contém uma implementação de uma REST API para o backend do site do DATCOM. Foi escrito em [Go](https://go.dev/) utilizando o framework [Gin](https://github.com/gin-gonic/gin/).

Estatisticas:

| Linguagem  | Arquivos | Linhas   | Blanks  | Comentários |
|:----------:|:--------:|:--------:|:-------:|:-----------:|
| Go         | 27       | 1999     | 371     | 366 |
| Bash       | 4        | 302      | 41      | 7 |
| Markdown   | 3        | 1105     | 204     | 0 |
| **Total**  | **34**   | **3406** | **616** | **373** |

## Docker

Modifique o GIN\_MODE caso necessário:

```bash
$ vim Dockerfile
ENV GIN_MODE=release
```

Inicialize os containers:

```bash
$ ./docker.sh <container> ...

# Inicia somente o gin_backend
$ ./docker.sh backend
# Inicia ambos os containers: postgresql e o gin_backend
$ ./docker.sh pg backend
```

O servidor estará esperando por conexões em 127.0.0.1:8000.

## Instruções

Instale as dependências:

```bash
$ go get .
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

Modifique o initializers/env.go para não utilizar o Docker:

```bash
$ vim initializers/env.go
var DATCOM_DOCKER bool = false
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

A documentação completa da API está disponível no arquivo [DOCUMENTATION-API.md](DOCUMENTATION-API.md). Ela contém informações detalhadas sobre o uso das endpoints (rotas), parâmetros e exemplos de requisições e respostas. Recomenda-se a leitura antes de iniciar o desenvolvimento.

## Segurança

Se você descobrir uma vulnerabilidade de segurança, consulte nossa [Política de Segurança](SECURITY.md) para mais detalhes.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use, modify, and distribute the code as needed. See the [LICENSE](LICENSE) file for more information.

## Image/Logo Credit

The coffee logo was generated by <a href="https://www.craiyon.com">Craiyon</a>.
