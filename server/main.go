package main

import (
    "net/http"
    "fmt"
    "encoding/json"
)

type JsonData struct {
    Message string
}

func createResponse() string {
    response := JsonData{Message: "Hello World"}
    outputJson, err := json.Marshal(&response)

    if err != nil {
        panic(err)
    }

    return string(outputJson)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    response := createResponse()

    w.Header().Set("Content-Type", "application/json")

    fmt.Fprintf(w, response)
}

func main() {
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8081", nil)
}
