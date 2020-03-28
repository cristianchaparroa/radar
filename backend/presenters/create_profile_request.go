package presenters

type CreateProfileRequest struct {
	Profile  Profile  `json:"profile"`
	Location Location `json:"location"`
}
