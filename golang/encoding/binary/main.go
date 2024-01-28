package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	b := []byte("hello world")
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.BigEndian, uint64(len(b))); err != nil {
		log.Fatal(err)
	}

	_, err := buf.Write(b)
	if err != nil {
		log.Fatal(err)
	}

	p := make([]byte, 8)
	if _, err := buf.Read(p); err != nil {
		log.Fatal(err)
	}

	pp := make([]byte, binary.BigEndian.Uint64(p))
	if _, err := buf.Read(pp); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(pp))
}
