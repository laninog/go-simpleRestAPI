package models

type User struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstName,omitempty" validate:"nonzero"`
	LastName  string   `json:"lastName,omitempty" validate:"nonzero"`
	Address   *Address `json:"address,omitempty" validate:"nonzero"`
}