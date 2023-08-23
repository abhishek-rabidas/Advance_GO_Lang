package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	r, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		fmt.Println("Error in fetching data")
	}

	handleResponse(r)

}

func handleResponse(response *http.Response) {
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(responseData))

}
