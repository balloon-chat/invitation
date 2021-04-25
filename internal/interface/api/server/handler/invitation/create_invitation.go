package invitation

import (
	"encoding/json"
	"github.com/balloon/go/invite/internal/domain/service"
	"github.com/balloon/go/invite/internal/domain/usecase"
	"log"
	"net/http"
)

type CreateInvitationRequest struct {
	TopicId string `json:"topicId"`
}

type CreateInvitationResponse struct {
	TopicId string `json:"topicId"`
	Code    []int  `json:"code"`
}

func CreateInvitation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if invitationService == nil {
		s, err := service.NewInvitationService()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		invitationService = s
	}

	var req CreateInvitationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	invitation, err := invitationService.CreateInvitation(req.TopicId)
	if err == usecase.InvitationAlreadyCreatedError {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println("error while creating invitation:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(CreateInvitationResponse{
		TopicId: string(invitation.TopicId),
		Code:    invitation.Code,
	})
	if err != nil {
		log.Println("error while encoding response json:", err)
		return
	}
}
