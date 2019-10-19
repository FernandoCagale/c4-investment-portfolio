package entity

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Portfolio struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Actives []*Active `json:"actives"`
}

func (e Portfolio) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Name, validation.Required, validation.Length(2, 50)),
	)
}

type Active struct {
	Code       string       `json:"code"`
	Operations []*Operation `json:"operations"`
}

type Operation struct {
	Value    float32 `json:"value"`
	Quantity int     `json:"quantity"`
}
