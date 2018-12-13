package models

type Address struct {
	City  string `json:"city,omitempty" validate:"nonzero"`
	State string `json:"state,omitempty" validate:"nonzero"`
}