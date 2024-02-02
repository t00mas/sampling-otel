package main

import (
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type App struct {
	next_service_endpoint string
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Start a span for the incoming request
	ctx, span := otel.Tracer(name).Start(r.Context(), "serveHTTP"+" - "+r.URL.Path)
	defer span.End()

	// Create a new request to the next service
	// the endpoint is the next service endpoint + the path of the incoming request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.next_service_endpoint+r.URL.Path, nil)
	if err != nil {
		http.Error(w, "couldn't create request to endpoint", http.StatusInternalServerError)
	}

	// Do the request to the next service, Transport is wrapped with otelhttp
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "couldn't do request to endpoint", http.StatusInternalServerError)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "couldn't read body", http.StatusInternalServerError)
	}

	w.Write(body)
}
