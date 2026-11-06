# Orders API Client — Go Learning Project

A Go HTTP-client exercise with a small Gin adapter. It loads product data locally, accepts an order request, resolves product names, and forwards the resulting payload to a separate orders API.

## Request flow

```text
POST localhost:8080/order
        → product lookup in data.json
        → Go API client
        → POST localhost:8085/orders
```

This repository is the client/adapter; it does not include the upstream orders service. The original project references [Demo Orders API](https://github.com/paranietharan/demo-orders-api-go-lang); access to that separate repository may be required to reproduce the original setup.

## Run locally

Use Go 1.24.5 or a compatible newer toolchain, as declared in [go.mod](go.mod).

```bash
git clone https://github.com/paranietharan/api-client.git
cd api-client
go mod download
go run .
```

Before sending an order, run a compatible upstream service accepting `POST /orders` at `http://localhost:8085`. Change `BaseUrl` in [helper/helper.go](helper/helper.go) if it runs elsewhere. Run this client from the repository root so it can read [data.json](data.json); the catalog must contain at least one item.

Example request to the local adapter:

```bash
curl -X POST http://localhost:8080/order   -H 'Content-Type: application/json'   -d '{"id":"demo-order","item":[{"name":"Burger","qty":"1"}]}'
```

The JSON field is `item` (singular), and product names must match the catalog. The current adapter looks up and forwards catalog items but does not copy the submitted `qty` into those items. Quantity handling is therefore a known limitation of this learning implementation.

## Source guide

- [main.go](main.go): Gin route and order-to-catalog mapping.
- [api](api): HTTP client, request/query building, and order submission.
- [dto/order.go](dto/order.go): Request and response structures.
- [data/data.go](data/data.go): Load and search the product catalog.
- [helper/helper.go](helper/helper.go): Upstream address and submission helper.

## Development checks

```bash
go build ./...
go test ./...
```

There are currently no `*_test.go` files, so the test command is not evidence of business-flow coverage. Request dumping is enabled in the submission helper; use dummy data for demonstrations.
