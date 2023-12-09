package ap

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest represents an Show request.
type ShowRequest struct {
	URI string `json:"uri"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	if r.URI == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "URI",
		}
	}

	return nil
}

// Show ap will return either a note or a user object.
func (s *Service) Show(uri string) (models.Ap, error) {
	var response models.Ap
	err := s.Call(
		&core.JSONRequest{Request: &ShowRequest{URI: uri}, Path: "/ap/show"},
		&response,
	)

	return response, err
}

// ShowUser will return a user object if the ap/show object type is User.
func (s *Service) ShowUser(uri string) (models.User, error) {
	apResult, err := s.Show(uri)

	user, ok := apResult.Object.(models.User)
	if !ok {
		return user, core.InvalidFieldReferenceError{Name: "ShowUser", Type: "", Reference: apResult.Type}
	}

	return user, err
}

// ShowNote will return a user object if the ap/show object type is Note.
func (s *Service) ShowNote(uri string) (models.Note, error) {
	apResult, err := s.Show(uri)

	note, ok := apResult.Object.(models.Note)
	if !ok {
		return note, core.InvalidFieldReferenceError{Name: "ShowUser", Type: "", Reference: apResult.Type}
	}

	return note, err
}
