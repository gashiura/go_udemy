package main
import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Your url path is %s", r.URL.Path[1:])
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hello World!!")
}
