package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/scortier/langtrans/cli"
)

var wg sync.WaitGroup

var sourceLang string
var targetLang string
var sourceText string

func init() {
	// Define command-line flags for source language, target language, and source text.
	flag.StringVar(&sourceLang, "source", "en", "Source language [en]")
	flag.StringVar(&targetLang, "target", "fr", "Target language [fr]")
	flag.StringVar(&sourceText, "text", "Hello, World!", "Text to translate")
}

func main() {
	fmt.Println("Language Translation CLI")

	// Parse command-line flags.
	flag.Parse()

	// Check if the user has set any flags. If not, display usage and exit.
	if flag.NFlag() == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Create a channel to receive the translated string.
	strChan := make(chan string)
	wg.Add(1)

	// Create a request body with user-provided data.
	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	// Start a goroutine to perform the translation.
	go cli.RequestTranslate(reqBody, strChan, &wg)

	// Receive and process the translated string.
	processedString := strings.ReplaceAll(<-strChan, "\n", "")

	// Display the source and translated text.
	fmt.Printf("Source: %s\nTarget: %s\n", sourceText, processedString)

	// Close the channel and wait for the translation to complete.
	close(strChan)
	wg.Wait()
}
