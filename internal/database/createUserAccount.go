package database

import (
	"github.com/gaia-j/go-bank-api/internal/structs"
)

func CreateUserAccount(userId int) error {

	AccoutData := structs.CreateAccount{
		Balance:    0,
		BalanceBtc: 0,
		Credit:     0,
		UserId:     userId,
	}

	_, err := Db.Exec(`
	INSERT INTO account (users_id, balance, balance_btc, credit) 
	VALUES ($1, $2, $3, $4)
	`,
		AccoutData.UserId,
		AccoutData.Balance,
		AccoutData.BalanceBtc,
		AccoutData.Credit,
	)

	if err != nil {
		return err
	}

	return nil

}
