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
			Field:   "AppID",
		}
	}

	return nil
}

// Show app.
func (s *Service) Show(uri string) (models.Ap, error) {
	var response any
	err := s.Call(
		&core.JSONRequest{Request: &ShowRequest{URI: uri}, Path: "/ap/show"},
		&response,
	)

	return response, err
}
