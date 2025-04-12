# Desafio PosGoExpert - Stress Test
Sistema CLI em Go que realiza testes de carga em um serviço web

## Objetivo
Permitir que o usuário forneça:
- Uma **URL**
- O número total de **requisições**
- A quantidade de chamadas **simultâneas (concorrência)**
E ao final, gerar um relatório detalhado da execução do teste.

## Parâmetros da CLI
- **--url**	URL do serviço a ser testado
- **--requests**	Número total de requisições a serem feitas
- **--concurrency**	Número de chamadas simultâneas

## Build do Docker
```
docker build -t stress_test .
```

## Executando o teste
```
docker run stress_test --url=http://teste.com --requests=1000 --concurrency=10
```

## Exemplo de saída (relatório)
```
Relatório Final:
Tempo total: 406.464553ms
Total de requests: 1000
Requests com status 200: 990
Distribuição dos códigos de status:
Status 200: 990
Status 500: 10
```
