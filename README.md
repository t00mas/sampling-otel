# Sampling with OpenTelemetry

Tail sampling, and probabilistic-based sampling, at collector level.
Metric generation at collector level.

## Quickstart

```shell
    docker-compose up --build -d 
```

Grafana reachable at localhost:3000

## Metrics

### servicegraph
```
    traces_service_graph_request_client_seconds_bucket
    traces_service_graph_request_client_seconds_count
    traces_service_graph_request_client_seconds_sum
    traces_service_graph_request_server_seconds_bucket
    traces_service_graph_request_server_seconds_count
    traces_service_graph_request_server_seconds_sum
    traces_service_graph_request_total
```

### spanmetrics

```
    calls
    duration_bucket
    duration_count
    duration_sum
```

## Dependencies

```
    grafana
    k6
    loki
    opentelemetry
    prometheus
    tempo
```