package entity

type LoginIntention struct {
	Email       string
	Password    string
	User        User
	Profile     Profile
	JwtResponse JwtResponse
}
