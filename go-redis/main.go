// package main

// import (
//     "context"
//     "encoding/json"
//     "fmt"
//     "log"

//     "github.com/redis/go-redis/v9"
//     "github.com/redis/go-redis/extra/redisotel/v9"
//     "go.opentelemetry.io/otel"
//     "go.opentelemetry.io/otel/attribute"
//     "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
//     "go.opentelemetry.io/otel/sdk/resource"
//     "go.opentelemetry.io/otel/sdk/trace"
//     "go.opentelemetry.io/otel/semconv/v1.4.0"
// )

// type Author struct {
//     Name string `json:"name"`
//     Age  int    `json:"age"`
// }

// func main() {
//     // Create a new OTLP exporter
//     exp, err := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
//     if err != nil {
//         log.Fatalf("failed to initialize OTLP exporter: %v", err)
//     }

//     // Create a new tracer provider with the exporter
//     tp := trace.NewTracerProvider(
//         trace.WithBatcher(exp),
//         trace.WithResource(resource.NewWithAttributes(
//             semconv.SchemaURL,
//             attribute.String("service.name", "go-redis-testing"),
//         )),
//     )
//     defer func() {
//         if err := tp.Shutdown(context.Background()); err != nil {
//             log.Fatalf("failed to shutdown TracerProvider: %v", err)
//         }
//     }()

//     // Set the global tracer provider
//     otel.SetTracerProvider(tp)

//     // Create a tracer
//     tracer := otel.Tracer("example-tracer")

//     // Start a span
//     ctx, span := tracer.Start(context.Background(), "main-span")
//     defer span.End()

//     rdb := redis.NewClient(&redis.Options{
//         Addr:     "localhost:6379",
//         Password: "",
//         DB:       0,
//     })

//     // Enable tracing instrumentation.
//     if err := redisotel.InstrumentTracing(rdb); err != nil {
//         panic(err)
//     }

//     author := Author{Name: "Elliot", Age: 25}
//     jsonData, err := json.Marshal(author)
//     if err != nil {
//         fmt.Println(err)
//     }

//     err = rdb.Set(ctx, "id1234", jsonData, 0).Err()
//     if err != nil {
//         fmt.Println(err)
//     }
//     val, err := rdb.Get(ctx, "id1234").Result()
//     if err != nil {
//         fmt.Println(err)
//     }
//     fmt.Println(val)

//     // Add some attributes to the span
//     span.SetAttributes(attribute.String("example-attribute", "example-value"))
// }

package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"

    "github.com/redis/go-redis/v9"
    "github.com/redis/go-redis/extra/redisotel/v9"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type Author struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// MultiExporter is a custom SpanExporter that forwards spans to multiple exporters.
type MultiExporter struct {
    exporters []trace.SpanExporter
}

func (m *MultiExporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
    for _, exporter := range m.exporters {
        if err := exporter.ExportSpans(ctx, spans); err != nil {
            return err
        }
    }
    return nil
}

func (m *MultiExporter) Shutdown(ctx context.Context) error {
    for _, exporter := range m.exporters {
        if err := exporter.Shutdown(ctx); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    // Create a new stdout exporter
    stdoutExp, err := stdouttrace.New(
        stdouttrace.WithPrettyPrint(),
    )
    if err != nil {
        log.Fatalf("failed to initialize stdout exporter: %v", err)
    }

    // Create a new OTLP HTTP exporter
    otlpExp, err := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
    if err != nil {
        log.Fatalf("failed to initialize OTLP exporter: %v", err)
    }

    // Create a multi-exporter
    multiExp := &MultiExporter{
        exporters: []trace.SpanExporter{stdoutExp, otlpExp},
    }

    // Create a new tracer provider with the multi-exporter
    tp := trace.NewTracerProvider(
        trace.WithBatcher(multiExp),
        trace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            attribute.String("service.name", "go-redis-testing"),
        )),
    )
    defer func() {
        if err := tp.Shutdown(context.Background()); err != nil {
            log.Fatalf("failed to shutdown TracerProvider: %v", err)
        }
    }()

    // Set the global tracer provider
    otel.SetTracerProvider(tp)

    // Create a tracer
    tracer := otel.Tracer("example-tracer")

    // Start a span
    ctx, span := tracer.Start(context.Background(), "main-span")
    defer span.End()

    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    // Enable tracing instrumentation.
    if err := redisotel.InstrumentTracing(rdb); err != nil {
        panic(err)
    }

    author := Author{Name: "******", Age: 25}
    jsonData, err := json.Marshal(author)
    if err != nil {
        fmt.Println(err)
    }

    err = rdb.Set(ctx, "id1234", jsonData, 0).Err()
    if err != nil {
        fmt.Println(err)
    }
    val, err := rdb.Get(ctx, "id1234").Result()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(val)

    // Add some attributes to the span
    span.SetAttributes(attribute.String("example-attribute", "example-value"))
}
