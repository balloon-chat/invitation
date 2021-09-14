package main

import (
	"fmt"
	"github.com/balloon/go/invite/internal/interface/api/server/handler/invitation"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-invitation", invitation.CreateInvitation)
	http.HandleFunc("/invitation-topic", invitation.GetTopicId)
	http.HandleFunc("/invitation-code", invitation.GetInvitationCode)

	p := 8000
	log.Println(fmt.Sprintf("Listening at http://localhost:%d", p))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), nil); err != nil {
		log.Fatalln(err)
	}
}
