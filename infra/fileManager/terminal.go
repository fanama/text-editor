package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

type Terminal struct {
	filename    string
	contents    []byte
	saveCommand string
}

func Open(filename string) (Terminal, error) {

	fmt.Printf("Editing file: %s\n", filename)

	contents, err := os.ReadFile(filename)
	if err != nil {

		_, err = os.Create(filename)

		if err != nil {

			fmt.Println("Error reading file:", err)
			return Terminal{}, err
		}
	}

	const save = ":w"

	return Terminal{filename, contents, save}, nil
}

func (this *Terminal) Save() error {
	err := os.WriteFile(this.filename, this.contents, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	return nil

}

func (this *Terminal) Edit() {

	fmt.Printf("\n\n\n%s", this.contents)
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == this.saveCommand {

			break
		}
		lines = append(lines, line)
	}

	for _, line := range lines {
		this.contents = append(this.contents, []byte(line)...)
		this.contents = append(this.contents, '\n')
	}
	this.Save()

}
