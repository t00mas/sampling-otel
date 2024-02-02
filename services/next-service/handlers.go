package main

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func doNormalStuff(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer(name).Start(r.Context(), "normal")
	defer span.End()
	doSomething(ctx)
	w.Write([]byte("Service is up!"))
}

func doSomething(ctx context.Context) {
	_, span := otel.Tracer(name).Start(ctx, "do something")
	defer span.End()
}

func doLongStuff(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer(name).Start(r.Context(), "long")
	defer span.End()
	doLongThing(ctx)
	w.Write([]byte("Service is up!"))
}

func doLongThing(ctx context.Context) {
	_, span := otel.Tracer(name).Start(ctx, "do long thing")
	defer span.End()

	time.Sleep(1000 * time.Millisecond)
}

func doErrorStuff(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer(name).Start(r.Context(), "error")
	defer span.End()
	doFailureThing(ctx)
	w.Write([]byte("Service is up!"))
}

func doFailureThing(ctx context.Context) {
	_, span := otel.Tracer(name).Start(ctx, "do another thing")
	span.SetStatus(codes.Error, "do another thing failed")
	defer span.End()
}

func Normal() http.Handler {
	return otelhttp.NewHandler(http.HandlerFunc(doNormalStuff), "Normal")
}

func Long() http.Handler {
	return otelhttp.NewHandler(http.HandlerFunc(doLongStuff), "Long")
}

func Error() http.Handler {
	return otelhttp.NewHandler(http.HandlerFunc(doErrorStuff), "Error")
}
