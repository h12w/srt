package main

import (
	"fmt"
	"log"
	"os"

	"strings"

	"h12.io/srt"
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
	last := ""
	for scanner.Scan() {
		this := strings.TrimSpace(scanner.Record().Text)
		a, b := cleanAdjacent(last, this)
		if a != "" {
			fmt.Println(a)
		}
		last = b
	}
}

func cleanAdjacent(a, b string) (string, string) {
	n := len(a)
	if n > len(b) {
		n = len(b)
	}
	for i := n; i >= 1; i-- {
		if a[len(a)-i:] == b[0:i] {
			return a[:len(a)-i], b[i:]
		}
	}
	return a, b
}
