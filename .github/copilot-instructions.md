# GitHub Copilot Instructions for astroapi-go

## Project Context

This is a pure Go HTTP client SDK for Astrology API v3 (`https://api.astrology-api.io`). Module: `github.com/astro-api/astroapi-go`. Go 1.23. The only external dependency is `testify` for tests — everything else uses standard library.

## Architecture Overview

```
client.go                    → AstrologyClient (root, holds all sub-clients)
categories/base.go           → BaseCategoryClient (shared Get/Post/Put/Delete/MakeRequest)
categories/<name>/<name>.go  → Category Client (embeds BaseCategoryClient)
categories/<name>/params.go  → Request param structs (json + validate tags)
categories/<name>/responses.go → Response type aliases (map[string]any)
shared/shared.go             → BirthData, Subject, DateTimeLocation, DateRange, AstrologyOptions
option/option.go             → Functional options (RequestOption = func(*RequestConfig))
errors/error.go              → AstrologyError (StatusCode, Message, Code, Body)
internal/transport/          → AuthTransport, RetryTransport (http.RoundTripper chain)
internal/validator/          → Struct tag validation (validate:"required")
internal/apijson/            → Field[T] generic optional/nullable JSON type
```

HTTP transport chain: `DefaultTransport → RetryTransport → AuthTransport → http.Client`

## Coding Patterns

When generating code for this project, follow these patterns:

### New API method

```go
func (c *Client) GetThing(ctx context.Context, params ThingParams, opts ...option.RequestOption) (*ThingResponse, error) {
    var out ThingResponse
    if err := c.Post(ctx, c.BuildURL(apiPrefix, "thing"), params, &out, opts...); err != nil {
        return nil, err
    }
    return &out, nil
}
```

- First param: `ctx context.Context`
- Last param: `opts ...option.RequestOption`
- Use `c.Get()` / `c.Post()` / `c.Put()` / `c.Delete()` — never raw `MakeRequest`
- Build URLs with `c.BuildURL(apiPrefix, "segments"...)` — never string concatenation
- Declare `var out ResponseType` then return `&out`

### New param struct

```go
type ThingParams struct {
    Subject shared.Subject           `json:"subject" validate:"required"`
    Options *shared.AstrologyOptions `json:"options,omitempty"`
}
```

- Required fields: `validate:"required"` tag
- Optional struct fields: pointer type + `omitempty`

### New response type

```go
type ThingResponse map[string]any
```

### Unit test

```go
package <name>_test  // always black-box

func TestClient_GetThing(t *testing.T) {
    client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, "/api/v3/<name>/thing", r.URL.Path)
        assert.Equal(t, http.MethodPost, r.Method)
        testutil.JSON(w, testutil.DataEnvelope(map[string]any{"result": "ok"}))
    })
    defer cleanup()

    result, err := client.<Name>.GetThing(ctx, params)
    require.NoError(t, err)
    assert.NotNil(t, result)
}
```

- Use `testutil.NewClient(t, handler)` for mocks, `testutil.NewIntegrationClient(t)` for real API
- Use `testutil.DataEnvelope()` / `testutil.ResultEnvelope()` for response wrapping
- Use `testutil.DefaultSubject()` / `testutil.DefaultDateTimeLocation()` for fixtures
- Mock-only tests: `if testutil.IsIntegration() { t.Skip("...") }`

## Import Aliases

```go
import astroapi "github.com/astro-api/astroapi-go"
import astroerrors "github.com/astro-api/astroapi-go/errors"
```

## Adding a New Category

1. Create `categories/<name>/` with 4 files: `<name>.go`, `params.go`, `responses.go`, `<name>_test.go`
2. Define `const apiPrefix = "api/v3/<name>"`
3. Client struct embeds `*categories.BaseCategoryClient`
4. Wire into `AstrologyClient` in `client.go` (add field + instantiate in `NewClient`)

## Versioning

Git tags following semver: `v{MAJOR}.{MINOR}.{PATCH}`

- PATCH (v0.1.0 → v0.1.1): bug fix, docs
- MINOR (v0.1.1 → v0.2.0): new category, new method (backwards-compatible)
- MAJOR (v0.2.0 → v1.0.0): breaking API change (removed/renamed types, changed signatures)

Release: `git tag v0.1.0 && git push origin v0.1.0` — CI creates GitHub Release + notifies Go Module Proxy.

Starting from v2+, module path must include major suffix (`github.com/astro-api/astroapi-go/v2`) and all internal imports must be updated.

## Commands

```bash
make build    # go build ./...
make test     # go test -race -count=1 -timeout 120s ./...
make lint     # golangci-lint run ./...
make fmt      # gofmt -w -s .
make vet      # go vet ./...
```
