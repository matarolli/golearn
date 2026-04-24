// Define que este arquivo pertence ao pacote principal, necessario para criar um programa executavel.
package main

// Importa o pacote fmt, usado para ler dados da entrada e imprimir resultados na tela.
import "fmt"

// Funcao principal: e o ponto de inicio da execucao do programa.
func main() {
	// Declara a variavel kelvin, do tipo inteiro, para armazenar o valor lido em Kelvin.
	var kelvin int

	// Le um numero digitado pelo usuario e guarda esse valor na variavel kelvin.
	fmt.Scan(&kelvin)

	// Converte o valor de Kelvin para Celsius usando a formula C = K - 273.
	celsius := kelvin - 273

	// Imprime na tela o valor convertido para graus Celsius.
	fmt.Println(celsius)
}
