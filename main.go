package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// write your code here
	var name string
	fmt.Scan(&name)
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Can't close file")
		}
	}(file)
	var tooWords = make(map[string]bool)
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		tooWords[scanner.Text()] = true
	}
	var inputWord string
	fmt.Scan(&inputWord)
	for inputWord != "exit" {
		var aWord = make([]string, len(inputWord))
		aWord = strings.Split(inputWord, " ")
		for i := 0; i < len(aWord); i++ {
			if tooWords[strings.ToLower(inputWord)] {
				for i := 0; i < len(inputWord); i++ {
					fmt.Printf("%c", '*')
				}
				fmt.Println()
			} else {
				fmt.Println(inputWord)
			}
		}
		fmt.Scan(&inputWord)
	}
	fmt.Println("Bye!")
}
