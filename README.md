# ArcgisHubWorldCountriesGeneralized SDK

Simplified global country boundaries optimised for display at scales of 1:5,000,000 and broader

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About ArcGIS Hub - World Countries Generalized

The **World Countries Generalized** layer is a hosted ArcGIS feature service published by [Esri](https://www.esri.com/) on ArcGIS Hub. It exposes a simplified, low-detail polygon representation of every country in the world, intended as a lightweight basemap or join layer for global visualisations.

The geometry is generalised for performance at small scales (roughly 1:5,000,000 and broader), so it draws quickly even when the whole world is in view. Esri compiles the data from [Garmin International](https://www.garmin.com/), the [CIA World Factbook](https://www.cia.gov/the-world-factbook/), and the [National Geographic Society](https://www.nationalgeographic.org/), and refreshes it annually to track changes to country names and significant borders.

What you typically get back:

- Country polygon geometries returned as GeoJSON
- Either a filtered subset (search by attribute) or the full global dataset in one request

Operational notes: the service is reachable over HTTPS with CORS enabled, so it can be queried directly from a browser. Returning the entire world in a single call is noticeably slower than a targeted query — community monitoring reports average response times around 6 seconds for the full dataset versus under half a second for a filtered search.

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

## 30-second quickstart

### TypeScript

```ts
import { ArcgisHubWorldCountriesGeneralizedSDK } from 'arcgis-hub-world-countries-generalized'

const client = new ArcgisHubWorldCountriesGeneralizedSDK({})

// List all features
const features = await client.Feature().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

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
| **Feature** | Country boundary polygons served as GeoJSON features; query a single country or retrieve the full world dataset from the FeatureServer layer. | `/0/query` |
| **Metadata** | Service- and layer-level metadata describing the feature service, including field definitions, geometry type, and capabilities exposed by the ArcGIS REST endpoint. | `/0` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from arcgishubworldcountriesgeneralized_sdk import ArcgisHubWorldCountriesGeneralizedSDK

client = ArcgisHubWorldCountriesGeneralizedSDK({})

# List all features
features, err = client.Feature(None).list(None, None)
```

### PHP

```php
<?php
require_once 'arcgishubworldcountriesgeneralized_sdk.php';

$client = new ArcgisHubWorldCountriesGeneralizedSDK([]);

// List all features
[$features, $err] = $client->Feature(null)->list(null, null);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go"

client := sdk.NewArcgisHubWorldCountriesGeneralizedSDK(map[string]any{})

// List all features
features, err := client.Feature(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "ArcgisHubWorldCountriesGeneralized_sdk"

client = ArcgisHubWorldCountriesGeneralizedSDK.new({})

# List all features
features, err = client.Feature(nil).list(nil, nil)
```

### Lua

```lua
local sdk = require("arcgis-hub-world-countries-generalized_sdk")

local client = sdk.new({})

-- List all features
local features, err = client:Feature(nil):list(nil, nil)
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
client = ArcgisHubWorldCountriesGeneralizedSDK.test(None, None)
result, err = client.Feature(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = ArcgisHubWorldCountriesGeneralizedSDK::test(null, null);
[$result, $err] = $client->Feature(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Feature(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = ArcgisHubWorldCountriesGeneralizedSDK.test(nil, nil)
result, err = client.Feature(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Feature(nil):load(
  { id = "test01" }, nil
)
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

## Using the ArcGIS Hub - World Countries Generalized

- Upstream: [https://services.arcgis.com/P3ePLMYs2RVChkJx/arcgis/rest/services/World_Countries_Generalized/FeatureServer](https://services.arcgis.com/P3ePLMYs2RVChkJx/arcgis/rest/services/World_Countries_Generalized/FeatureServer)
- API docs: [https://freepublicapis.com/arcgis-hub-world-countries-generalized](https://freepublicapis.com/arcgis-hub-world-countries-generalized)

---

Generated from the ArcGIS Hub - World Countries Generalized OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
