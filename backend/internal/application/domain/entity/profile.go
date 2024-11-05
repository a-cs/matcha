package entity

type ProfileIntention struct {
	Username string
	User     User
	Profile  Profile
}

type Profile struct {
	ID                 uint64
	UserID             uint64
	FirstName          string
	LastName           string
	Location           string
	LikesCounter       uint64
	GenderID           uint64
	TagsList           []Tags
	Biography          string
	SexualPreferenceID uint64
	Pictures           []Picture
	ViewCounter        uint64
	IsOnline           bool
	LastOnlineAt       string
	AccountStatus      string
}

type Tags struct {
	ID   uint64
	Name string
}

type Picture struct {
	ID  uint64
	Url string
}
