package database

import (
	"context"
	"github.com/gaia-j/go-bank-api/internal/structs"
)

func CreateAddress(userId int, userData structs.RegisterUser) error {
	AddressData := structs.CreateAddress{
		UserId:   userId,
		Street:   userData.Street,
		Number:   userData.Number,
		Neighbor: userData.Neighbor,
		Comp:     userData.Comp,
		City:     userData.City,
		State:    userData.State,
		Cep:      userData.Cep,
	}

	_, err := Db.Exec(context.Background(), `
	INSERT INTO address (user_id, street, number, comp, city, state, cep, neighbor) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`,
		AddressData.UserId, AddressData.Street, AddressData.Number,
		AddressData.Comp, AddressData.City, AddressData.State, AddressData.Cep, AddressData.Neighbor)

	if err != nil {
		return err
	}

	return nil

}
