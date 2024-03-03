package signValidation

import (
	"fmt"
	"github.com/gaia-j/go-bank-api/internal/structs"
	"strings"
)

func ValidateLoginUserData(userData *structs.LoginUser) error {

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

	return nil

}
