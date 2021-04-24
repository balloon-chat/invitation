package usecase

import (
	"github.com/balloon/go/invite/internal/domain/model"
	"testing"
)

func TestInvitationUseCase_GetTopicId(t *testing.T) {
	r := testInvitationRepository{}
	u := NewInvitationUseCase(r)

	t.Run("get topic id", func(t *testing.T) {
		invitation := model.NewInvitation("test")
		_ = r.Save(invitation)

		topicId, err := u.GetTopicId(invitation.Code)
		if err != nil {
			t.Error(err)
		}

		if topicId != invitation.TopicId {
			t.Errorf("Expected: %s, Actual: %s", invitation.TopicId, topicId)
		}
	})

	t.Run("get topic id with invalid invitation code", func(t *testing.T) {
		code := []int{1, 2, 3}
		_, err := u.GetTopicId(code)

		if err != InvalidInvitationCodeError {
			t.Errorf("returned wrong error: %v", err)
		}
	})
}
