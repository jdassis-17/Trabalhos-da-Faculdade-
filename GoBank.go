package main

import (
	"errors"
	"fmt"
)

type Conta struct {
	Titular string
	Numero  int
	Saldo   float64
}

func (c *Conta) Depositar(valor float64) error {
	if valor <= 0 {
		return errors.New("Quantia a ser depositada deve ser maior que zero")
	}
	c.Saldo += valor
	fmt.Printf("Depósito de R$ %.2f realizado com sucesso!\n", valor)
	return nil
}

func (c *Conta) Sacar(valor float64) error {
	if valor <= 0 {
		return errors.New("Quantia de saque deve ser maior que zero")
	}
	if c.Saldo < valor {
		return errors.New("Saldo insuficiente para realizar o saque")
	}
	c.Saldo -= valor
	fmt.Printf("Saque de R$ %.2f realizado com sucesso!\n", valor)
	return nil
}

func main() {
	minhaConta := Conta{Titular: "João", Numero: 407806, Saldo: 2400.0}

	fmt.Printf("Saldo inicial de %s (Conta %d): R$ %.2f\n", minhaConta.Titular, minhaConta.Numero, minhaConta.Saldo)

	err := minhaConta.Depositar(321.50)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Printf("Saldo atualizado: R$ %.2f\n", minhaConta.Saldo)

	err = minhaConta.Sacar(1800.0)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Printf("Saldo atualizado: R$ %.2f\n", minhaConta.Saldo)

	err = minhaConta.Sacar(300.0)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	err = minhaConta.Depositar(-520.0)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Printf("Saldo atualizado: R$ %.2f\n", minhaConta.Saldo)

}
