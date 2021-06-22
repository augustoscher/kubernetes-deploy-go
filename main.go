package main

import (
	"fmt"
	"log"
	"net/http"
)

const webContent = "augustoscher:v100"

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Printf("Running on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
