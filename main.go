package main

import (
	"fmt"
	"os"
	// "io"
	"bufio"
	"log"
)

func main() {
	FileRead()
}

// FileRead is FileReader
func FileRead() {
	file, err := os.Open(`base-text/test.html`)
	if err != nil {
		return
	}
	defer file.Close()


	// return bufio.NewScanner(file)

}

func FileWirte(file file) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
