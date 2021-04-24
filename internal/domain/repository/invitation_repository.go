package repository

import (
	"errors"
	"github.com/balloon/go/invite/internal/domain/model"
)

var (
	InvitationNotFoundError = errors.New("invitation not found")
)

type InvitationRepository interface {
	Save(invitation *model.Invitation) error

	// FindByTopicId TopicIdにより招待を探す。 もし、招待が存在しない場合は InvitationNotFoundError を返す
	FindByTopicId(topicId model.TopicId) (*model.Invitation, error)

	// FindByTopicId 招待コードにより招待を探す。 もし、招待が存在しない場合は InvitationNotFoundError を返す
	FindByInvitationCode(code model.InvitationCode) (*model.Invitation, error)
}
