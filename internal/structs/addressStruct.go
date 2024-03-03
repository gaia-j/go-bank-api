package structs

type CreateAddress struct {
	UserId   int
	Street   string
	Number   string
	Comp     string
	Neighbor string
	City     string
	State    string
	Cep      string
}

type Address struct {
	Id        int
	UserId    int
	Street    string
	Number    string
	Comp      string
	Neighbor  string
	City      string
	State     string
	Cep       string
	CreatedAt string
	UpdatedAt string
}
