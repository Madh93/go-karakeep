# go-karakeep

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue)](https://go.dev/doc/install)
[![Latest tag](https://img.shields.io/github/v/tag/Madh93/go-karakeep?label=go%20module)](https://github.com/Madh93/go-karakeep/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/Madh93/go-karakeep.svg)](https://pkg.go.dev/github.com/Madh93/go-karakeep)
[![License](https://img.shields.io/badge/License-MIT-brightgreen)](LICENSE)

[Go](https://go.dev/) client library for [Karakeep](https://karakeep.app).

The `go-karakeep` client is auto-generated using the
[`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen) tool, which allows convert [Karakeep OpenAPI](https://github.com/karakeep-app/karakeep/blob/v0.24.1/packages/open-api/karakeep-openapi-spec.json) specification to Go code.

## Requirements

- [Go 1.23](https://golang.org/dl/) or higher.
- A valid API key from [Karakeep](https://docs.karakeep.app/screenshots#settings).

## Installation

To install `go-karakeep`, use `go get`:

```sh
go get github.com/Madh93/go-karakeep
```

## Usage

Here is a basic example of how to use the `go-karakeep` library:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/Madh93/go-karakeep"
)

func main() {
    // Basic configuration
    apiUrl := "https://<YOUR_KARAKEEP_HOSTNAME>/api/v1" // Replace this with your API URL
    apiKey := "<YOUR_KARAKEEP_API_KEY>"                 // Replace this with your actual token

    // Set up Bearer authentication
    auth := func(ctx context.Context, req *http.Request) error {
        req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
        return nil
    }

    // Create the Karakeep client
    client, err := karakeep.NewClient(apiUrl, karakeep.WithRequestEditorFn(auth))
    if err != nil {
        log.Fatalf("Error creating client: %v", err)
    }

    log.Printf("Hello world from %s", client.Server)
}
```

For more code examples, check out the [examples](examples) directory.

## Documentation

For detailed usage and API documentation, refer to the [GoDoc](https://pkg.go.dev/github.com/Madh93/go-karakeep).

Additionally, it's recommended to check out the latest [Karakeep API documentation](https://docs.karakeep.app/api) for more information.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any bug fixes or enhancements.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a Pull Request.

## License

This project is licensed under the [MIT license](LICENSE).
