package entity

type PasswordIntention struct {
	Email          string
	User           User
	JwtObj         JwtObj
	NewPassword    string
	HashedPassword string
	JwtToken       string
}
