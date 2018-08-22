package main

import (
	"net/http"
	"github.com/nlopes/slack"
	"fmt"
	"encoding/json"
	"github.com/JinOketani/slack_image_cmd/lib"
	"os"
)

func main() {
	http.HandleFunc("/cmd", func(w http.ResponseWriter, r *http.Request) {
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !s.ValidateToken(os.Getenv("VERIFICATION_TOKEN")) {
			fmt.Println(s)
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
	})
	fmt.Println("[Info] Server start")
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
