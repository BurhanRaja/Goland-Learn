package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://localhost:3000/learn?coursename=reactjs&user=burhan"

func main() {
	result, _ := url.Parse(myurl)
	fmt.Println(result)

	fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Path)

	// qparams := result.Query()

	// fmt.Println(qparams["coursename"])

	// for _, val := range qparams {
	// 	defer fmt.Println(val)
	// }

	// Create a url
	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "localhost:3000",
		RawQuery: "coursename=nodejs",
		Path:     "/learn",
	}

	fmt.Println(partsOfUrl.String())

}
