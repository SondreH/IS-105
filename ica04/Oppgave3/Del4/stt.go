package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
)

type jsonStruktur struct {
	StatusCode int `json:"StatusCode"`
	Headers    struct {
		Connection            []string `json:"Connection"`
		ContentDisposition    []string `json:"Content-Disposition"`
		ContentLength         []string `json:"Content-Length"`
		ContentSecurityPolicy []string `json:"Content-Security-Policy"`
		ContentType           []string `json:"Content-Type"`
		Date                  []string `json:"Date"`
		SessionName           []string `json:"Session-Name"`
		XContentTypeOptions   []string `json:"X-Content-Type-Options"`
		XDpWatsonTranID       []string `json:"X-Dp-Watson-Tran-Id"`
		XFrameOptions         []string `json:"X-Frame-Options"`
		XGlobalTransactionID  []string `json:"X-Global-Transaction-Id"`
		XXSSProtection        []string `json:"X-Xss-Protection"`
	} `json:"Headers"`
	Result struct {
		Results []struct {
			Final        bool `json:"final"`
			Alternatives []struct {
				Transcript string          `json:"transcript"`
				Confidence float64         `json:"confidence"`
				Timestamps [][]interface{} `json:"timestamps"`
			} `json:"alternatives"`
			WordAlternatives []struct {
				StartTime    float64 `json:"start_time"`
				EndTime      float64 `json:"end_time"`
				Alternatives []struct {
					Confidence float64 `json:"confidence"`
					Word       string  `json:"word"`
				} `json:"alternatives"`
			} `json:"word_alternatives"`
		} `json:"results"`
		ResultIndex int `json:"result_index"`
	} `json:"Result"`
}

var jsonUt jsonStruktur

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
			Keywords:                  []string{"Keyword"},
			KeywordsThreshold:         core.Float32Ptr(0.5),
		},
	)

	if responseErr != nil {
		log.Fatal("Failed getting the response", responseErr)
	}

	//json.Unmarshal([]byte(response.Result), jsonUt)
	fmt.Println(response)
	//fmt.Println(response)

}
