package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/jixunmoe-go/netease/crypto"
	"log"
	"os"
)

func main() {
	argv := os.Args
	l := log.New(os.Stderr, "", 0)

	if len(argv) <= 1 {
		l.Printf("%s: Response packet decryptor.\n", argv[0])
		l.Printf("Usage:\n")
		l.Printf("%s <response body in hex or b64 str>", argv[0])
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

	resp := crypto.DecryptEAPIResponse(payload)
	if err != nil {
		l.Println("could not decrypt data")
		os.Exit(1)
		return
	}

	fmt.Printf("Size: %d\n", len(resp))
	fmt.Println(string(resp))
}
