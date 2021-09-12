package invite

import (
	"github.com/balloon/go/invite/env"
	"github.com/balloon/go/invite/internal/interface/api/server/handler/invitation"
	"net/http"
)

func init() {
	env.LoadEnv()
}

func CreateInvitation(w http.ResponseWriter, r *http.Request) {
	invitation.CreateInvitation(w, r)
}

func GetTopicId(w http.ResponseWriter, r *http.Request) {
	invitation.GetTopicId(w, r)
}

func GetInvitationCode(w http.ResponseWriter, r *http.Request) {
	invitation.GetInvitationCode(w, r)
}
