package validators

import "fmt"

func Cpf(cpf string) error {

	if cpf == "" {
		return fmt.Errorf("o campo 'cpf' deve ser preenchido")
	}

	if len(cpf) != 11 {
		return fmt.Errorf("o campo 'cpf' deve ter 11 digitos")
	}

	allEqual := true

	for i := 0; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			allEqual = false
			break
		}
	}

	if allEqual {
		return fmt.Errorf("todos os digítos do cpf são iguais")
	}

	bytesCpf := []byte(cpf)

	firstDigit := int(bytesCpf[9] - '0')
	secondDigit := int(bytesCpf[10] - '0')

	firstDigitVerify := 0
	secondDigitVerify := 0

	for i := 0; i <= 9; i++ {
		digit := int(bytesCpf[i] - '0')
		multiply := 10 - i
		if i < 9 {
			firstDigitVerify += digit * multiply
		}
		secondDigitVerify += digit * (multiply + 1)
	}

	firstDigitRes := (firstDigitVerify * 10) % 11

	secondDigitRes := (secondDigitVerify * 10) % 11

	if firstDigitRes == 10 {
		firstDigitRes = 0
	}

	if secondDigitRes == 10 {
		secondDigitRes = 0
	}

	if firstDigit != firstDigitRes || secondDigit != secondDigitRes {
		return fmt.Errorf("cpf inválido")
	}

	return nil

}
