// Example 04-http-nethttp — expose /version via stdlib net/http.
//
// Mounts the version handler on a stdlib ServeMux at the default route,
// then starts a server on :8080. Hit http://localhost:8080/version with
// curl to see the JSON.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	binethttp "github.com/ubgo/buildinfo/contrib/buildinfo-nethttp"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	binethttp.Mount(mux) // GET /version

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Println("listening on :8080 — try: curl http://localhost:8080/version")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Println("shutdown error:", err)
		os.Exit(1)
	}
}
