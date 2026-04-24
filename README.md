# gokelvin

## Problema

Um professor de ensino medio solicitou aos seus alunos que desenvolvessem um programa para converter o valor do ponto de ebulicao da agua de Kelvin para graus Celsius.

Formula informada:

```text
C = K - 273
```

Onde:

- `C` representa a temperatura em graus Celsius.
- `K` representa a temperatura em Kelvin.

## Entrada

O programa recebe um numero inteiro representando a temperatura em Kelvin.

Exemplo:

```text
373
```

## Saida

O programa imprime o valor convertido para graus Celsius.

Exemplo:

```text
100
```

## Solucao implementada em Go

Arquivo: `convert.go`

```go
package main

import "fmt"

func main() {
	var kelvin int

	fmt.Scan(&kelvin)

	celsius := kelvin - 273
	fmt.Println(celsius)
}
```

## Como executar

```bash
go run convert.go
```
