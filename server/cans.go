package server

import "net/http"

func main() {
    panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}
