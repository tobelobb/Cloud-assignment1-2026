# Copilot Instructions for countryinfo

## Project Overview
A Go REST API that aggregates country information from external APIs (REST Countries and Currency APIs). The service provides endpoints to query country data, exchange rates, and system status.

## Architecture

### Service Structure
- **main.go**: Server initialization and route registration (`:8080`)
- **handlers/**: HTTP request handlers for each endpoint
  - `status.go`: System health check endpoint returning API uptime
  - `info.go`: Country information aggregation including name, continent(s), population, area, languages, borders with exchange rates (skeleton only), flag and capital city
  - `exchange.go`: Border/currency exchange data (skeleton only)

### External Dependencies
- REST Countries API: `http://129.241.150.113:8080/v3.1/` (country data)
- Currency API: `http://129.241.150.113:9090/currency/` (exchange rates)

### Data Flow Pattern
1. Client → HTTP handler in `handlers/`
2. Handler extracts parameters from request
3. Handler calls external APIs via clients (not yet implemented)
4. Response aggregated and returned as JSON

## Critical Patterns & Conventions

### Handler Pattern
All handlers follow this signature:
```go
func HandlerName(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // Handle logic
    json.NewEncoder(w).Encode(resp)
}
```

### Response Format
Use `map[string]interface{}` with `json.NewEncoder` for flexibility:
```go
resp := map[string]interface{}{"key": value}
json.NewEncoder(w).Encode(resp)
```

### Uptime Tracking
- `startTime` variable in status handler captures server startup time
- Use `time.Since(startTime).Seconds()` for uptime calculation

## Known Structure Issues
- **clients/ directory**: Empty - needs HTTP client implementations for REST Countries and Currency APIs
- **models/ directory**: Empty - needs struct definitions for API responses
- **Incomplete handlers**: `info.go` and `exchange.go` contain only comments, no implementation
- **main.go**: Routes defined but not all registered (`handleEndpoint1` unused, `stausendpoint` unused)

## Next Steps for AI Assistance
1. Create client functions in `clients/` to fetch from external APIs
2. Define response models in `models/` for type safety
3. Implement `ExchangeHandler` and `InfoHandler` logic
4. Wire all routes properly in `main.go`

## Development Notes
- All code is in a single `main` package (main.go) and `handlers` package
- External APIs are hardcoded in main.go—consider moving to config
- errors should be handled gracefully with appropriate HTTP status codes and messages