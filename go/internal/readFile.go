package internal

import (
	"bufio"
	"os"
)

func readLines(filename string, fileChannel chan string) {
	file, _ := os.Open("../inputs/" + filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileChannel <- scanner.Text()
	}
	close(fileChannel)
	file.Close()
}

func ReadLines(filename string) chan string {
	fileChannel := make(chan string)

	go readLines(filename, fileChannel)

	return fileChannel
}
