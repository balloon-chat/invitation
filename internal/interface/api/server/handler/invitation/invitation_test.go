package invitation

import (
	"github.com/balloon/go/invite/internal/domain/model"
	"github.com/balloon/go/invite/internal/domain/usecase"
	"github.com/balloon/go/invite/internal/infrastructure/inmemory"
)

type testInvitationService struct {
	u *usecase.InvitationUseCase
}

func newTestInvitationService() *testInvitationService {
	r, _ := inmemory.NewInvitationRepository()
	u := usecase.NewInvitationUseCase(r)
	return &testInvitationService{
		u: u,
	}
}

func (s *testInvitationService) CreateInvitation(topicId string) (*model.Invitation, error) {
	return s.u.CreateInvitation(topicId)
}

func (s *testInvitationService) GetTopicId(code []int) (string, error) {
	topicId, err := s.u.GetTopicId(code)
	return string(topicId), err
}
