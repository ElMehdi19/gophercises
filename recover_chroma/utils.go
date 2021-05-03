package main

import "os"

func isDevMode() bool {
	return os.Getenv("dev_mode") == "true"
}

func funcThatPanics() {
	panic("Oh jeez Rick!!")
}
