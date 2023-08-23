package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {

	for {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 199
		randomID := rand.Intn(max-min+1) + min

		r, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(randomID))

		if err != nil {
			fmt.Println("Error in fetching data")
		}

		go handleResponse(r)
	}

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
