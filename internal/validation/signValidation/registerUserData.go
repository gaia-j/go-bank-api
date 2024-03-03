package signValidation

import (
	"fmt"
	"github.com/gaia-j/go-bank-api/internal/structs"
	"github.com/gaia-j/go-bank-api/internal/validation/validators"
	"strings"
)

func ValidateRegisterUserData(userData *structs.RegisterUser) error {

	if err := validators.Cpf(userData.Cpf); err != nil {
		return err
	}

	if userData.Name == "" {
		return fmt.Errorf("o campo 'Name' deve ser preenchido")
	}

	if len(userData.Name) < 3 {
		return fmt.Errorf("o campo 'Name' deve ter pelo menos 3 digitos")
	}

	if userData.Surname == "" {
		return fmt.Errorf("o campo 'Surname' deve ser preenchido")
	}

	if len(userData.Surname) < 3 {
		return fmt.Errorf("o campo 'Surname' deve ter pelo menos 3 digitos")
	}

	if userData.Password == "" {
		return fmt.Errorf("o campo 'Password' deve ser preenchido")
	}

	if len(userData.Password) < 8 {
		return fmt.Errorf("o campo 'Password' deve ter pelo menos 8 digitos")
	}

	if userData.Email == "" {
		return fmt.Errorf("o campo 'Email' deve ser preenchido")
	}

	if !strings.ContainsRune(userData.Email, '@') {
		return fmt.Errorf("insira um email válido")
	}

	if userData.Street == "" {
		return fmt.Errorf("o campo 'Street' deve ser preenchido")
	}

	if len(userData.Street) < 3 {
		return fmt.Errorf("o campo 'Street' deve ter pelo menos 3 digitos")
	}

	if userData.Number == "" {
		return fmt.Errorf("o campo 'Number' deve ser preenchido")
	}

	if len(userData.Number) < 1 {
		return fmt.Errorf("o campo 'Number' deve ter pelo menos 1 digito")
	}

	if userData.City == "" {
		return fmt.Errorf("o campo 'City' deve ser preenchido")
	}

	if userData.State == "" {
		return fmt.Errorf("o campo 'State' deve ser preenchido")
	}

	if userData.Neighbor == "" {
		return fmt.Errorf("o campo 'Neighbor' deve ser preenchido")
	}

	if userData.Cep == "" {
		return fmt.Errorf("o campo 'Cep' deve ser preenchido")
	}

	if len(userData.Cep) != 8 {
		return fmt.Errorf("o campo 'Cep' deve ter 8 digitos")
	}

	return nil
}
