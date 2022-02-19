package fly

import (
	"fmt"
	"log"
	"net/http"
)

// HealthCheck creates an HTTP server for fly.io healthchecks, listening on
// port 8080.
func HealthCheck() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("error starting server:", err)
	}
}
