package main

import (
	"fmt"
	"log"
	"muellsammler/client"
	"os"
)

func main() {
	log.Println("Hello muellsammler!")

	response, err := client.NewClient().ReadStreets("https://web.c-trace.de/bremenabfallkalender/(S(nnititlfxtt4vnvhqcudbrlk))/Data/Strassen")

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for _, element := range response.Streets {
		log.Println(element)
	}
}
