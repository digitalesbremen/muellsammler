package main

import (
	"fmt"
	"log"
	"muellsammler/client"
	"os"
)

const (
	BaseURL = "https://web.c-trace.de"
)

func main() {
	log.Println("Hello muellsammler!")

	client := client.NewClient(BaseURL)

	response, err := client.ReadStreets("/bremenabfallkalender/(S(nnititlfxtt4vnvhqcudbrlk))/Data/Strassen")

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for _, element := range response.Streets {
		log.Println(element)
	}
}
