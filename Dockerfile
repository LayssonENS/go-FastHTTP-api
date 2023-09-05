# Use a imagem oficial do Golang como base
FROM golang:1.21 as builder

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie os arquivos do projeto para o container
COPY . .

# Compile o projeto
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./app/main.go

# Use uma imagem scratch para uma imagem final mais leve
FROM scratch

# Copie o binário compilado do builder para a imagem final
COPY --from=builder /app/main /main

# Exponha a porta 9000
EXPOSE 9000

# Defina o comando padrão para executar o binário
CMD ["/main"]
