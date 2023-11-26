package entity

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type CPF struct {
	numero string
}

func Validar(input string) error {

	cpf := CPF{
		numero: input,
	}

	if cpf.numero == "" {
		return errors.New("Informar o CPF")
	}

	if len(cpf.numero) != 11 {
		return errors.New("CPF deve conter 11 digitos")
	}

	pattern := regexp.MustCompile("[0-9]+")
	resultado := pattern.FindAllString(cpf.numero, -1)

	if cpf.numero != resultado[0] {
		return errors.New("CPF deve conter somente números")
	}

	return nil
}

func ValidarDigito(cpf string) bool {

	err := Validar(cpf)
	if err != nil {
		panic(err)
	}
	digitoEsperado1, digitoEsperado2 := SepararDigito(cpf)
	peso1, peso2 := SomarPesos(cpf)
	digito1Calculado1, digitoCalculado2 := CalcularDigito(peso1, peso2)
	if digitoEsperado1 != digito1Calculado1 || digitoEsperado2 != digitoCalculado2 {
		return false
	}
	return true
}

func SepararDigito(cpf string) (string, string) {
	digito := cpf[9:]
	arrayDigito := strings.Split(digito, "")
	digito1 := arrayDigito[0]
	digito2 := arrayDigito[1]
	return digito1, digito2
}

func CalcularDigito(peso1 int, peso2 int) (string, string) {
	resto1 := 11 - (peso1 % 11)
	resto2 := 11 - (peso2 % 11)
	if resto1 >= 10 {
		resto1 = 0
	}
	if resto2 >= 10 {
		resto2 = 0
	}
	return strconv.Itoa(resto1), strconv.Itoa(resto2)
}

func SomarPesos(cpf string) (int, int) {
	soma1 := 0
	contador1 := 0
	arrayCpf1 := strings.Split(cpf[:9], "")
	arrayMultiplicadores1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}

	for _, numero := range arrayCpf1 {
		x, err := strconv.Atoi(numero)
		if err != nil {
			panic(err)
		}
		soma1 += x * arrayMultiplicadores1[contador1]
		contador1 += 1
	}

	soma2 := 0
	contador2 := 0
	arrayCpf2 := strings.Split(cpf[:10], "")
	arrayMultiplicadores2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	for _, numero := range arrayCpf2 {
		x, err := strconv.Atoi(numero)
		if err != nil {
			panic(err)
		}
		soma2 += x * arrayMultiplicadores2[contador2]
		contador2 += 1
	}
	return soma1, soma2
}
