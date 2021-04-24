package model

import (
	"math/rand"
	"time"
)

const (
	CodeLength = 8
)

type InvitationCode []int

type TopicId string

type Invitation struct {
	Code InvitationCode

	TopicId TopicId

	CreatedAt time.Time
}

func NewInvitation(topicId string) *Invitation {
	return &Invitation{
		Code:      generateInvitationCode(CodeLength),
		TopicId:   TopicId(topicId),
		CreatedAt: time.Now(),
	}
}

// generateInvitationCode 指定された長さで、0~9の値の数列を作成する。
func generateInvitationCode(length int) []int {
	code := make([]int, length)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		code[i] = rand.Intn(10)
	}

	return code
}

func (a InvitationCode) IsEqual(b InvitationCode) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
