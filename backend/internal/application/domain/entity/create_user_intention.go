package entity

type CreateUserIntention struct {
	CreateUser     CreateUser
	User           User
	Profile        Profile
	HashedPassword string
	SlugID         string
}
