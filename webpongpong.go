package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

const retries = 100
const jsonAnswer = "{\"length\": %v, \"code\": %v, \"retries\": %v, \"duration\": %v}"

var pongUrl string

func get(url string) (content string, code int, err error) {

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

func pingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i := 1; i <= retries; i++ {
		message, _, err := get(pongUrl + "/pong/" + params["length"])
		if err == nil {
			fmt.Fprintf(w, message)
			return
		}
	}
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte("Service unavailable"))
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["length"])
	if err != nil {
		panic(err)
	}
	i -= 4
	result := bytes.NewBufferString("po")
	for j := 0; j < i; j++ {
		result.WriteString("o")
	}
	result.WriteString("ng")
	w.Write(result.Bytes())
}

func Start(port int, ponghost string, pongport int) {
	myport := fmt.Sprintf(":%v", port)

	pongUrl = fmt.Sprintf("http://%s:%v", ponghost, pongport)

	r := mux.NewRouter()
	r.HandleFunc("/ping/{length:[0-9]+}", pingHandler)
	r.HandleFunc("/pong/{length:[0-9]+}", pongHandler)

	fmt.Printf("Ping service is up and listening on port %v\n", port)
	fmt.Printf("Pong service assumed to be reachable at %s\n", pongUrl)
	http.ListenAndServe(myport, r)
}

func main() {
	Start(8080, "localhost", 8080)
	os.Exit(0)
}
