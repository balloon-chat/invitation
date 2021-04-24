package usecase

import (
	"github.com/balloon/go/invite/internal/domain/repository"
)

type InvitationUseCase struct {
	repository repository.InvitationRepository
}

func NewInvitationUseCase(r repository.InvitationRepository) *InvitationUseCase {
	return &InvitationUseCase{
		repository: r,
	}
}
