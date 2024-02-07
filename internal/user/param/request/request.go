package request

type CreateUser struct {
	Username    string
	PhoneNumber string
	Name        string
	Password    string
	Gender      string
}

type GenerateUserToken struct {
	Username string
	Password string
}
