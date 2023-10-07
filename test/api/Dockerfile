# stage de build
FROM golang:1.21.1 AS build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/main.go

# stage imagem final
FROM scratch
WORKDIR /app
COPY --from=build /app ./
EXPOSE 3000
CMD [ "./api", "env=docker" ]