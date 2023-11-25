package entity

import (
	"regexp"
	"testing"
)

func TestTamanhoCpf(t *testing.T) {
	cpf := "14538220620"
	resultado := len(cpf)
	esperado := 11
	if esperado != resultado {
		t.Errorf("esperado %d, resultado %d", esperado, resultado)
	}
}

func TestRegexCpf(t *testing.T) {
	cpf := "14538220620"
	pattern := regexp.MustCompile("[0-9]+")
	resultado := pattern.FindAllString(cpf, -1)

	if cpf != resultado[0] {
		t.Errorf("esperado %s, resultado %s", cpf, resultado[0])
	}
}

func TestPeso(t *testing.T) {
	cpf := "14538220620"
	peso1, peso2 := SomarPesos(cpf)
	esperado1 := 185
	esperado2 := 220
	if peso1 != esperado1 {
		t.Errorf("esperado1 %d, resuldado1 %d", esperado1, peso1)
	}

	if peso2 != esperado2 {
		t.Errorf("esperado2 %d, resuldado2 %d", esperado2, peso2)
	}
}

func TestCalcularDigito(t *testing.T) {
	peso1 := 185
	peso2 := 220
	cpf := "14538220620"

	esperado1, esperado2 := SepararDigitoDoCpfInformado(cpf)
	digito1, digito2 := CalcularDigito(peso1, peso2)

	if digito1 != esperado1 {
		t.Errorf("esperado1 %s , resultado1 %s", esperado1, digito1)
	}

	if digito2 != esperado2 {
		t.Errorf("esperado2 %s , resultado2 %s", esperado2, digito2)
	}
}

func TestValidarCpf(t *testing.T) {
	cpf := "15776285860"
	peso1, peso2 := SomarPesos(cpf)
	digito1, digito2 := CalcularDigito(peso1, peso2)
	esperado1, esperado2 := SepararDigitoDoCpfInformado(cpf)
	if digito1 != esperado1 {
		t.Errorf("esperado1 %s , resultado1 %s", esperado1, digito1)
	}
	if digito2 != esperado2 {
		t.Errorf("esperado2 %s , resultado2 %s", esperado2, digito2)
	}
}

func TestCpf(t *testing.T) {

	validaCpf := func(t *testing.T, digitoEsperado1, digitoEsperado2, digito1Calculado1, digitoCalculado2 string) {
		t.Helper()
		if digitoEsperado1 != digito1Calculado1 || digitoEsperado2 != digitoCalculado2 {
			t.Errorf("Primeiro Digito %s != %s , Segundo Digito %s != %s", digitoEsperado1, digito1Calculado1, digitoEsperado2, digitoCalculado2)
		}
	}

	t.Run("validar cpf", func(t *testing.T) {
		cpf := "81499852045"
		digitoEsperado1, digitoEsperado2 := SepararDigitoDoCpfInformado(cpf)
		peso1, peso2 := SomarPesos(cpf)
		digito1Calculado1, digitoCalculado2 := CalcularDigito(peso1, peso2)
		validaCpf(t, digitoEsperado1, digitoEsperado2, digito1Calculado1, digitoCalculado2)
	})

	t.Run("validar cpf", func(t *testing.T) {
		cpf := "46885462039"
		digitoEsperado1, digitoEsperado2 := SepararDigitoDoCpfInformado(cpf)
		peso1, peso2 := SomarPesos(cpf)
		digito1Calculado1, digitoCalculado2 := CalcularDigito(peso1, peso2)
		validaCpf(t, digitoEsperado1, digitoEsperado2, digito1Calculado1, digitoCalculado2)
	})

	t.Run("validar cpf", func(t *testing.T) {
		cpf := "13830579071"
		digitoEsperado1, digitoEsperado2 := SepararDigitoDoCpfInformado(cpf)
		peso1, peso2 := SomarPesos(cpf)
		digito1Calculado1, digitoCalculado2 := CalcularDigito(peso1, peso2)
		validaCpf(t, digitoEsperado1, digitoEsperado2, digito1Calculado1, digitoCalculado2)
	})
}

func TestArrayCpf(t *testing.T) {

	validaCpf := func(t *testing.T, cpf string) {
		t.Helper()

		err := ValidarCpf(cpf)

		if err != nil {
			t.Errorf("CPF inválido, %s", err)
		}

		digitoEsperado1, digitoEsperado2 := SepararDigitoDoCpfInformado(cpf)
		peso1, peso2 := SomarPesos(cpf)
		digito1Calculado1, digitoCalculado2 := CalcularDigito(peso1, peso2)

		if digitoEsperado1 != digito1Calculado1 || digitoEsperado2 != digitoCalculado2 {
			t.Errorf("Primeiro Digito %s != %s , Segundo Digito %s != %s", digitoEsperado1, digito1Calculado1, digitoEsperado2, digitoCalculado2)
		}
	}

	t.Run("validar cpf", func(t *testing.T) {
		arrayCpf := []string{"28222527002", "60344902005", "41072820072", "55085950070", "59896444005", "65108853017"}
		for _, cpf := range arrayCpf {
			validaCpf(t, cpf)
		}
	})

}
