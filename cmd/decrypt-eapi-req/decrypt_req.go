package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/JixunMoe-go/NetEaseAPI/crypto"
	"log"
	"os"
)

func main() {
	argv := os.Args
	l := log.New(os.Stderr, "", 0)

	if len(argv) <= 1 {
		l.Printf("%s: Request packet decryptor.\n", argv[0])
		l.Printf("Usage:\n")
		l.Printf("%s <request body in hex or b64 str>", argv[0])
		l.Printf("\n")
		return
	}

	payload, err := hex.DecodeString(argv[1])
	if err != nil {
		payload, err = base64.StdEncoding.DecodeString(argv[1])
		if err != nil {
			l.Println("could not decode argument (hex/base64)")
			os.Exit(1)
			return
		}
	}

	req, err := crypto.DecryptEAPIRequestBody(payload)
	if err != nil {
		l.Println("could not decrypt data")
		os.Exit(1)
		return
	}

	fmt.Printf("Path: %s\n", req.Path)
	fmt.Printf("Data: %s\n", req.Data)
}
