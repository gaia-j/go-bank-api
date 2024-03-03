package database

import (
	"github.com/gaia-j/go-bank-api/internal/structs"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(
	registerUserData structs.RegisterUser) (int, error) {

	var createdUserId int

	var hashedPassword []byte

	bytePassword := []byte(registerUserData.Password)

	hashedPassword, errPass := bcrypt.GenerateFromPassword(bytePassword, 10)

	if errPass != nil {
		return 0, errPass
	}

	err := Db.QueryRow(
		"INSERT INTO users(name, surname, cpf, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		registerUserData.Name, registerUserData.Surname, registerUserData.Cpf, registerUserData.Email,
		string(hashedPassword)).
		Scan(&createdUserId)

	if err != nil {
		return 0, err
	}
	return createdUserId, nil
}
