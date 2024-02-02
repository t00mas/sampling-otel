package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// InitTracerProvider creates a new trace provider instance and registers it as global trace provider.
func InitTracerProvider() (func(context.Context) error, error) {
	ctx := context.Background()

	// Create a new OTLP trace exporter
	traceExporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Create a new batch span processor
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)

	// Create a new tracer provider with the batch processor we created earlier, and with a sampler that samples all traces.
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
	)

	// Register the trace provider as the global trace provider.
	otel.SetTracerProvider(tracerProvider)

	// Set the global propagator to TraceContext
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Return a function that can be used to shutdown the trace provider.
	return tracerProvider.Shutdown, nil
}
