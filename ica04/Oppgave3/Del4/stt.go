package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
)

func main() {

	speechToText, speechToTextErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
		IAMApiKey: "gLQV9wfEpXPRSaayEdbadTQv-fxAjH9NoN2F6oLytmOu",
		URL:       "https://gateway-lon.watsonplatform.net/speech-to-text/api",
	})

	if speechToTextErr != nil {
		panic(speechToTextErr)
	}

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}

	var audioFile io.ReadCloser
	audioFile, audioFileErr := os.Open(os.Args[1])
	if audioFileErr != nil {
		log.Fatal("Failed reading file", audioFileErr)
	}

	response, responseErr := speechToText.Recognize(
		&speechtotextv1.RecognizeOptions{
			Audio:                     &audioFile,
			ContentType:               core.StringPtr(speechtotextv1.RecognizeOptions_ContentType_AudioWav),
			Timestamps:                core.BoolPtr(true),
			WordAlternativesThreshold: core.Float32Ptr(0.9),
			Keywords:                  []string{"colorado", "tornado", "tornadoes"},
			KeywordsThreshold:         core.Float32Ptr(0.5),
		},
	)
	if responseErr != nil {
		log.Fatal("Failed getting the response", responseErr)
	}

	result := speechToText.GetRecognizeResult(response)
	b, err := json.MarshalIndent(result, "", "  ")
	// Dette er bare en test jeg la til fordi jeg ikke forstår hvor null kommer.
	if err != nil {
		log.Fatal("Json error", err)
	}
	fmt.Println(string(b))

}
