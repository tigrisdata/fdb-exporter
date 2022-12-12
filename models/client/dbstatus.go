package clientmodel

type DatabaseStatus struct {
	Available bool `json:"available"`
	Healthy   bool `json:"healthy"`
}
