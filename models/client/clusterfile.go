package clientmodel

type ClusterFile struct {
	Path     string `json:"path"`
	UpToDate bool   `json:"up_to_date"`
}
