package main

import "fmt"

type contaCorrente struct {
	titular  string
	nAgencia int
	nConta   int
	saldo    float64
}

func main() {

	contaAlfredo := contaCorrente{"Alfredo", 154, 666, 5.24}

	fmt.Println(contaAlfredo)

}
