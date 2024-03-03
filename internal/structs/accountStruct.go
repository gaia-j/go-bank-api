package structs

type CreateAccount struct {
	Balance        int64
	AccountAddress string
	BalanceBtc     int64
	Credit         int64
	UserId         int
}

type Account struct {
	Id         int
	Balance    int64
	BalanceBtc int64
	Credit     int64
	UserId     int
	CreatedAt  string
	UpdatedAt  string
}
