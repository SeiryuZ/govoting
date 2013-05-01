package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello")
}

func initHandler() {
	http.HandleFunc("/", rootHandler)
}

func main() {

	initHandler()
	fmt.Println("Starting listening")

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}

}
