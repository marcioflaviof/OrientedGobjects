package main

import (
	"OrientedGobjects/contas"
	"fmt"
)

func main() {

	// contaAlfredo := contas.ContaCorrente{model.Titular{"Alfredo", "0852369741", "Cantor"}, 154, 666, 0}
	// contaJeanne := contas.ContaCorrente{model.Titular{"Jeanne", "063215568", "Tradutora"}, 154, 666, 0}

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(500)

	fmt.Println(contaExemplo.ObterSaldo())

}
