package main

import (
	"os"

	"./runecount"
)

func main() {
	var a = os.Args[1]
	runecount.FindAll(a)
	runecount.FileToByteslice(a)
}
