package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

// c *ContaCorrente faz o método ser referência da struct ContaCorrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor >= 0 {
		c.Saldo += valor
		return "Depósito realizado com sucesso, valor de", valor
	}

	return "Valor do depósito menor que zero, valor informador de", valor
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor < c.Saldo && valor > 0 {
		c.Saldo -= valor
		contaDestino.Depositar(valor)
		return true
	}

	return false
}
