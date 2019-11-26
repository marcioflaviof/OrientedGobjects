package main

import "fmt"

type contaCorrente struct {
	titular  string
	nAgencia int
	nConta   int
	saldo    float64
}

func (c *contaCorrente) sacar(saque float64) string {

	podeSacar := c.saldo > saque
	naoNegativo := saque > 0

	if podeSacar && naoNegativo {

		c.saldo -= saque

		return "Saque realizado com sucesso"

	} else {

		return "Por algum motivo não foi possível realizar o saque"
	}

}

func main() {

	contaAlfredo := contaCorrente{"Alfredo", 154, 666, 5.24}

	fmt.Printf("%s, %.2f \n", contaAlfredo.sacar(6), contaAlfredo.saldo)

}
