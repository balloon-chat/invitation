package invite

import (
	"github.com/balloon/go/invite/internal/interface/api/server/handler/invitation"
	"net/http"
)

func CreateInvitation(w http.ResponseWriter, r *http.Request) {
	invitation.CreateInvitation(w, r)
}

func GetTopicId(w http.ResponseWriter, r *http.Request) {
	invitation.GetTopicId(w, r)
}
