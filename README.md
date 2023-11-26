# golang-003-validar-cpf

Esta aplicação desenvolvida em GoLang verifica se um número de CPF é valido.

## Regra de Negócios

CPF utilizado no exemplo: 51365179052

1. Selecionar os 09 e 10 primeirios digitos do CPF, para multiplicar cada algarismos conforme abaixo:

| Soma 1      | Soma 2       |
|-------------|--------------|
| 5 x 10 = 50 | 5 x 11 = 55  |    
| 1 x 09 = 9  | 1 x 10 = 10  |
| 3 x 08 = 24 | 3 x 09 = 27  | 
| 6 x 07 = 42 | 6 x 08 = 48  | 
| 5 x 06 = 30 | 5 x 07 = 35  | 
| 1 x 05 = 05 | 1 x 06 = 06  |
| 7 x 04 = 28 | 7 x 05 = 35  | 
| 9 x 03 = 27 | 9 x 04 = 36  | 
| 0 x 02 = 00 | 0 x 03 = 00  |
|             | 5 x 02 = 10  |  
| Total: 215  | Total: 262   |

2. O próximo passo é descobrir o resto da divisão.

| Resto 1              | Resto 2               |
------------------------------------------------
| 215 / 11 = 19,545454 | 262 / 11 = 23,818181  |
| 215 - (19 * 11) = 6  | 262 - (23 * 11) = 9   |


3. Agora subtrair 11 do resto da divisão, para descobrir o digito verificador.

| Penúltimo Digito | Último Digito   |
|------------------|------------------
| 11 - 6 = 5       | 11 - 9 = 2      | 


4. Resto da divisão maior ou igual a 10.

Caso o resto da divisão for maior ou igual a 10, o resto deverá ser igual a zero.


## Executar aplicação

Subir o container:

``` bash
docker-compose up -d --build
```

Acessar o container: 

``` bash
docker-compose exec goapp bash
```

Executar os testes em todos os diretórios:

``` bash
go test ./...
```

Validr uma lista de CPFs:

``` bash
go run cmd/main.go
```

## Referências




