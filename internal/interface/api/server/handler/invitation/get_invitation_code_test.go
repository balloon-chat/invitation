package invitation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetInvitationCode(t *testing.T) {
	service := newTestInvitationService()
	invitationService = service
	server := httptest.NewServer(http.HandlerFunc(GetInvitationCode))
	defer server.Close()

	t.Run("get invitation code", func(t *testing.T) {
		service.Repository.Reset()
		invitation, _ := invitationService.CreateInvitation("test")

		req := GetInvitationCodeRequest{TopicId: string(invitation.TopicId)}
		bs, _ := json.Marshal(req)
		res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(bs))
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("the status code is not %d, but %d", http.StatusOK, res.StatusCode)
		}

		var body GetInvitationCodeResponse
		err = json.NewDecoder(res.Body).Decode(&body)
		if err != nil {
			t.Errorf("cannot decode response: %v", err)
		}

		if !invitation.Code.IsEqual(body.Code) {
			t.Errorf("Expected: %v, Actual: %v", invitation.Code, body.Code)
		}
	})

	t.Run("get invitation code before creating a invitation", func(t *testing.T) {
		service.Repository.Reset()
		req := GetInvitationCodeRequest{TopicId: "test"}
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
