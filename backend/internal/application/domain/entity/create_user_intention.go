package entity

type CreateUserIntention struct {
	User           CreateUser
	HashedPassword string
	SlugID         string
}
