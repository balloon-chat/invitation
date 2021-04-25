package usecase

import "github.com/balloon/go/invite/internal/domain/model"

// GetInvitationCode 話題の招待コードを取得
func (u *InvitationUseCase) GetInvitationCode(topicId model.TopicId) (model.InvitationCode, error) {
	invitation, err := u.repository.FindByTopicId(topicId)
	if err != nil {
		return nil, err
	}

	return invitation.Code, nil
}
