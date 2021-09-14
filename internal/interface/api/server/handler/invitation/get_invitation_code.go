package invitation

import (
	"encoding/json"
	"github.com/balloon/go/invite/env"
	"github.com/balloon/go/invite/internal/domain/repository"
	"github.com/balloon/go/invite/internal/domain/service"
	"github.com/balloon/go/invite/internal/domain/usecase"
	"log"
	"net/http"
)

type GetInvitationCodeRequest struct {
	TopicId string `json:"topicId"`
}

type GetInvitationCodeResponse struct {
	Code []int `json:"code"`
}

func GetInvitationCode(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", env.ClientEntryPoint)
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case http.MethodPost:
		break
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
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

	var req GetInvitationCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	code, err := invitationService.GetInvitationCode(req.TopicId)
	if err == usecase.InvalidInvitationCodeError {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err == repository.InvitationNotFoundError {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println("error while getting invitation:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(GetInvitationCodeResponse{
		Code: code,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error while encoding response json:", err)
		return
	}
}
