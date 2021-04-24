package invitation

import (
	"bytes"
	"encoding/json"
	"github.com/balloon/go/invite/internal/domain/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTopicId(t *testing.T) {
	invitationService = newTestInvitationService()
	server := httptest.NewServer(http.HandlerFunc(GetTopicId))
	defer server.Close()

	t.Run("get topic id", func(t *testing.T) {
		invitation, _ := invitationService.CreateInvitation("test")

		req := GetTopicIdRequest{Code: invitation.Code}
		bs, _ := json.Marshal(req)
		res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(bs))
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("the status code is not %d, but %d", http.StatusOK, res.StatusCode)
		}

		var body GetTopicIdResponse
		err = json.NewDecoder(res.Body).Decode(&body)
		if err != nil {
			t.Errorf("cannot decode response: %v", err)
		}

		if body.TopicId != string(invitation.TopicId) {
			t.Errorf("Expected: %s, Actual: %s", invitation.TopicId, body.TopicId)
		}
	})

	t.Run("get topic id with invalid invitation code", func(t *testing.T) {
		req := GetTopicIdRequest{Code: []int{1, 2, 3}}
		bs, _ := json.Marshal(req)
		res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(bs))
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("the status code is not %d, but %d", http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("get topic id before creating a invitation", func(t *testing.T) {
		code := make([]int, model.CodeLength)
		req := GetTopicIdRequest{Code: code}
		bs, _ := json.Marshal(req)
		res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(bs))
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("the status code is not %d, but %d", http.StatusBadRequest, res.StatusCode)
		}
	})
}
