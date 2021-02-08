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
