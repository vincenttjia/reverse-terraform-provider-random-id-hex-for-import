package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
  // Insert your hexcode here
	s := "971c842539f7ec43"
	decoded, err := hex.DecodeString(s)

	if err != nil {
		log.Fatal(err)
	}

	id := base64.RawURLEncoding.EncodeToString(decoded)
	fmt.Println(id)
}
