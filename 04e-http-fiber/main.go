// Example 04e-http-fiber — expose /version via Fiber.
package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"

	bfiber "github.com/ubgo/buildinfo/contrib/buildinfo-fiber"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	bfiber.Mount(app) // GET /version

	go func() {
		log.Println("listening on :8080 — try: curl http://localhost:8080/version")
		if err := app.Listen(":8080"); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = app.ShutdownWithContext(shutdownCtx)
}
