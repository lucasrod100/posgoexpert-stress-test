# Etapa de build
FROM golang:1.23 AS builder

WORKDIR /app

# Copia os arquivos do projeto
#COPY go.mod go.sum ./
#RUN go mod download
COPY . .

RUN go mod tidy

# Compila o binário de forma estática para evitar problemas de glibc
RUN CGO_ENABLED=0 go build -o /app/stress_test ./cmd/main.go && chmod +x /app/stress_test
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /app/stress_test ./cmd/main.go && chmod +x /app/stress_test

# Etapa final (imagem mínima para execução)
FROM debian:bullseye-slim

WORKDIR /root/

# 🛠️ Instala os certificados de CA para HTTPS funcionar
RUN apt-get update && apt-get install -y ca-certificates

# Copia o binário da fase de build
COPY --from=builder /app/stress_test /stress_test

# Define o comando padrão
ENTRYPOINT ["/stress_test"]