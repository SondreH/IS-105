package main

import (
	"encoding/json"
	"fmt"
	"log"

	speech "github.com/SondreH/IS-105/ica04/Oppgave3/Del2/go-speak"
)

// Eksperimentering:
type jsonStruktur struct {
	SpeakingText string `json:"_text"`
	Outcomes     []struct {
		Confidence interface{} `json:"confidence"`
		Intent     string      `json:"intent"`
		Text       string      `json:"_text"`
		Entities   struct {
		} `json:"entities"`
	} `json:"outcomes"`
	WARNING string `json:"WARNING"`
	MsgID   string `json:"msg_id"`
}

var jsonUt jsonStruktur

func main() {
	speech.SetWitKey("CEG4C6LHPECM7AFEG7VMP5VGZVN55IJL") //Wit API Key MUST be set before calling any other Wit.AI functions
	jsonData := []byte(speech.SendWitVoice("ClearSample.wav"))
	//jsonData := (speech.SendWitVoice("ClearSample.wav"))

	/*dekodet := json.NewDecoder(strings.NewReader(jsonData))
	if err := dec.Decode(&verdier); err != nil	{
		log.Fatalf("oops: %v", err)
	}*/

	err := json.Unmarshal(jsonData, &jsonUt)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Din tale betyr: ", jsonUt.SpeakingText)

	//testing
	/*data := verdier.(map[string]interface{})

	for beskrivelse, verdier := range data {
		switch verdier := verdier.(type) {
		case []interface{}:
			for ant, innhold := range verdier {
				fmt.Println(beskrivelse, ":", innhold)
				fmt.Println("Antall prøver/maps: ", ant+1)
			}
		default:
			fmt.Println(beskrivelse, verdier)
		}
	}*/
	//fmt.Println(speechText.Outcomes.Confidence, speechText.SpeakingText)

	//fmt.Println(strings.Trim(msg, "{_text :")

}

/*
func main(){
  speech.SetWitKey("//Your Wit API Key here") //Currently ContinuousRecognition() uses wit.ai for speech recognition
  speech.ContinuousRecognition()
}
*/
