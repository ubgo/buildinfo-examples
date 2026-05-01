// Example 05-with-otel-resource — build an OpenTelemetry resource.Resource
// from buildinfo metadata and print its attributes.
//
// In a real service you would pass this resource into a TracerProvider /
// MeterProvider so every span and metric carries the build attributes
// without per-span repetition.
package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"

	botel "github.com/ubgo/buildinfo/contrib/buildinfo-otel"
)

func main() {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			attribute.String("service.name", "myapi"),
			attribute.String("service.namespace", "demo"),
		),
		resource.WithAttributes(botel.Attributes()...),
	)
	if err != nil {
		fmt.Println("resource.New error:", err)
		return
	}

	fmt.Println("--- resource attributes ---")
	for _, a := range res.Attributes() {
		fmt.Printf("%-22s = %v\n", a.Key, a.Value.Emit())
	}
}
