package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	action = flag.String("action", "", "encrypt/decrypt")
	file   = flag.String("file", "", "file path")
)

func init() {
	flag.Parse()
}

func main() {
	f, err := ioutil.ReadFile(*file)
	fmt.Println(len(f))

	if err != nil {
		log.Fatalf("read file fail, error: %s", err)
	}

	var content []byte
	switch *action {
	case "encrypt":
		buffer := f[:2]
		encrypt, err := AesCFBEncrypt(buffer)
		if err != nil {
			log.Fatalf("encrypt error, error: %s", err)
		}
		content = append(encrypt, f[2:]...)

	case "decrypt":
		buffer := f[:18]
		decrypt, err := AesCFBDecrypt(buffer)
		if err != nil {
			log.Fatalf("decrypt error, error: %s", err)
		}
		content = append(decrypt, f[18:]...)
	default:
		log.Fatal("unsupported action")
	}

	fmt.Println(len(content))
	err = ioutil.WriteFile(*file, content, 644)
	if err != nil {
		log.Fatalf("write file error, error: %s", err)
	}
}
