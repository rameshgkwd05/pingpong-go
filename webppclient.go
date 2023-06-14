package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get(url string) (content string, code int, err error) {
	fmt.Println("GET ", url)
	res, err := http.Get(url)
	if err != nil {
		return "", http.StatusServiceUnavailable, err
	}

	message, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", http.StatusServiceUnavailable, err
	}

	return string(message), res.StatusCode, err
}

func main() {
	fmt.Println("testing pipngpong")
	for i := 0; i < 5; i++ {
		message, _, err := get("http://localhost:8080" + "/ping/" + fmt.Sprintf("%v", i+4))
		if err == nil {
			fmt.Println(message)
		}
		// message2, _, err2 := get("http://localhost:8081" + "/ping/" + fmt.Sprintf("%v", i+4))
		// if err2 == nil {
		// 	fmt.Println(message2)
		// }
	}

}
