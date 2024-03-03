package misc

import "math/rand"

func generateAccountAddress() string {

	elegibleLetters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	elegibleNumbers := []byte("0123456789")

	var generatedAddress []byte

	for i := 0; i < 9; i++ {

		randomLetter := rand.Intn(len(elegibleLetters))
		randomNumber := rand.Intn(len(elegibleNumbers))

		if i < 3 || i > 7 {
			generatedAddress = append(generatedAddress, elegibleLetters[randomLetter])
		} else {
			generatedAddress = append(generatedAddress, elegibleNumbers[randomNumber])
		}
	}

	return string(generatedAddress)

}
