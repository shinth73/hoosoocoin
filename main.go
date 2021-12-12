package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Hello from home~~")
}

func main() {
	http.HandleFunc("/", home)
    fmt.Printf("Listening on http://127.0.0.1%s\n", port)
	log.Fatal(http.ListenAndServe(port,nil))
}
