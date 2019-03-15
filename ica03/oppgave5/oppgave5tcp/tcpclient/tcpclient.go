Package tcpclient dials at an adress and displays the result
package tcpclient

import (
	"fmt"
	"io/ioutil"
	"net"
)

// Dial dials at the given adress and prints out a string of the result
func Dial() {
	fmt.Println("TCP dialed")
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}
