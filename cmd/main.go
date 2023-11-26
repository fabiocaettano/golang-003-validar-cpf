package main

import (
	"fmt"

	CPF "github.com/fabiocaettano74/validar-cpf/internal/cpf/entity"
)

func main() {
	arrayCpf := []string{"51365179052", "51365179051", "60344902005", "60344902050", "41072820072", "41072820027", "55085950070", "59896444005", "65108853017"}

	for _, cpf := range arrayCpf {
		result := CPF.ValidarDigito(cpf)
		if result {
			fmt.Printf("CPF %s é válido", cpf)
		} else {
			fmt.Printf("CPF %s não é valido", cpf)
		}
		fmt.Printf("\n")
	}
}
