package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 4 {
		log.Fatalf("Remember to use the format: go run main.go file1 file2 file3")
	}

	file1name := os.Args[1]
	file2name := os.Args[2]
	file3name := os.Args[3]

	file1, err := os.Open(file1name)
	if err != nil {
		log.Fatalf("Error reading argument: %v", err)
	}
	defer file1.Close()

	file2, err := os.Open(file2name)
	if err != nil {
		log.Fatalf("Error reading argument: %v", err)
	}
	defer file2.Close()

	file3, err := os.Create(file3name)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file3.Close()

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	writer3 := bufio.NewWriter(file3)
	defer writer3.Flush()

	for scanner1.Scan() && scanner2.Scan() {
		line1 := scanner1.Text()
		line2 := scanner2.Text()

		if line1 == "" && line2 == "" {
			writer3.WriteString("\n")
		} else if line1 == "" && line2 != "" {
			writer3.WriteString(line2 + "\n")
		} else if line2 == "" && line1 != "" {
			writer3.WriteString(line1 + "\n")
		} else {

			elements1 := strings.Split(line1, " ")
			elements2 := strings.Split(line2, " ")
			if len(elements1) != 2 || len(elements2) != 2 {
				log.Fatalf("Invalid line format")
			}

			line3 := elements1[0] + " " + elements2[0] + " "
			num1, err := strconv.Atoi(elements1[1])
			if err != nil {
				log.Fatalf("could not convert number: %v", err)
			}
			num2, err := strconv.Atoi(elements2[1])
			if err != nil {
				log.Fatalf("could not convert number: %v", err)
			}
			numbers3 := num1 + num2

			line3 += strconv.Itoa(numbers3) + "\n"
			_, err = writer3.WriteString(line3)
			if err != nil {
				log.Fatalf("Error writing to file3: %v", err)
			}
		}

	}

	for scanner1.Scan() {
		writer3.WriteString(scanner1.Text() + "\n")
	}

	for scanner2.Scan() {
		writer3.WriteString(scanner2.Text() + "\n")
	}

	if err := scanner1.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	if err := scanner2.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

}
