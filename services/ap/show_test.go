package ap_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/ap"
	"github.com/yitsushi/go-misskey/test"
)

// TestService_Show tests the Show method of the Ap service.
func TestService_Show(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/ap/show",
		RequestData:  &ap.ShowRequest{},
		ResponseFile: "note.json",
		StatusCode:   http.StatusOK,
	})

	resp, err := client.Ap().Show("8y7eu4z91l")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Note", resp.Type)

	note, ok := resp.Object.(models.Note)
	assert.True(t, ok, "The type assertion should be possible as it's a Note type")
	assert.Equal(t, "9n1hunzmvlv20008", note.ID)
	assert.Equal(t, "88v9vu5nbu", note.UserID)
}

// TestService_ShowNote tests the ShowNote method of the Ap service.
func TestService_ShowNote(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/ap/show",
		RequestData:  &ap.ShowRequest{},
		ResponseFile: "note.json",
		StatusCode:   http.StatusOK,
	})

	note, err := client.Ap().ShowNote("8y7eu4z91l")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "agdfg98g8g7ggg", note.ID)
	assert.Equal(t, "pzbbeer89g", note.UserID)
}

// TestService_ShowUser tests the ShowUser method of the Ap service.
func TestService_ShowUser(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/ap/show",
		RequestData:  &ap.ShowRequest{},
		ResponseFile: "user.json",
		StatusCode:   http.StatusOK,
	})

	user, err := client.Ap().ShowUser("8y7eu4z91l")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "88v9vu5nbu", user.ID)
	assert.Equal(t, "kiki_test", user.Username)
	assert.True(t, user.IsBot)
}

// TestShowRequest_Validate tests the Validate method of the ShowRequest.
func TestShowRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			ap.ShowRequest{},
		},
		[]core.BaseRequest{
			ap.ShowRequest{URI: "asd"},
		},
	)
}

// ExampleService_Show demonstrates how to use Ap.Show.
func ExampleService_Show() {
	client, _ := misskey.NewClientWithOptions(
		misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")),
	)

	resp, err := client.Ap().Show("8y7eu4z91l")
	if err != nil {
		log.Printf("[Ap/Show] %s", err)

		return
	}

	// You can use the type assertion to determine the type of the object.
	switch object := resp.Object.(type) {
	case models.User:
		log.Printf("[Ap/Show] %s", object.Username)
	case models.Note:
		log.Printf("[Ap/Show] %s", object.Text)
	}
}

// ExampleService_ShowNote demonstrates how to use Ap.ShowNote.
func ExampleService_ShowNote() {
	client, _ := misskey.NewClientWithOptions(
		misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")),
	)

	resp, err := client.Ap().ShowNote("8y7eu4z91l")
	if err != nil {
		log.Printf("[Ap/ShowNote] %s", err)

		return
	}

	log.Printf("[Ap/Show] %s", resp.Text)
}

// ExampleService_ShowUser demonstrates how to use Ap.ShowUser.
func ExampleService_ShowUser() {
	client, _ := misskey.NewClientWithOptions(
		misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")),
	)

	resp, err := client.Ap().ShowUser("8y7eu4z91l")
	if err != nil {
		log.Printf("[Ap/Show] %s", err)

		return
	}

	log.Printf("[Ap/Show] %s", resp.Username)
}
