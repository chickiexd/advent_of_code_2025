package utils

import (
	"fmt"
	"log"
	"os"
)

func GetInput(day int) string {
	cwd, err := os.Getwd()
	inputPath := fmt.Sprintf("%s/solutions/day_%02d/input.txt", cwd, day)
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func GetTestInput(day int) string {
	cwd, err := os.Getwd()
	inputPath := fmt.Sprintf("%s/solutions/day_%02d/example.txt", cwd, day)
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
