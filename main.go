package main

import (
	"fmt"
	"log"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
		return 
	}
	defer file.Close()

	eightByteData := make([]byte, 8)

	for {
		currentLine := ""
		noOfBytes, err := file.Read(eightByteData)

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
			return
		}

		lines := strings.Split(eightByteData, "\n")  
		for i, line := range lines {
			currentLine = currentLine + line
			
		} 

	}

}
