package main

import (
	"fmt"
	"os"
)

func main () {
	file, err := os.Open("test.txt")

	if err != nil {
		// handle the error here
		fmt.Println("file open failed:", err)
		return
	}

	defer file.Close()

	// get the file size
	stat, err := file.Stat()

	if err != nil {
		fmt.Println("file stat faild:", err)
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	if err != nil {
		fmt.Println("file read failed:", err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}
