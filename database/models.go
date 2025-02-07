package database

type Director struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName" validate:"required"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName" validate:"required"`
}

type Actor struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName" validate:"required"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName" validate:"required"`
}

type Film struct {
	ID         int    `json:"id"`
	Title      string `json:"title" validate:"required"`
	DirectedBy int    `json:"directedBy" validate:"required"`
	Logline    string `json:"logline" validate:"required"`
	Year       int    `json:"year" validate:"required,min=1900,max=2040"`
}

type Character struct {
	ID           int    `json:"id"`
	Name         string `json:"name" validate:"required"`
	PortrayedBy  int    `json:"portrayedBy" validate:"required"`
	FeaturedIn   int    `json:"featuredIn" validate:"required"`
	DiesInTheEnd bool   `json:"diesInTheEnd" validate:"boolean"`
}
