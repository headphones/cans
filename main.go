package main
import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler)
	r.PathPrefix("/").Handler(
		handlers.LoggingHandler(os.Stdout,
			http.FileServer(
				http.Dir("/Users/pnovotnak/go/src/github.com/headphones/cans/www/app"))))
    //r.HandleFunc("/socket.io", SocketHandler)
	http.ListenAndServe(":8000", r)
}
