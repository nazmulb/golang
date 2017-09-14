package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test1.txt")

	if err != nil {
		fmt.Println("file create failed:", err)
		return
	}

	defer file.Close()

	file.WriteString("testing...")
}
