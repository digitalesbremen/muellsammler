package main

import (
	"fmt"
	"log"
	"muellsammler/client"
	"os"
)

func main() {
	log.Println("Hello muellsammler!")

	response, err := client.NewClient().GetContent("https://web.c-trace.de/bremenabfallkalender/(S(nnititlfxtt4vnvhqcudbrlk))/Data/Strassen")

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for _, element := range response.Addresses {
		log.Println(element)
	}
}
