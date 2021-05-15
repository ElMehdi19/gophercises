package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isDevMode() bool {
	return os.Getenv("dev_mode") == "true"
}

func funcThatPanics() {
	panic("Oh jeez Rick!!")
}

type sourceCodeFile struct {
	path    string
	lineNum int
	addr    string
}

func (s sourceCodeFile) link() string {
	return fmt.Sprintf("<a href='/debug/?path=%s' target='blank'>%s</a>:%d %s", s.path, s.path, s.lineNum, s.addr)
}

func parseStackLine(line string) *sourceCodeFile {
	var code sourceCodeFile
	var err error
	parts := strings.Split(line, ":")
	code.path = parts[0]

	parts = strings.Fields(parts[1])
	code.lineNum, err = strconv.Atoi(parts[0])

	if err != nil {
		log.Printf("error: %s", err.Error())
		return nil
	}
	if len(parts) > 1 {
		code.addr = parts[1]
	}

	return &code
}

func createLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	// for _, line := range lines {
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "\t") {
			line := parseStackLine(lines[i])
			if line == nil {
				log.Printf("couldn't parse line %d: %s", i, lines[i])
				continue
			}
			lines[i] = line.link()
		}
	}

	return strings.Join(lines, "\n")
}
