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
