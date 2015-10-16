package main
import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong\n"))
}

func main() {
	log.Println("")
	log.Println("   `( ◔ ౪◔)´")
	log.Println("")
	log.Println(" Server running on :8000")
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler)
	r.PathPrefix("/").Handler(
		handlers.LoggingHandler(os.Stdout,
			http.FileServer(
				http.Dir("/var/www/"))))
    //r.HandleFunc("/socket.io", SocketHandler)
	http.ListenAndServe(":8000", r)
}
