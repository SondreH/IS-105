package main

import (
	"github.com/karlosmartines/karls-windows/Golang/is105-ica03/oppgave5tcp/tcpclient"
	"github.com/karlosmartines/karls-windows/Golang/is105-ica03/oppgave5tcp/tcpserver"
)

func main() {
	go tcpclient.Dial()
	tcpserver.ServeTCP()

}
