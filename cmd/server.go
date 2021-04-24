package main

import (
	"github.com/balloon/go/invite/internal/interface/api/server/handler/invitation"
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		log.Fatalln("Environment variable GOOGLE_APPLICATION_CREDENTIALS is empty")
	}

	http.HandleFunc("/create", invitation.CreateInvitation)
	http.HandleFunc("/get", invitation.GetTopicId)
	log.Println("Listening at http://localhost:8080 ")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
