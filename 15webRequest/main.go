package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	response, err := http.Get(url)
	checkErr(err)

	fmt.Printf("Response is : %T\n", response)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	fmt.Println(string(data))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
