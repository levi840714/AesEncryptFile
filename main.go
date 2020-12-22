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
	fmt.Println(string(f))
	if err != nil {
		log.Fatalf("read file fail, error: %s", err)
	}

	var content string
	switch *action {
	case "encrypt":
		content, err = AesCFBEncrypt(string(f))
		if err != nil {
			log.Fatalf("encrypt error, error: %s", err)
		}
	case "decrypt":
		content, err = AesCFBDecrypt(string(f))
		if err != nil {
			log.Fatalf("decrypt error, error: %s", err)
		}
	}

	err = ioutil.WriteFile(*file, []byte(content), 644)
	if err != nil {
		log.Fatalf("write file error, error: %s", err)
	}
}
