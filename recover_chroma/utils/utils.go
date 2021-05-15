package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func IsDevMode() bool {
	return os.Getenv("dev_mode") == "true"
}

func FuncThatPanics() {
	panic("Oh jeez Rick!!")
}

type sourceCodeFile struct {
	path    string
	lineNum int
	addr    string
}

func (s sourceCodeFile) link() string {
	val := url.Values{}
	val.Add("path", s.path)
	val.Add("line", fmt.Sprint(s.lineNum))
	log.Println("values:", val.Encode())
	return fmt.Sprintf("\t<a href='/debug/?%s' target='blank'>%s</a>:%d %s", val.Encode(), s.path, s.lineNum, s.addr)
}

func parseStackLine(line string) (sourceCodeFile, error) {
	var code sourceCodeFile
	var err error
	if !strings.HasPrefix(line, "\t") {
		return code, fmt.Errorf("no match")
	}

	parts := strings.Split(line, ":")
	code.path = strings.TrimSpace(parts[0])

	parts = strings.Fields(parts[1])
	code.lineNum, err = strconv.Atoi(parts[0])

	if err != nil {
		log.Printf("error: %s", err.Error())
		return code, err
	}
	if len(parts) > 1 {
		code.addr = parts[1]
	}

	return code, nil
}

func CreateLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	// for _, line := range lines {
	for i := 0; i < len(lines); i++ {
		if line, err := parseStackLine(lines[i]); err == nil {
			lines[i] = line.link()
		}
	}

	return strings.Join(lines, "\n")
}
