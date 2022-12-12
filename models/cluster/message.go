package clustermodel

type Message struct {
	Description string  `json:"description"`
	Issues      []Issue `json:"issues"`
	Name        string  `json:"name"`
}

type Issue struct {
	Addresses   []string `json:"addresses"`
	Count       int      `json:"count"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
}
