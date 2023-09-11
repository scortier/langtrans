package cli

import (
	"log"
	"net/http"
	"sync"

	gabs "github.com/Jeffail/gabs/v2"
)

// RequestBody represents the request body for translation.
type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

// const URL = "https://translation.googleapis.com/language/translate/v2"
const URL = "https://translate.googleapis.com/translate_a/single"

// RequestTranslate sends a translation request to Google Translate API.
func RequestTranslate(body *RequestBody, strChan chan string, wg *sync.WaitGroup) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %s", err)
	}

	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)

	req.URL.RawQuery = query.Encode()

	if err != nil {
		log.Fatalf("Error encoding raw query: %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making HTTP request: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		log.Fatal("You have exceeded the rate limit")
		strChan <- "You have been rate limited, try again later"
		wg.Done()
		return
	}

	parsedJSON, err := gabs.ParseJSONBuffer(resp.Body)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}
	// fmt.Println(parsedJSON.String())

	nestedOne, err := parsedJSON.ArrayElement(0)
	if err != nil {
		log.Fatalf("Error getting array element 0: %s", err)
	}

	nestedTwo, err := nestedOne.ArrayElement(0)
	if err != nil {
		log.Fatalf("Error getting array element 0: %s", err)
	}

	translatedString, err := nestedTwo.ArrayElement(0)
	if err != nil {
		log.Fatalf("Error getting array element 0: %s", err)
	}

	strChan <- translatedString.Data().(string)
	wg.Done()
}
