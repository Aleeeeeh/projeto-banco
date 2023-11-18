package contas

import "banco/clientes"

/*
public -> Começa com letra maiuscula
private -> Começa com letra minuscula
Com isso private só é possível acessar de fato a propriedade dentro do mesmo arquivo, conforme
exemplo da propriedade saldo
*/
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

// c *ContaCorrente faz o método ser referência da struct ContaCorrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor >= 0 {
		c.saldo += valor
		return "Depósito realizado com sucesso no valor de", valor
	}

	return "Valor do depósito menor que zero, valor informador de", valor
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor < c.saldo && valor > 0 {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	}

	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
