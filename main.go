package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoGopher := ContaCorrente{}
	contaDoGopher.saldo = 500
	fmt.Println("Saldo em conta:", contaDoGopher.saldo)
	var saque float64
	fmt.Println("Digite o valor do saque: ")
	fmt.Scan(&saque)

	fmt.Println(contaDoGopher.Sacar(saque))
	fmt.Println("Saldo atualizado:", contaDoGopher.saldo)
}

// c *ContaCorrente faz o método ser referência da struct ContaCorrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}
