package contas

import "OrientedGobjects/model"

type ContaPoupanca struct {
	Titular  model.Titular `json:"titular"`
	Agencia  int           `json:"agencia"`
	Conta    int           `json:"conta"`
	Operacao int           `json:"operacao"`
	saldo    float64       `json:"saldo"`
}

// Sacar = Withdraw
func (c *ContaPoupanca) Sacar(saque float64) string {

	podeSacar := c.saldo > saque

	naoNegativo := saque > 0

	if podeSacar && naoNegativo {

		c.saldo -= saque

		return "Saque realizado com sucesso"

	} else {

		return "Por algum motivo não foi possível realizar o saque"
	}
}

// Depositar = Deposit
func (c *ContaPoupanca) Depositar(deposito float64) string {

	naoNegativo := deposito > 0

	if naoNegativo {

		c.saldo += deposito

		return "Deposito realizado com sucesso"

	} else {

		return "Por algum motivo não foi possível realizar o depósito"
	}

}

// Transferencia = transfer
func (c *ContaPoupanca) Transferencia(valor float64, conta *ContaPoupanca) string {

	if c.saldo > valor && valor > 0 {

		c.saldo -= valor
		conta.saldo += valor
		return "O valor foi depositado com sucesso"
	} else {

		return "O deposito não foi possível"
	}

}

// ObterSaldo Check how much money you have
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
