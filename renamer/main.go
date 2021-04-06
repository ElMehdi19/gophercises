package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type file struct {
	filename, oldPath, newPath string
}

func main() {
	dir := "./sample"
	files, err := nonRecursive(dir)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, file := range files {
		fmt.Printf("mv %s => %s\n", file.oldPath, file.newPath)
		if err := os.Rename(file.oldPath, file.newPath); err != nil {
			panic(err)
		}
	}
}

func nonRecursive(dir string) ([]file, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []file

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if _, err := match(entry.Name(), 0); err == nil {
			files = append(files, file{
				filename: entry.Name(),
				oldPath:  fmt.Sprintf("%s/%s", dir, entry.Name()),
			})
		}
	}

	for i := 0; i < len(files); i++ {
		parsed, _ := match(files[i].filename, len(files))
		files[i].newPath = fmt.Sprintf("%s/%s", dir, parsed)
	}

	return files, nil
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
