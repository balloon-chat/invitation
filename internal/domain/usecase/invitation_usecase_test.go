package usecase

import (
	"github.com/balloon/go/invite/internal/domain/model"
	"github.com/balloon/go/invite/internal/domain/repository"
	"reflect"
)

var invitations []*model.Invitation

type testInvitationRepository struct {
}

func (t testInvitationRepository) Save(invitation *model.Invitation) error {
	invitations = append(invitations, invitation)
	return nil
}

func (t testInvitationRepository) FindByTopicId(topicId model.TopicId) (*model.Invitation, error) {
	for _, invitation := range invitations {
		if invitation.TopicId == topicId {
			return invitation, nil
		}
	}
	return nil, repository.InvitationNotFoundError
}

func (t testInvitationRepository) FindByInvitationCode(code model.InvitationCode) (*model.Invitation, error) {
	for _, invitation := range invitations {
		if reflect.DeepEqual(invitation.Code, code) {
			return invitation, nil
		}
	}
	return nil, repository.InvitationNotFoundError
}

func (t testInvitationRepository) Reset() {
	invitations = []*model.Invitation{}
}
