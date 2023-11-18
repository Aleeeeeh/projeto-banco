package interfaces

type direcionaConta interface {
	/*Ambas as contas implementa esse método "Sacar". Em Go basta ter um método com esse nome
	sem a necessidade das classes usar a palavra reservada "implements", bastando apenas ter
	o método com esse nome em cada classe
	*/
	Sacar(valor float64) string
}

func PagarBoleto(conta direcionaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}
