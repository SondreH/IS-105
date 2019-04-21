package main

import (
	speech "github.com/SondreH/IS-105/ica04/Oppgave3/Del2/go-speak"
)

//example
//var key = string //note that this package is imported as speech

func main() {
	speech.SetWitKey("ffdb4901977bc08490ab580e93be771c7a0a2529") //Wit API Key MUST be set before calling any other Wit.AI functions
	speech.SendWitVoice("SondreOgTarIS105.wav")
}

/*
func main(){
  speech.SetWitKey("//Your Wit API Key here") //Currently ContinuousRecognition() uses wit.ai for speech recognition
  speech.ContinuousRecognition()
}
*/
