package lib

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"strings"
	"net/http"
	"os"
)

type Search struct {
	Key      string
	EngineId string
	Type     string
	Count    string
}

type Result struct {
	Kind string `json:"kind"`
	URL struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Items []struct {
		Link        string `json:"link"`
	} `json:"items"`
}

func SearchImage(word string) string {
	baseUrl := "https://www.googleapis.com/customsearch/v1"
	s := Search{os.Getenv("CUSTOM_SEARCH_KEY"), os.Getenv("CUSTOM_SEARCH_ENGINE_ID"), "image", "1"}
	word = strings.TrimSpace(word)
	url := baseUrl + "?key=" + s.Key + "&cx=" + s.EngineId + "&searchType=" + s.Type + "&num=" + s.Count + "&q=" + word
	return ParseJson(url)
}

func ParseJson(url string) string {
	var imageUrl = "not search image"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("get error:", err)
	}
	if response != nil {
		defer response.Body.Close()
	}

	byteArray, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes := ([]byte)(byteArray)
	data := new(Result)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		log.Println("json error:", err)
	}
	if data.Items != nil {
		imageUrl = data.Items[0].Link
		log.Println(imageUrl)
	}

	return imageUrl
}
