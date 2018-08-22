package slack_image_cmd

import (
	"net/http"
	"fmt"
	"github.com/nlopes/slack"
	"os"
	"encoding/json"
	"google.golang.org/appengine"
	"github.com/JinOketani/slack_image_cmd/lib"
)

func init() {
	http.HandleFunc("/cmd", handler)
	appengine.Main()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi!")
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
}
