package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}

	hash := sha256.Sum256(file)
	fmt.Printf("%x", hash)

}
