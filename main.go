package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteGopher := clientes.Titular{Nome: "Gopher", Cpf: "123.456.789.10", Profissao: "Dev"}
	contaDoGopher := contas.ContaCorrente{Titular: clienteGopher, NumeroAgencia: 6882, NumeroConta: 55669}

	var deposito float64
	fmt.Println("Digite o valor para depósito: ")
	fmt.Scan(&deposito)
	status, valor := contaDoGopher.Depositar(deposito)
	fmt.Println(status, valor)
	fmt.Println("Saldo em conta do", contaDoGopher.Titular.Nome, ":", contaDoGopher.ObterSaldo())
	fmt.Println()

	var saque float64
	fmt.Println("Digite o valor do saque: ")
	fmt.Scan(&saque)
	fmt.Println(contaDoGopher.Sacar(saque))
	fmt.Println("Saldo atualizado da conta do", contaDoGopher.Titular.Nome, ":", contaDoGopher.ObterSaldo())
	fmt.Println()

	clienteGopherzinho := clientes.Titular{Nome: "Gopherzinho", Cpf: "111.222.333.44", Profissao: "devinho"}
	contaGopherzinho := contas.ContaCorrente{Titular: clienteGopherzinho, NumeroAgencia: 6882, NumeroConta: 5559874}

	var valorTransferencia float64
	fmt.Println("Insira valor da transferência para conta do", contaGopherzinho.Titular.Nome)
	fmt.Scan(&valorTransferencia)
	transferencia := contaDoGopher.Transferir(valorTransferencia, &contaGopherzinho)
	fmt.Println("Resposta da transação:", transferencia)
	fmt.Println(contaDoGopher)
	fmt.Println(contaGopherzinho)
	fmt.Println("Obrigado e volte sempre !")
}
