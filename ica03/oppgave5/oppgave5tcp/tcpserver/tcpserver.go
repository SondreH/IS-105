// Package tcpserver is a TCP-server which serves a JSON-structure from the jsonemail-package.
package tcpserver

import (
	"fmt"
	"net"

	"github.com/shammers95/IS-105/ica03/oppgave5/oppgave5tcp/jsonemail"
)

// handler is retrieving a JSON-structure and writing it to a connection
func handler(c net.Conn) {
	json := jsonemail.EncodeJSON()
	c.Write(json)
	c.Close()
}

// ServeTCP starts the server and listens for dialers
func ServeTCP() {
	fmt.Println("TCP served")
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
