package models

import "encoding/json"

type Type struct {
	Type string `json:"type"` // "User" or "Note"
}

type Ap struct {
	Type
	Object any `json:"object"` // models.User or models.Note
}

func (a *Ap) UnmarshalJSON(data []byte) error {
	var t Type
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	switch t.Type {
	case "User":
		a.Object = new(User)
	case "Note":
		a.Object = new(Note)
	}
	return json.Unmarshal(data, &a.Object)
}
