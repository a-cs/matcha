package entity

type Profile struct {
	ID                 uint64
	UserID             uint64
	FirstName          string
	LastName           string
	Location           string
	LikesCounter       uint64
	GenderID           uint64
	TagsList           interface{}
	Biography          string
	SexualPreferenceID uint64
	Pictures           interface{}
	ViewCounter        uint64
	IsOnline           bool
	LastOnlineAt       string
	AccountStatus      string
}
