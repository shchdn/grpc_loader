package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	var filepath string
	myId := generateClientId()
	client, err := createClient()
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Pass me a filepath:")
		fmt.Scanln(&filepath)
		for {
			err := uploadFileToServer(client, myId, filepath)
			if err == nil {
				break
			}
		}
	}
}
