package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var dataApi = os.Getenv("DATA_API")

// This should be split in a separate function
func main() {
	http.HandleFunc("/clicks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://challenge1.kpmg.interviews.bajescu.com")

		req, err := http.NewRequest(r.Method, fmt.Sprintf("%s/clicks", dataApi), nil)
		if err != nil {
			fmt.Fprint(w, "Couldn't make the request.")
			return
		}

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Fprint(w, "Couldn't connect to data server")
			return
		}
		defer r.Body.Close()

		responseBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Fprint(w, "Couldn't read response")
			return
		}
		w.Write(responseBytes)

		clicks, err := strconv.Atoi(string(responseBytes))
		if err != nil {
			log.Printf("Couldn't convert response to int: %s", err)
		}
		if clicks != 0 && clicks % 5 == 0 {
			fmt.Fprint(w, " WELL DONE")
		}
	})

	var addr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}