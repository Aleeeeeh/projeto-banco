package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoGopher := contas.ContaCorrente{Titular: "Gopher"}

	var deposito float64
	fmt.Println("Digite o valor para depósito: ")
	fmt.Scan(&deposito)
	status, valor := contaDoGopher.Depositar(deposito)
	fmt.Println(status, valor)

	fmt.Println("Saldo em conta do gopher:", contaDoGopher.Saldo)
	var saque float64
	fmt.Println("Digite o valor do saque: ")
	fmt.Scan(&saque)
	fmt.Println(contaDoGopher.Sacar(saque))
	fmt.Println("Saldo atualizado da conta do gopher:", contaDoGopher.Saldo)

	contaGopherzinho := contas.ContaCorrente{Titular: "Gopherzinho"}

	var valorTransferencia float64
	fmt.Println("Insira valor da transferência para conta do gopherzinho: ")
	fmt.Scan(&valorTransferencia)
	transferencia := contaDoGopher.Transferir(valorTransferencia, &contaGopherzinho)
	fmt.Println("Resposta da transação:", transferencia)
	fmt.Println(contaDoGopher)
	fmt.Println(contaGopherzinho)
	fmt.Println("Obrigado e volte sempre !")
}
