package tcpclient

import (
	"fmt"
	"io/ioutil"
	"net"
)

// Dial dials
func Dial() {
	fmt.Println("TCP dialed")
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(bs)
}
