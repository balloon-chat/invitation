package repository

import (
	"github.com/balloon/go/invite/internal/domain/model"
	"github.com/balloon/go/invite/internal/domain/util"
	"time"
)

const (
	KeyCode      = "code"
	KeyTopicId   = "topicId"
	KeyCreatedAt = "createdAt"
)

type InvitationEntity struct {
	Code      int       `json:"code"`
	TopicId   string    `json:"topicId"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewInvitationEntity(invitation model.Invitation) *InvitationEntity {
	return &InvitationEntity{
		Code:      util.ArrayToInt(invitation.Code),
		TopicId:   string(invitation.TopicId),
		CreatedAt: invitation.CreatedAt,
	}
}

func (e InvitationEntity) ToMap() map[string]interface{} {
	return map[string]interface{}{
		KeyCode:      e.Code,
		KeyTopicId:   e.TopicId,
		KeyCreatedAt: e.CreatedAt,
	}
}

func (e InvitationEntity) ToInvitation() *model.Invitation {
	return &model.Invitation{
		Code:      util.IntToArray(e.Code, model.CodeLength),
		TopicId:   model.TopicId(e.TopicId),
		CreatedAt: e.CreatedAt,
	}
}
