package usecase

import (
	"github.com/balloon/go/invite/internal/domain/model"
	"testing"
)

func TestInvitationUseCase_GetInvitationCode(t *testing.T) {
	r := testInvitationRepository{}
	u := NewInvitationUseCase(r)

	t.Run("get invitation code", func(t *testing.T) {
		r.Reset()
		invitation := model.NewInvitation("test")
		_ = r.Save(invitation)

		code, err := u.GetInvitationCode(invitation.TopicId)
		if err != nil {
			t.Error(err)
		}

		if !code.IsEqual(invitation.Code) {
			t.Errorf("Expected: %v, Actual: %v", invitation.Code, code)
		}
	})

	t.Run("get invitation code before save", func(t *testing.T) {
		r.Reset()
		invitation := model.NewInvitation("test")
		_, err := u.GetInvitationCode(invitation.TopicId)
		if err == nil {
			t.Errorf("did not return error")
		}
	})
}
