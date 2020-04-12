package domain

// CurrentProfile contains the las state of all information of
// one profile.
type CurrentProfile struct {
	Profile Profile
	Status Status;
	Location Location;
}