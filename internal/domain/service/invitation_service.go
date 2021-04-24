package service

import (
	"github.com/balloon/go/invite/internal/domain/model"
	"github.com/balloon/go/invite/internal/domain/usecase"
	"github.com/balloon/go/invite/internal/infrastructure/firebase"
)

type InvitationService interface {
	CreateInvitation(topicId string) (*model.Invitation, error)
	GetTopicId(code []int) (string, error)
}

type InvitationServiceImpl struct {
	usecase *usecase.InvitationUseCase
}

func NewInvitationService() (*InvitationServiceImpl, error) {
	r, err := firebase.NewFirestoreInvitationRepository()
	if err != nil {
		return nil, err
	}

	u := usecase.NewInvitationUseCase(r)
	s := &InvitationServiceImpl{
		u,
	}

	return s, nil
}

func (s *InvitationServiceImpl) CreateInvitation(topicId string) (*model.Invitation, error) {
	return s.usecase.CreateInvitation(topicId)
}

func (s *InvitationServiceImpl) GetTopicId(code []int) (string, error) {
	topicId, err := s.usecase.GetTopicId(code)
	if err != nil {
		return "", err
	}
	return string(topicId), err
}
