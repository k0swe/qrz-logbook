[![PkgGoDev](https://pkg.go.dev/badge/github.com/k0swe/qrz-logbook)](https://pkg.go.dev/github.com/k0swe/qrz-logbook)
[![Go Report Card](https://goreportcard.com/badge/github.com/k0swe/qrz-logbook)](https://goreportcard.com/report/github.com/k0swe/qrz-logbook)

# Go API client for QRZ.com Logbook

A GoLang client library for QRZ.com's Logbook data service. The service 
provides real-time access to information from QRZ.com logbooks. The API is
documented [here](https://www.qrz.com/docs/logbook/QRZLogbookAPI.html).

This client library was generated based on the OpenAPI specification in
the `api/openapi.yaml` file. However, the API itself is not well-described by
OpenAPI, so the generated library is supplemented with `wrapper.go`.

A simple application to demonstrate how to integrate the library is located in
`cmd/qrz-logbook/main.go`.
 
A QRZ.com XML subscription is required to take full advantage of the API. A
description of subscription plans and rates is available on the 
[QRZ.com website](http://www.qrz.com/i/subscriptions.html).
