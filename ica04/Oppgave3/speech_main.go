package main

//import (
//	"github.com/SondreH/IS-105/ica04/Oppgave3/gspeech-rec-master/"
//	"github.com/SondreH/IS-105/ica04/Oppgave3/go-speak-master"
//)

func main() {
	continuous.start()

	wit.SetWitKey("qP7byWDB8PjOvVPeGPdLkS-6pvOj_DBP_F4WkQo5dkhm")
	wit.SendWitVoice("github.com/SondreH/IS-105/ica04/Oppgave3/gspeech-rec-master/hellothisiskarlmartin.m4a")
}
