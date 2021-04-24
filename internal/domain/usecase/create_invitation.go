package usecase

import (
	"errors"
	"fmt"
	"github.com/balloon/go/invite/internal/domain/model"
	"github.com/balloon/go/invite/internal/domain/repository"
)

var (
	InvitationAlreadyCreatedError = errors.New("invitation is already created")
	InvalidInvitationCodeError    = errors.New("invalid invitation code")
)

// CreateInvitation 一意の招待コードを持つ招待を作成
func (u *InvitationUseCase) CreateInvitation(topicId string) (*model.Invitation, error) {
	var invitation *model.Invitation

	// 指定されたTopicIdに対して、すでに招待が作成されていた場合、何もしない
	invitation, err := u.repository.FindByTopicId(model.TopicId(topicId))
	if err != nil && err != repository.InvitationNotFoundError {
		return nil, err
	}
	if invitation != nil {
		return nil, InvitationAlreadyCreatedError
	}

	// 招待コードが衝突しなくなるまで、招待を作成
	for {
		invitation = model.NewInvitation(topicId)
		_, err := u.repository.FindByInvitationCode(invitation.Code)
		if err != nil && err != repository.InvitationNotFoundError {
			return nil, err
		}
		break
	}

	err = u.repository.Save(invitation)
	if err != nil {
		return nil, fmt.Errorf("error while saveing new invitation: %v", err)
	}

	return invitation, nil
}
