# official Go image
FROM golang:1.23.5-alpine

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o /src/main /src/main.go
EXPOSE 8000
ENV GIN_MODE=debug
CMD ["/src/main"]
