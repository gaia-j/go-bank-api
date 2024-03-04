package database

import (
	"context"
	"github.com/gaia-j/go-bank-api/internal/structs"
)

func CreateUserAccount(userId int) error {

	AccountAddress, errAddress := GenerateUniqueAccountAddress()

	if errAddress != nil {
		return errAddress
	}

	AccountData := structs.CreateAccount{
		Balance:        0,
		AccountAddress: AccountAddress,
		BalanceBtc:     0,
		Credit:         0,
		UserId:         userId,
	}

	_, err := Db.Exec(context.Background(), `
	INSERT INTO account (user_id, account_address, balance, balance_btc, credit)
	VALUES ($1, $2, $3, $4, $5)
	`,
		AccountData.UserId,
		AccountData.AccountAddress,
		AccountData.Balance,
		AccountData.BalanceBtc,
		AccountData.Credit,
	)

	if err != nil {
		return err
	}

	return nil

}
