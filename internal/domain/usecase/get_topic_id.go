package usecase

import (
	"github.com/balloon/go/invite/internal/domain/model"
)

// GetTopicId 招待コードから対応する話題のIDを取得
func (u *InvitationUseCase) GetTopicId(code []int) (model.TopicId, error) {
	if len(code) != model.CodeLength {
		return "", InvalidInvitationCodeError
	}

	invitation, err := u.repository.FindByInvitationCode(code)
	if err != nil {
		return "", err
	}

	return invitation.TopicId, nil
}
