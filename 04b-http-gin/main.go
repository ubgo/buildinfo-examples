// Example 04b-http-gin — expose /version via Gin, with auth middleware
// protecting an /internal/version variant.
//
// Demonstrates that the same Mount API works with the framework-typed
// middleware shape (gin.HandlerFunc) the user already knows.
package main

import (
	"context"
	"crypto/subtle"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	bgin "github.com/ubgo/buildinfo/contrib/buildinfo-gin"
)

const internalKey = "secret-rotate-me"

func internalKeyAuth(expected string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if subtle.ConstantTimeCompare(
			[]byte(c.GetHeader("X-Internal-Key")),
			[]byte(expected),
		) != 1 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// Public version (e.g. for k8s probes / load balancer health pages).
	bgin.Mount(r)

	// Internal version, gated behind X-Internal-Key.
	internal := r.Group("/internal", internalKeyAuth(internalKey))
	bgin.Mount(internal, bgin.WithPath("/version"))

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Println("listening on :8080")
		log.Println("  public:   curl http://localhost:8080/version")
		log.Println("  internal: curl -H 'X-Internal-Key: " + internalKey + "' http://localhost:8080/internal/version")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
}
