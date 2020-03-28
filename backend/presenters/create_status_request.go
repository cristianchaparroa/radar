package presenters

type CreateStatus struct {
	Status Status `json:"status"`
	Location Location `json:"location"`
}
