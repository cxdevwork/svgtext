package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/svg/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}
