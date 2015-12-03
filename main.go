package main
import (
	"flag"
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
	webroot := "/var/www"
	bind := "localhost:8000"

	flag.StringVar(&webroot, "webroot", "/var/www", "web application root")
	flag.StringVar(&bind, "bind", "localhost:8000", "<host>:<port> or just <host> or <port>")
	flag.Parse()

	log.Println("")
	log.Println("   `( ◔ ౪◔)´")
	log.Println("")
	log.Println(" Server running on :8000")

	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler)
	r.HandleFunc("/settings", SettingsHandler)
	r.HandleFunc("/test/lastfm", testArtistSearch)

	r.PathPrefix("/").Handler(
		handlers.LoggingHandler(os.Stdout,
			http.FileServer(http.Dir(webroot))))
    //r.HandleFunc("/socket.io", SocketHandler)
	err := http.ListenAndServe(bind, r)
	if err != nil {
		panic(err)
	}
}
