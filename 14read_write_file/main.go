package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is a file to be saved for test."
	writeFile("./myFile.txt", content)
	readFile("./myFile.txt")
}

func writeFile(filename string, content string) {

	file, err := os.Create(filename)
	checkErr(err)

	length, err := io.WriteString(file, content)
	checkErr(err)

	fmt.Println("Length of content ", length)
	defer file.Close()
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkErr(err)

	fmt.Println("The data of myFile is:\n", string(data))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
