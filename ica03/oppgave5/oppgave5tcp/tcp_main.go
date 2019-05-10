// Package main in this program starts a TCP-server which serves a JSON-structure
// and then starts a client to dial at the same adress to get the JSON-structure.
package main

import (
	"github.com/sondreh/IS-105/ica03/oppgave5/oppgave5tcp/tcpserver"
)

// Function main runs tcpclients Dial and tcpservers ServeTCP
func main() {
	//go tcpclient.Dial()
	tcpserver.ServeTCP()

}
