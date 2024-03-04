package database

import (
	"context"
	"math/rand"
)

func generateAccountAddress() (string, error) {

	elegibleLetters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	elegibleNumbers := []byte("0123456789")

	var generatedAddressBytes []byte

	var generatedAddress string

	for i := 0; i < 9; i++ {

		randomLetter := rand.Intn(len(elegibleLetters))
		randomNumber := rand.Intn(len(elegibleNumbers))

		if i < 3 || i > 7 {
			generatedAddressBytes = append(generatedAddressBytes, elegibleLetters[randomLetter])
		} else {
			generatedAddressBytes = append(generatedAddressBytes, elegibleNumbers[randomNumber])
		}
	}

	generatedAddress = string(generatedAddressBytes)

	return generatedAddress, nil

}

func GenerateUniqueAccountAddress() (string, error) {

	accountAddress, err := generateAccountAddress()

	if err != nil {
		return "", err
	}

	var accountAddressCount int

	errGot := Db.QueryRow(context.Background(), "SELECT COUNT(*) FROM account WHERE account_address = $1", accountAddress).Scan(&accountAddressCount)

	if errGot != nil {
		return "", errGot

	}

	if accountAddressCount > 0 {
		return GenerateUniqueAccountAddress()
	}

	return accountAddress, nil

}
