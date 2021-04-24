package invitation

import (
	"encoding/json"
	"github.com/balloon/go/invite/internal/domain/repository"
	"github.com/balloon/go/invite/internal/domain/service"
	"github.com/balloon/go/invite/internal/domain/usecase"
	"log"
	"net/http"
)

type GetTopicIdRequest struct {
	Code []int `json:"code"`
}

type GetTopicIdResponse struct {
	TopicId string `json:"topicId"`
}

func GetTopicId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Methods", "POST")
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

	var req GetTopicIdRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	topicId, err := invitationService.GetTopicId(req.Code)
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

	err = json.NewEncoder(w).Encode(GetTopicIdResponse{
		TopicId: topicId,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error while encoding response json:", err)
		return
	}
}
