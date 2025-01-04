package main

import (
	"log"
	"net/http"
    "fmt"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "ok")
}
func main() {

    http.HandleFunc("/", handleHello)

    log.Println("Start Server")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Println("error to start webserver")
    }
    log.Println("Finish")
}
