package main

import (
	"net/http"
	"github.com/nlopes/slack"
	"encoding/json"
	"github.com/JinOketani/slack_image_cmd/lib"
	"os"
)

func main() {
	http.HandleFunc("/cmd", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(os.Getenv("VERIFICATION_TOKEN")) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/image":
		response := &slack.Msg{Text: lib.SearchImage(s.Text), ResponseType: "in_channel"}
		image, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(image)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
