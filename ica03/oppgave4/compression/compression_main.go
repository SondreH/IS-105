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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func hexEncode(data []byte) ([]byte, int) {
	hexEncoding := make([]byte, hex.EncodedLen(len(data)))
	len := hex.Encode(hexEncoding, data)

	// fmt.Printf("Hexadecimal encoding: %s\n", hexEncoding)

	// fmt.Printf("Hexadecimal encoding length: %d\n", len)

	return hexEncoding, len
}

func base64Encode(data []byte /*, hexEnLen int, pWriter io.Writer*/) ([]byte, int) {

	// base64Encoding := make([]byte, base64.EncodedLen(hexEnLen))

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

func gzipEncode(data []byte) ([]byte, int) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	zw.Write(data)
	gzipSlice := buf.Bytes()
	fmt.Printf("gzipSlice: %v\n", gzipSlice)
	gzipSlice, gzipLen := hexEncode(gzipSlice)
	return gzipSlice, gzipLen
}

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

	b64Encoding, b64EnLen := base64Encode(data /*, hexEnLen, pWriter*/)
	fmt.Printf("Base64 hex encoding: %x\n", b64Encoding)
	fmt.Printf("Base64 hex encoding length: %d\n", b64EnLen)

	gzipEncoding, gzipLen := gzipEncode(b64Encoding)
	fmt.Printf("Gzip hex encoding: %x\n", gzipEncoding)
	fmt.Printf("Gzip hex encoding length: %d\n", gzipLen)
}
