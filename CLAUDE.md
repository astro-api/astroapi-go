# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`astroapi-go` is a pure Go HTTP client SDK for the Astrology API v3 (`https://api.astrology-api.io`). It is not a server or CLI — it's a library. The only external dependency is `testify` for test assertions; everything else uses the standard library.

**Module:** `github.com/astro-api/astroapi-go` | **Go version:** 1.24

## Common Commands

```bash
make build        # go build ./...
make test         # go test -race -count=1 -timeout 120s ./...
make test-cover   # test with coverage report
make lint         # golangci-lint run ./...
make fmt          # gofmt -w -s .
make vet          # go vet ./...
make tidy         # go mod tidy && go mod verify
```

Run a single test:
```bash
go test -race -run TestGetPositions ./categories/data/
```

Integration tests require `ASTROLOGY_API_KEY` env var (unit tests always run without it):
```bash
ASTROLOGY_API_KEY=your-key go test -race ./...
```

## Architecture

### Client Structure

`AstrologyClient` (in `client.go`) is the root. It holds sub-clients for each API category (Data, Charts, Analysis, etc.), all wired in `NewClient()`. Configuration comes from functional options (`option.RequestOption`) or environment variables (`ASTROLOGY_API_KEY`, `ASTROLOGY_API_BASE_URL`).

### Category Clients

Every category lives in `categories/<name>/` with exactly four files:
- `<name>.go` — `Client` struct + methods
- `params.go` — request parameter structs
- `responses.go` — response types (mostly `map[string]any` aliases)
- `<name>_test.go` — unit + integration tests

Each category client embeds `*categories.BaseCategoryClient` (defined in `categories/base.go`), which provides `Get`, `Post`, `Put`, `Delete`, `BuildURL`, and `MakeRequest`.

### HTTP Transport Chain

Built in `categories/base.go` via `buildHTTPClient`:
```
http.DefaultTransport → RetryTransport → AuthTransport → http.Client
```
Both `RetryTransport` (exponential backoff with body replay) and `AuthTransport` (Bearer token, Content-Type, Accept headers) implement `http.RoundTripper` and live in `internal/transport/`.

### Request Lifecycle

1. Method call (e.g. `client.Data.GetPositions(ctx, params, opts...)`)
2. `validator.Validate(params)` checks `validate:"required"` struct tags
3. Config cloned, per-request options applied
4. GET: params → query string via JSON round-trip + reflection; POST/PUT/DELETE: params → JSON body
5. Transport chain handles auth + retries
6. Response envelope (`{"data":...}` or `{"result":...}`) auto-unwrapped before JSON unmarshal
7. Non-2xx responses → `*astroerrors.AstrologyError`

### Key Patterns

- **Functional options:** `option.RequestOption` is `func(*requestconfig.RequestConfig)`. Options apply at client creation or per-request.
- **Generic Field type:** `internal/apijson.Field[T]` distinguishes absent vs null vs present. Exposed as `astroapi.Field[T]` with helpers `F()`, `Null[T]()`, `String()`, `Int()`, `Float()`, `Bool()`.
- **Nested sub-clients:** Only `insights/` has sub-clients (Relationship, Pet, Wellness, Financial, Business), each also embedding `*BaseCategoryClient`.

### Shared Types

`shared/shared.go` contains common request types used across categories: `BirthData`, `Subject`, `DateTimeLocation`, `DateRange`, `AstrologyOptions`, `ReportOptions`.

### Error Handling

`errors/error.go` defines `AstrologyError` with `StatusCode`, `Message`, `Code`, `Body` fields and helpers `IsNotFound()`, `IsRateLimit()`, `IsServerError()`, `IsClientError()`. Import as `astroerrors "github.com/astro-api/astroapi-go/errors"`.

## Testing

- **Unit tests** use `testutil.NewClient(t, mockHandler)` which creates `httptest.NewServer` with retries disabled.
- **Integration tests** use `testutil.NewIntegrationClient(t)` which skips if no API key. Named with `_Integration` suffix.
- Tests are black-box (`package <category>_test`).
- Test helpers in `internal/testutil/`: `JSON()`, `DataEnvelope()`, `ResultEnvelope()`, `DefaultSubject()`, `AssertNoError()`.

## Versioning & Releases

Versions are git tags following semver: `v{MAJOR}.{MINOR}.{PATCH}`

- **PATCH** (v0.1.0 → v0.1.1): bug fix, docs
- **MINOR** (v0.1.1 → v0.2.0): new category, new method (backwards-compatible)
- **MAJOR** (v0.2.0 → v1.0.0): breaking API change (removed/renamed types, changed signatures)

Release flow:
```bash
git tag v0.1.0
git push origin v0.1.0
```

CI (`release.yml`) automatically runs tests, creates GitHub Release with changelog, and pings Go Module Proxy.

Starting from v2+, the module path must include major suffix (`github.com/astro-api/astroapi-go/v2`) and all internal imports must be updated.

## Adding a New Category

1. Create `categories/<name>/` with the four standard files
2. Define `const apiPrefix = "api/v3/<name>"` and a `Client` struct embedding `*categories.BaseCategoryClient`
3. Add params with `json` + `validate` tags, response types, and methods using `c.Post(ctx, c.BuildURL(apiPrefix, "endpoint"), params, &out, opts...)`
4. Wire into `AstrologyClient` in `client.go`
5. Write tests using `testutil.NewClient` (unit) and `testutil.NewIntegrationClient` (integration)
