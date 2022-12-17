# Squeeze Go Client

This project implements an HTTP client for the Squeeze API.

**Maturity**: Experimental. We have already created multiple HTTP clients internally but never a fully featured
client which supported all API endpoints of Squeeze. We will try to generate an API client based on the OpenAPI
specification once the [OpenAPI generator](https://openapi-generator.tech/) is able to generate a fully working
client.

**Support**: Currently, we will not officially support this client. Still, if you want to use the client and have
problems or questions, feel free to open an issue. We would love to encourage usage of this client and the HTTP
api of our products in general.

## Versioning

Squeeze has multiple APIs (and will probably receive more APIs in the future). This client is versioned to match
the versions of the API.

- The v1 API `/api/v1` will not be supported
- The v2 API `/api/v2` will be supported by this client. The `main` branch is used for this.
- Git tags will be used to mark client releases matching to the Squeeze API. If you use Squeeze 2.3.1, you must also
use the client version 2.3.1.

## Usage

Add client as dependency:

```go
go get github.com/dexpro-solutions-gmbh/squeeze-go-client
```

### Authentication

Currently, only authentication via API key is supported. We will support for JWT based authentication
in the future.

```go
client := squeeze_go_client.NewClient("https://your.squeeze.one/api/v2")
client.ApiKey = "..."
```

## Tests

To run tests on this project, use these ENV variables to configure the Squeeze API to be used when testing:

- SQZ_BASE_PATH: The base path of the Squeeze API. Example: `https://your.squeeze.one/api/v2`
- SQZ_KEY: API key to authenticate with
