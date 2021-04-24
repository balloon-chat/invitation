package usecase

import (
	"github.com/balloon/go/invite/internal/domain/model"
	"reflect"
	"testing"
)

func TestCreateInvitation(t *testing.T) {
	r := testInvitationRepository{}
	u := NewInvitationUseCase(r)

	t.Run("generate unique invitation code", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			topicId := string(rune(i))
			invitation, err := u.CreateInvitation(topicId)
			if err != nil {
				t.Error(err)
			}
			if invitation.TopicId != model.TopicId(topicId) {
				t.Errorf("Invitation#TopicId is wrong: Expected %s, Acctual %s", topicId, invitation.TopicId)
			}

			sameCode := 0
			for _, inv := range invitations {
				if reflect.DeepEqual(inv.Code, invitation.Code) {
					sameCode++
				}
			}
			if sameCode > 1 {
				t.Errorf("Generated same invitation code: %v", invitation.Code)
			}
		}
	})

	t.Run("don't create invitation of same topic", func(t *testing.T) {
		i := model.NewInvitation("test")
		_ = r.Save(i)
		_, err := u.CreateInvitation(string(i.TopicId))
		if err != InvitationAlreadyCreatedError {
			t.Error("Same invitation of topic is created.")
		}
	})
}
