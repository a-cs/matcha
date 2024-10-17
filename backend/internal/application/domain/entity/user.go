package entity

type CreateUser struct {
	Email     string
	Username  string
	Password  string
	FirstName string
	LastName  string
}

type User struct {
	ID            uint64
	Email         string
	Password      string
	Username      string
	ActiveMatches interface{}
	AccountStatus string
	SlugID        string
}
