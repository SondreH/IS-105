// Package main in compression.go is a program for compressing data to gzip format.
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

// Func check checks for an error.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Func hexEncode takes data in a slice of bytes, encodes it in hexadecimals according to the Unicode-model
// and returns it as a slice of bytes and an integer of how many bytes it is.
func hexEncode(data []byte) ([]byte, int) {
	hexEncoding := make([]byte, hex.EncodedLen(len(data)))
	len := hex.Encode(hexEncoding, data)
	return hexEncoding, len
}

// Func base64Encode take data in a slice of bytes, encodes it in base64 and returns it as a slice of bytes
// and an integer of how many bytes it is.
func base64Encode(data []byte) ([]byte, int) {
	
	asciislice := make([]byte, 0, len(data))
	buf := bytes.NewBuffer(asciislice)

	b64Writer := base64.NewEncoder(base64.StdEncoding, buf)

	_, err := b64Writer.Write(data)
	check(err)
	b64Writer.Close()

	b64Slice := buf.Bytes()

	fmt.Printf("b64Slice: %s\n", b64Slice)

	b64Slice, b64EnLen := hexEncode(b64Slice)

	return b64Slice, b64EnLen
}

// Func gzipEncode takes data as a slice of bytes, encodes it in gzip-format and returns it as a slice
// of bytes and an integer of how many bytes it is.
func gzipEncode(data []byte) ([]byte, int) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	zw.Write(data)
	gzipSlice := buf.Bytes()
	fmt.Printf("gzipSlice: %v\n", gzipSlice)
	gzipSlice, gzipLen := hexEncode(gzipSlice)
	return gzipSlice, gzipLen
}

// Func main takes an argument from the command line and reads the information. It then calls the above
// encoding functions to encode it in hexadecimals, base64 and gzip and prints the results of each part
// of the process.
func main() {
	if len(os.Args) < 2 {
		panic("1")
	}

	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	fmt.Printf("Original text: %s\n", data)

	hexEncoding, hexEnLen := hexEncode(data)
	fmt.Printf("Hexadecimal encoding: %s\n", hexEncoding)
	fmt.Printf("Hexadecimal encoding length: %d\n", hexEnLen)

	b64Encoding, b64EnLen := base64Encode(data)
	fmt.Printf("Base64 hex encoding: %x\n", b64Encoding)
	fmt.Printf("Base64 hex encoding length: %d\n", b64EnLen)

	gzipEncoding, gzipLen := gzipEncode(b64Encoding)
	fmt.Printf("Gzip hex encoding: %x\n", gzipEncoding)
	fmt.Printf("Gzip hex encoding length: %d\n", gzipLen)
}
