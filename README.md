# Go Template

## Overview

This repository contains a Go template for building microservices using gRPC. It uses Uber's FX for dependency injection and buf.build and connect.build to handle gRPC requests efficiently.

## Directory Structure

The project is organized as follows:

```
go-template/
├── api/
│   └── auth/
│       └── v1/
│           └── auth.proto
└── internal/
    ├── auth/
    │   └── router/
    ├── config/
    ├── logger/
    └── server/
```

`api/` : Contains protobuf files, such as `go-template/api/auth/v1/auth.proto`

`internal/`: Contains internal packages like auth, config, logger, and server

## Build and Development

`make protogen`: Runs `buf.build`'s generator to create files for the Go server and the TypeScript client.

`make watch`: Utilizes `air` to run the development server, with live reloading.

## Packages and Routing

Each package can have its own router, contributing to the main router defined in the `server` package. Each service handler needs to be added in `go-template/internal/server/routes.go`

`internal/server/config.go` defines environment variables used in the code

## Environment variables and configuration
`internal/config/routes.go` defines route handlers and gRPC reflection
