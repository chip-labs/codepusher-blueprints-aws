# Etapa 1: Build
FROM golang:1.24-alpine AS builder

# Define o diretório de trabalho para o build
WORKDIR /app

# Copia os arquivos de go.mod e go.sum (para cache de dependências)
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod download

# Copia todo o restante do código
COPY . .

# Compila o projeto
RUN go build -o fiber-app ./cmd/server

# Etapa 2: Final (imagem leve)
FROM alpine:latest

# Define o diretório de trabalho na imagem final
WORKDIR /app

# Copia o binário compilado da etapa de build
COPY --from=builder /app/fiber-app .

# Expõe a porta que a aplicação usa (3000, por exemplo)
EXPOSE 3000

# Comando que roda a aplicação
ENTRYPOINT ["./fiber-app"]
