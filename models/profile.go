package models

// Profile contains all player profile data
type Profile struct {
	FirstName string
	LastName  string
	Has18     bool
	Blocked   bool
}

// IsFull indicates if profile has all fields filled up
func (p *Profile) IsFull() bool {
	return p.FirstName != "" && p.LastName != "" && p.Has18
}
