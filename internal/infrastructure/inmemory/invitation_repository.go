package inmemory

import (
	"github.com/balloon/go/invite/internal/domain/model"
	"github.com/balloon/go/invite/internal/domain/repository"
	"log"
	"sync"
)

type InMemoryInvitationRepository struct {
	lock        sync.RWMutex
	invitations []*model.Invitation
}

func NewInvitationRepository() (*InMemoryInvitationRepository, error) {
	invitations := []*model.Invitation{}
	return &InMemoryInvitationRepository{
		invitations: invitations,
	}, nil
}

func (r *InMemoryInvitationRepository) Save(invitation *model.Invitation) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.invitations = append(r.invitations, invitation)

	log.Println("New Invitation Added:", invitation)

	return nil
}

func (r *InMemoryInvitationRepository) FindByTopicId(topicId model.TopicId) (*model.Invitation, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, v := range r.invitations {
		if v.TopicId == topicId {
			return v, nil
		}
	}

	return nil, repository.InvitationNotFoundError
}

func (r *InMemoryInvitationRepository) FindByInvitationCode(code model.InvitationCode) (*model.Invitation, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, v := range r.invitations {
		if code.IsEqual(v.Code) {
			return v, nil
		}
	}

	return nil, repository.InvitationNotFoundError
}
