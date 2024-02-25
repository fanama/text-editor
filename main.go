package main

import (
	filemanager "fanama/text-editor/infra/fileManager"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Go Text Editor!")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]

	editor, err := filemanager.Open(filename)

	if err != nil {
		return
	}

	editor.Edit()
}
