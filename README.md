# [**_FullCycle_**] Desafio - Client-Server-API

## Como Rodar

```
docker compose up --build
go run client/client.go
```

## Problema
Você precisará nos entregar dois sistemas em Go:
- client.go
- server.go

## ToDo - Requisitos
- [X] O `server.go` deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
- [X] O `client.go` deverá realizar uma requisição HTTP no `server.go` solicitando a cotação do dólar.
- [X] O `client.go` precisará receber do `server.go` apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package **"context"**, 
  - [X] o `client.go` terá um timeout máximo de 300ms para receber o resultado do `server.go`.
- [X] Usando o package **"context"**, o `server.go` deverá registrar no banco de dados **SQLite** cada cotação recebida, sendo que:
  - [X] o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e;
  - [X] o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
- [X] Os 3 contextos _deverão retornar erro nos logs_ caso o tempo de execução seja insuficiente.
- [X] O `client.go` terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
**O endpoint necessário gerado pelo `server.go` para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.**