package entity

//Portfolio entity
type Portfolio struct {
	Name    string    `json:"name"`
	Actives []*Active `json:"actives"`
}

//Active entity
type Active struct {
	Code       string       `json:"code"`
	Operations []*Operation `json:"operations"`
}

//Operation entity
type Operation struct {
	Value    float32 `json:"value"`
	Quantity int     `json:"quantity"`
}
