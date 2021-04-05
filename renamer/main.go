package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fileName := "birthday_001.txt" // -> Birthday - 1 of 4.txt
	parsed, err := match(fileName, 4)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	fmt.Printf("%s -> %s\n", fileName, parsed)
}

func match(fileName string, total int) (string, error) {
	parts := strings.Split(fileName, ".")
	extension := parts[len(parts)-1]
	tmp := strings.Join(parts[:len(parts)-1], ".")
	parts = strings.Split(tmp, "_")
	name := strings.Join(parts[:len(parts)-1], "_")
	num, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match the pattern", fileName)
	}
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), num, total, extension), nil
}
