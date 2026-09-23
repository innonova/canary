// canary: answers with its version, host and release tag. Deployed first
// on every host of the iac estate so the deploy path and the routing can
// be checked independently of any real service.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var version = "dev" // set at build time with -ldflags "-X main.version=..."

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8100"
	}
	host, _ := os.Hostname()
	started := time.Now().UTC()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"service": "canary",
			"version": version,
			"release": os.Getenv("RELEASE_TAG"),
			"host":    host,
			"started": started.Format(time.RFC3339),
			"path":    r.URL.Path,
		})
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "ok") })
	log.Printf("canary %s on %s listening on :%s", version, host, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
