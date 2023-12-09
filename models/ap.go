package models

import (
	"encoding/json"

	"github.com/yitsushi/go-misskey/core"
)

// Ap represents an Ap from /ap endpoints.
type Ap struct {
	Type   string `json:"type"`   // "User" or "Note"
	Object any    `json:"object"` // models.User or models.Note
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It will coerce the "object" field into the correct type (User/Note).
func (a *Ap) UnmarshalJSON(data []byte) error {
	var t struct {
		Type string `json:"type"` // "User" or "Note"
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	a.Type = t.Type

	switch t.Type {
	case "User":
		var u struct {
			Object User `json:"object"`
		}

		if err := json.Unmarshal(data, &u); err != nil {
			return err
		}

		a.Object = u.Object

		return nil
	case "Note":
		var n struct {
			Object Note `json:"object"`
		}

		if err := json.Unmarshal(data, &n); err != nil {
			return err
		}

		a.Object = n.Object

		return nil
	default:
		return core.NotImplementedYet{Reason: t.Type + " is not User/Note"}
	}
}
