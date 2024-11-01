package entity

type JwtObj struct {
	UserID   uint64
	Username string
	Iat      int64
}

type JwtResponse struct {
	Token         string
	ProfileStatus string
	UserID        uint64
}
