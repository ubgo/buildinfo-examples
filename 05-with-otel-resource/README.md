# 05-with-otel-resource

Build an OpenTelemetry `resource.Resource` from `buildinfo` metadata and print its attributes.

## What it demonstrates

- `botel.Attributes()` returns `[]attribute.KeyValue` ready to drop into `resource.WithAttributes(...)`.
- The build attributes (`build.version`, `build.commit`, `build.branch`, …) coexist with your service attributes (`service.name`, `service.namespace`).
- In a real service, the resulting resource is passed to `sdktrace.NewTracerProvider(sdktrace.WithResource(res))` and `sdkmetric.NewMeterProvider(sdkmetric.WithResource(res))` so every span and metric carries the build context automatically.

## Run

```sh
go run .
```

## Expected output

```
--- resource attributes ---
build.branch           = unknown
build.commit           = ...
build.go_version       = go1.24.x
build.goarch           = ...
build.goos             = ...
build.modified         = false
build.time             = ...
build.version          = dev
service.name           = myapi
service.namespace      = demo
telemetry.sdk.language = go
telemetry.sdk.name     = opentelemetry
...
```

(`telemetry.sdk.*` attributes are added automatically by the OTEL SDK; the `build.*` and `service.*` ones come from this example.)

## Wiring into TracerProvider / MeterProvider (sketch)

```go
res, _ := resource.New(ctx,
    resource.WithAttributes(attribute.String("service.name", "myapi")),
    resource.WithAttributes(botel.Attributes()...),
)

tp := sdktrace.NewTracerProvider(
    sdktrace.WithResource(res),
    sdktrace.WithBatcher(otlpExporter),
)
otel.SetTracerProvider(tp)

mp := sdkmetric.NewMeterProvider(
    sdkmetric.WithResource(res),
    sdkmetric.WithReader(sdkmetric.NewPeriodicReader(prometheusExporter)),
)
otel.SetMeterProvider(mp)
```

Every span and metric exported from this point on will carry the `build.*` resource attributes.
