package contas

//ContaCorrente é a conta corrente
type ContaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Sacar(saque float64) string {

	podeSacar := c.Saldo > saque

	naoNegativo := saque > 0

	if podeSacar && naoNegativo {

		c.Saldo -= saque

		return "Saque realizado com sucesso"

	} else {

		return "Por algum motivo não foi possível realizar o saque"
	}
}

func (c *ContaCorrente) Depositar(deposito float64) string {

	naoNegativo := deposito > 0

	if naoNegativo {

		c.Saldo += deposito

		return "Deposito realizado com sucesso"
	} else {

		return "Por algum motivo não foi possível realizar o depósito"
	}

}

func (c *ContaCorrente) Transferencia(valor float64, conta *ContaCorrente) string {

	if c.Saldo > valor && valor > 0 {

		c.Saldo -= valor
		conta.Saldo += valor
		return "O valor foi depositado com sucesso"
	} else {

		return "O deposito não foi possível"
	}

}
