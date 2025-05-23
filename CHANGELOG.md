# 0.18.2 (May 13, 2025)
* Added support for `jsonapi` `links` and `meta` for list responses.
* Upgraded to latest `github.com/hashicorp/jsonapi` package.

# 0.18.1 (Sep 27, 2024)
* Use github.com/google/go-cmp to test response body.

# 0.18.0 (Aug 30, 2024)
* Switch to standard `slog` package for logging.
* Added a `slog.Logger` to `http.Request.Context()`.
* Added a `slog.Logger` to `json.Request`.

# 0.17.2 (Apr 02, 2024)
* PanicMiddleware no longer writes response code (to avoid panic from duplicate write).

# 0.17.1 (Mar 11, 2024)
* PanicMiddleware prints stack trace to stderr.

# 0.17.0 (Mar 11, 2024)
* Configured `http.Server` with no timeouts.
* Added `AdjustServerFn` to `api.Server` that allows configuration of the underlying `http.Server`.

# 0.16.1 (Mar 07, 2024)
* Added `errors.CaptureMiddleware` to report API errors.
* Added macros to `errors` for creating path/querystring/payload parse errors.
* Added request parsers that automate parsing and error wrapping.

# 0.16.0 (Mar 06, 2024)
* Improved `recovery.PanicMiddleware` to inject `PanicError` (adheres to `error`) instead of `interface{}`.
* `PanicError` contains stack trace as `StackTrace() []byte`.
* Added `jwt.ClaimsMiddleware[T]` which retrieves JWT token from request, then parses the token into custom claims.

# 0.15.0 (Mar 04, 2024)
* Reimplemented `intercept.Middleware` to use [`httpsnoop`](https://github.com/felixge/httpsnoop).

# 0.14.1 (Mar 01, 2024)
* Added `MiddlewareChain` to create a set of middlewares without using a router. This is helpful to apply middleware to a single HTTP handler.

# 0.14.0 (Jan 24, 2024)
* Added `ErrorCode` to `ApiError` to emit application-specific error codes.
* Changed `BadRequestError.Details` from `[]string` to `map[string]string`

# 0.13.1 (Jun 21, 2023)
* Added `errors.ObscureInteralErrorsMiddleware` to prevent sensitive errors from reaching users.
  * A developer can optionally log the original error message to logs.

## 0.13.0 (Jun 21, 2023)
* Refactor `ValidationErrors` to a slice for easier interaction.
* Added helpers to `ValidationErrors` to swap between map and slice.

## 0.12.5 (Dec 16, 2022)

* Reporting json encoding errors to logs when sending json response.

## 0.12.4 (Nov 29, 2022)

* Fixed `ApiError` from hiding injected `Err`.

## 0.12.3 (Oct 20, 2022)

* Added `NewApiError` and `NewBadRequestError` for easier error construction.

## 0.12.2 (Oct 20, 2022)

* Improved formatting of `errors.BadRequestError`.

## 0.12.1 (Sep 14, 2022)

* Added support for `errors` package in `jsonapi` package.

## 0.12.0 (Sep 14, 2022)

* Aligned `SendError(error)` in `jsonapi` package with `json` package.

## 0.11.0 (Feb 10, 2022)

* Added `captureBody` toggle on `intercept.Middleware` to prevent capturing body if only intercepting status code/response duration.

## 0.10.1 (Dec 09, 2021)

* Added `http.Hijacker` pass-through for all response writers (including `intercept.Middleware`).

## 0.10.0 (Oct 14, 2021)

* Added `LaunchTLS` to `Server`.

## 0.9.0 (Aug 18, 2021)

* Added `errors` package for default error formats.

## 0.8.0 (Aug 09, 2021)

* Replace panic middleware with custom and added user-supplied fn.

## 0.7.3 (Feb 08, 2021)

* Fixed CORS preflight from responding with 405 (Method Not Allowed) when a route was not registered.

## 0.7.2 (Feb 05, 2021)

* Added `logging.FallbackBehavior`.

## 0.7.1 (Feb 05, 2021)

* Fix busted go version.

## 0.7.0 (Feb 05, 2021)

* Replaced logging middleware with `intercept.Middleware` and logging macros.

## 0.6.2 (Nov 27, 2020)

* Fixed emission of errors through `jsonapi` interfaces.

## 0.6.1 (Nov 21, 2020)

* Added `jwt.Middleware()` to parse, but not verify incoming jwt token in `Authorization` header as `Bearer` token.

## 0.6.0 (Nov 21, 2020)

* Add `jsonapi` package to support JSONAPI.

## 0.5.5 (Nov 21, 2020)

* Fixed `EndpointGroup` prefix.

## 0.5.4 (Nov 20, 2020)

* Added `EndpointGroup` to locate several endpoints under a single prefixed path.

## 0.5.3 (Oct 14, 2020)

* Fixed bug in `AuthenticateMiddleware` prohibiting upstream proxying when error occurs.

## 0.5.2 (Oct 13, 2020)

* Added `auth` package with `AuthenticateMiddleware` and `AuthorizeMiddleware`.

## 0.5.0 (Oct 08, 2020)

* Added support for `Access-Control-Allow-Credentials` for cors middleware.
  * By default, this is set to `true`

## 0.4.3 (Aug 10, 2020)

* Amended endpoint logger to log outgoing data for 4xx, 5xx responses

## 0.4.2 (Aug 10, 2020)

* Providing default CORS setup that works properly
* Providing mechanism to flow NotFound and MethodNotAllowed through middlewares
* Setting up `app1` example to log NotFound and MethodNotAllowed and utilize default CORS setup

## 0.4.1 (Jul 14, 2020)

* Added `standard.NewEnpoint` helper to make concise syntax.

## 0.4.0 (Jul 14, 2020)

* Rebuilt entire library to be simpler and not obscure golang http constructs.
* Added `gzip`, `recovery` middleware
* Added `options` endpoint

## 0.3.0 (Oct 16, 2018)

* BREAKING: Removed `Unlogged()` from endpoints.
* Added support for middleware
* Added endpoint logger to migrate endpoint logging to middleware.

## 0.2.5 (Sep 18, 2018)

* Added support for reading query string from request.

## 0.2.4 (Sep 14, 2018)

* Added `Raw.Endpoint`, allowing a `http.HandlerFunc` to be used as an `Endpoint`. 

## 0.2.3 (Aug 20, 2018)

* Added `Endpoints` with a helper `Print`.

## 0.2.2 (Aug 15, 2018)

* Added ability to not log requests to an endpoint. Useful for health endpoints that are repeatedly hit.

## 0.2.1 (Aug 11, 2018)

* Using t.Run to run tests.

## 0.2.0 (Jul 27, 2018)

* Rebuilt api server contextual wrapping.

## 0.1.3 (Jun 19, 2018)

* Fixed response header `Content-Type` to `application/json`.

## 0.1.2 (May 18, 2018)

* Refactoring logging of request.

## 0.1.1 (May 18, 2018)

FEATURES:

* Logging duration of request.

## 0.1.0 (Apr 25, 2018)

FEATURES:

* Initial implementation of go-api.
* Implemented json endpoints.
