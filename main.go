package main

import (
	"OrientedGobjects/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {

	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	// contaAlfredo := contas.ContaCorrente{model.Titular{"Alfredo", "0852369741", "Cantor"}, 154, 666, 0}
	// contaJeanne := contas.ContaCorrente{model.Titular{"Jeanne", "063215568", "Tradutora"}, 154, 666, 0}

	contaExemplo := contas.ContaPoupanca{}
	contaExemplo.Depositar(500)
	PagarBoleto(&contaExemplo, 70)

	fmt.Println(contaExemplo)

}
