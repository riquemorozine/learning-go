package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("archive.txt")

	if err != nil {
		panic(err)
	}

	length, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	//length, err := f.WriteString("Hello, World!")
	//if err != nil {
	//	panic(err)
	//}

	println(length, "bytes written successfully")

	f.Close()

	// reading
	file, err := os.ReadFile("archive.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// reading with buffer

	file2, err := os.Open("archive.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 2)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("archive.txt")
	if err != nil {
		panic(err)
	}
}
