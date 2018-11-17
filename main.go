package slack_image_cmd

import (
	"net/http"
	"github.com/nlopes/slack"
	"os"
	"encoding/json"
	"google.golang.org/appengine"
	"strings"
	"log"
	"io/ioutil"
	"google.golang.org/appengine/urlfetch"
)

type Search struct {
	Key      string
	EngineId string
	Type     string
	Count    string
}

type Result struct {
	Kind string `json:"kind"`
	URL  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Items []struct {
		Link        string `json:"link"`
	} `json:"items"`
}

func Init() {
	http.HandleFunc("/cmd", handler)
	appengine.Main()
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
		response := &slack.Msg{Text: SearchImage(r, s.Text), ResponseType: "in_channel"}
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

func SearchImage(r *http.Request, word string) string {
	baseUrl := "https://www.googleapis.com/customsearch/v1"
	s := Search{os.Getenv("CUSTOM_SEARCH_KEY"), os.Getenv("CUSTOM_SEARCH_ENGINE_ID"), "image", "1"}
	word = strings.TrimSpace(word)
	url := baseUrl + "?key=" + s.Key + "&cx=" + s.EngineId + "&searchType=" + s.Type + "&num=" + s.Count + "&q=" + word
	return ParseJson(r, url)
}

func ParseJson(r *http.Request, url string) string {
	var imageUrl = "not search image"
	ctx := appengine.NewContext(r)
	httpClient := urlfetch.Client(ctx)

	response, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(ctx, "html: %v", err)
	}
	if response != nil {
		defer response.Body.Close()
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes := ([]byte)(body)
	data := new(Result)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		log.Println("json error:", err)
	}

	if data.Items != nil {
		imageUrl = data.Items[0].Link
	}

	return imageUrl
}
