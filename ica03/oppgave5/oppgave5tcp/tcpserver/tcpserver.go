package tcpserver

import (
	"fmt"
	"net"

	"github.com/karlosmartines/karls-windows/Golang/is105-ica03/oppgave5tcp/jsonemail"
)

func handler(c net.Conn) {
	json := jsonemail.EncodeJSON()
	c.Write(json)
	c.Close()
}

// ServeTCP starts the server
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
