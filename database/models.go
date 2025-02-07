package database

type Director struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}

type Actor struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}

type Film struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	DirectedBy int    `json:"directedBy"`
	Logline    string `json:"logline"`
	Year       int    `json:"year"`
}

type Character struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PortrayedBy  int    `json:"portrayedBy"`
	FeaturedIn   int    `json:"featuredIn"`
	DiesInTheEnd bool   `json:"diesInTheEnd"`
}
