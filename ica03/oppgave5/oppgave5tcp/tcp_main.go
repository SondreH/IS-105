package main

import (
	"github.com/shammers95/IS-105/ica03/oppgave5/oppgave5tcp/tcpclient"
	"github.com/shammers95/IS-105/ica03/oppgave5/oppgave5tcp/tcpserver"
)

func main() {
	go tcpclient.Dial()
	tcpserver.ServeTCP()

}
