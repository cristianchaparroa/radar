package presenters


type ProfileResponse struct {
	Profile Profile `json:"profile"`
	Status  Status  `json:"status"`
}
