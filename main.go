package main

import (
	"OrientedGobjects/contas"
	"fmt"
)

func main() {

	contaAlfredo := contas.ContaCorrente{"Alfredo", 154, 666, 5.24}
	contaJeanne := contas.ContaCorrente{"Jeanne", 154, 666, 5524}

	fmt.Printf("%s, %.2f \n", contaAlfredo.Sacar(6), contaAlfredo.Saldo)

	fmt.Printf("%s, %.2f \n", contaJeanne.Depositar(6000), contaJeanne.Saldo)

	fmt.Printf("%s, primeiro Saldo %.2f, segundo Saldo %.2f\n", contaJeanne.Transferencia(2000, &contaAlfredo), contaJeanne.Saldo, contaAlfredo.Saldo)

}
