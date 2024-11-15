package entity

type LoginIntention struct {
	Email         string
	Password      string
	User          User
	ProfileStatus string
	JwtResponse   JwtResponse
}
