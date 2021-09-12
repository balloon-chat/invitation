package main

import (
	"fmt"
	"github.com/balloon/go/invite/env"
	"github.com/balloon/go/invite/internal/interface/api/server/handler/invitation"
	"log"
	"net/http"
	"os"
)

func main() {
	if env.DEBUG && os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		log.Fatalln("Environment variable GOOGLE_APPLICATION_CREDENTIALS is empty")
	}

	http.HandleFunc("/create-invitation", invitation.CreateInvitation)
	http.HandleFunc("/invitation-topic", invitation.GetTopicId)
	http.HandleFunc("/invitation-code", invitation.GetInvitationCode)

	p := 8000
	log.Println(fmt.Sprintf("Listening at http://localhost:%d", p))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), nil); err != nil {
		log.Fatalln(err)
	}
}
