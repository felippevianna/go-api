FROM golang:1.25.1 AS builder

WORKDIR /app
COPY . .
# comando para instalar as dependencias do projeto
RUN go mod download 
RUN go build -o main .

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]


# Buildar e rodar a aplicaão
# docker build -t go-api .
# docker run -p 8080:8080 go-api

# Exemplo de comando para instalar uma dependencia
# docker run --rm -v ${PWD}:/app -w /app golang:1.25.1 go get github.com/gorilla/mux

