# Transfers API

Docker
```
docker build . -t transfer-server

docker-compose up -d
```

Comandos para rodar os testes:
```
go test ./...

go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
