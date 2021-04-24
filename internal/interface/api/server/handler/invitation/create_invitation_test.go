package invitation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateInvitation(t *testing.T) {
	invitationService = newTestInvitationService()
	server := httptest.NewServer(http.HandlerFunc(CreateInvitation))
	defer server.Close()

	t.Run("create new invitation", func(t *testing.T) {
		req := CreateInvitationRequest{TopicId: "test"}
		bs, _ := json.Marshal(req)
		res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(bs))
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("the status code is not %d, but %d", http.StatusOK, res.StatusCode)
		}

		var body CreateInvitationResponse
		err = json.NewDecoder(res.Body).Decode(&body)
		if err != nil {
			t.Errorf("cannot decode response: %v", err)
		}

		if body.TopicId != req.TopicId {
			t.Errorf("the wrong topic's invitation was created")
		}
	})

	t.Run("create invitation with same topic id", func(t *testing.T) {
		req := CreateInvitationRequest{TopicId: "test"}
		bs, _ := json.Marshal(req)
		_, _ = http.Post(server.URL, "application/json", bytes.NewBuffer(bs))
		res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(bs))
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("the status code is not %d, but %d", http.StatusBadRequest, res.StatusCode)
		}
	})
}
