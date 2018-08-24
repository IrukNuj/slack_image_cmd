package lib

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"strings"
	"net/http"
	"os"
	"google.golang.org/appengine"
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
	Queries struct {
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
			SearchType     string `json:"searchType"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
			SearchType     string `json:"searchType"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title string `json:"title"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []struct {
		Kind        string `json:"kind"`
		Title       string `json:"title"`
		HTMLTitle   string `json:"htmlTitle"`
		Link        string `json:"link"`
		DisplayLink string `json:"displayLink"`
		Snippet     string `json:"snippet"`
		HTMLSnippet string `json:"htmlSnippet"`
		Mime        string `json:"mime"`
		Image       struct {
			ContextLink     string `json:"contextLink"`
			Height          int    `json:"height"`
			Width           int    `json:"width"`
			ByteSize        int    `json:"byteSize"`
			ThumbnailLink   string `json:"thumbnailLink"`
			ThumbnailHeight int    `json:"thumbnailHeight"`
			ThumbnailWidth  int    `json:"thumbnailWidth"`
		} `json:"image"`
	} `json:"items"`
}

func CreateUrl(word string) string {
	baseUrl := "https://www.googleapis.com/customsearch/v1"
	s := Search{os.Getenv("CUSTOM_SEARCH_KEY"), os.Getenv("CUSTOM_SEARCH_ENGINE_ID"), "image", "1"}
	word = strings.TrimSpace(word)
	url := baseUrl + "?key=" + s.Key + "&cx=" + s.EngineId + "&searchType=" + s.Type + "&num=" + s.Count + "&q=" + word
	return url
}

func SearchImage(r *http.Request, text string) string {
	var imageUrl = "not search image"
	ctx := appengine.NewContext(r)
	httpClient := urlfetch.Client(ctx)

	resp, err := httpClient.Get(CreateUrl(text))
	if err != nil {
		log.Fatal(ctx, "html: %v", err)
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
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
