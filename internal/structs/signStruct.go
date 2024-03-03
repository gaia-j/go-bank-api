package structs

type RegisterAdmin struct {
	User     string
	Password string
	Cpf      string
}
type RegisterUser struct {
	Name     string
	Surname  string
	Cpf      string
	UserId   string
	Password string
	Email    string
	Street   string
	Neighbor string
	Cep      string
	Number   string
	Comp     string
	State    string
	City     string
}

type LoginUser struct {
	Email    string
	Password string
}
