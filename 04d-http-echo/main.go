// Example 04d-http-echo — expose /version via Echo.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	becho "github.com/ubgo/buildinfo/contrib/buildinfo-echo"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	e := echo.New()
	e.HideBanner = true

	becho.Mount(e) // GET /version

	go func() {
		log.Println("listening on :8080 — try: curl http://localhost:8080/version")
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = e.Shutdown(shutdownCtx)
}
