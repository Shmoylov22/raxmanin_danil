package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	str := "Hello World."

	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}
}
