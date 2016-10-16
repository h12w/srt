package main

import (
	"fmt"
	"log"
	"os"

	"h12.me/srt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("srt [file.srt]")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := srt.NewScanner(file)
	previousText := ""
	for scanner.Scan() {
		text := scanner.Record().Text
		if text != previousText {
			fmt.Println(text)
			previousText = text
		}
	}
}
