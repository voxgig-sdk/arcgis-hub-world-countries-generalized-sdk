# ArcgisHubWorldCountriesGeneralized SDK

ArcGIS Hub - World Countries Generalized client, generated from the OpenAPI spec.

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## Try it

**TypeScript**
```bash
npm install arcgis-hub-world-countries-generalized
```

**Python**
```bash
pip install arcgis-hub-world-countries-generalized-sdk
```

**PHP**
```bash
composer require voxgig/arcgis-hub-world-countries-generalized-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go
```

**Ruby**
```bash
gem install arcgis-hub-world-countries-generalized-sdk
```

**Lua**
```bash
luarocks install arcgis-hub-world-countries-generalized-sdk
```

## Quickstart

### TypeScript

```ts
import { ArcgisHubWorldCountriesGeneralizedSDK } from 'arcgis-hub-world-countries-generalized'

const client = new ArcgisHubWorldCountriesGeneralizedSDK({
  apikey: process.env.ARCGIS-HUB-WORLD-COUNTRIES-GENERALIZED_APIKEY,
})

// List all features
const features = await client.Feature().list()
console.log(features.data)
```

See the [TypeScript README](ts/README.md) for the full guide.

## Surfaces

| Surface | Path |
| --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | `go-cli/` |
| **MCP server** | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o arcgis-hub-world-countries-generalized-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "arcgis-hub-world-countries-generalized": {
      "command": "/abs/path/to/arcgis-hub-world-countries-generalized-mcp"
    }
  }
}
```

## Entities

The API exposes 2 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Feature** |  | `/0/query` |
| **Metadata** |  | `/0` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
import os
from arcgishubworldcountriesgeneralized_sdk import ArcgisHubWorldCountriesGeneralizedSDK

client = ArcgisHubWorldCountriesGeneralizedSDK({
    "apikey": os.environ.get("ARCGIS-HUB-WORLD-COUNTRIES-GENERALIZED_APIKEY"),
})

# List all features
features, err = client.Feature().list()
print(features)
```

### PHP

```php
<?php
require_once 'arcgishubworldcountriesgeneralized_sdk.php';

$client = new ArcgisHubWorldCountriesGeneralizedSDK([
    "apikey" => getenv("ARCGIS-HUB-WORLD-COUNTRIES-GENERALIZED_APIKEY"),
]);

// List all features
[$features, $err] = $client->Feature()->list();
print_r($features);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go"

client := sdk.NewArcgisHubWorldCountriesGeneralizedSDK(map[string]any{
    "apikey": os.Getenv("ARCGIS-HUB-WORLD-COUNTRIES-GENERALIZED_APIKEY"),
})

// List all features
features, err := client.Feature(nil).List(nil, nil)
fmt.Println(features)
```

### Ruby

```ruby
require_relative "ArcgisHubWorldCountriesGeneralized_sdk"

client = ArcgisHubWorldCountriesGeneralizedSDK.new({
  "apikey" => ENV["ARCGIS-HUB-WORLD-COUNTRIES-GENERALIZED_APIKEY"],
})

# List all features
features, err = client.Feature().list
puts features
```

### Lua

```lua
local sdk = require("arcgis-hub-world-countries-generalized_sdk")

local client = sdk.new({
  apikey = os.getenv("ARCGIS-HUB-WORLD-COUNTRIES-GENERALIZED_APIKEY"),
})

-- List all features
local features, err = client:Feature():list()
print(features)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = ArcgisHubWorldCountriesGeneralizedSDK.test()
const result = await client.Feature().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = ArcgisHubWorldCountriesGeneralizedSDK.test()
result, err = client.Feature().load({"id": "test01"})
```

### PHP

```php
$client = ArcgisHubWorldCountriesGeneralizedSDK::test();
[$result, $err] = $client->Feature()->load(["id" => "test01"]);
```

### Golang

```go
client := sdk.Test()
result, err := client.Feature(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = ArcgisHubWorldCountriesGeneralizedSDK.test
result, err = client.Feature().load({ "id" => "test01" })
```

### Lua

```lua
local client = sdk.test()
local result, err = client:Feature():load({ id = "test01" })
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

---

Generated from the ArcGIS Hub - World Countries Generalized OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
