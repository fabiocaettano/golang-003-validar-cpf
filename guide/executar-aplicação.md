## Executar aplicação

1. Subir o container:

``` bash
docker-compose up -d --build
```

2. Acessar o container: 

``` bash
docker-compose exec goapp bash
```

3. Executar os testes em todos os diretórios:

``` bash
go test ./...
```

4. Validr uma lista de CPFs:

``` bash
go run cmd/main.go
```