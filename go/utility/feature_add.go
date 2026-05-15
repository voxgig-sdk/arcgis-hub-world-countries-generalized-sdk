package utility

import "github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/core"

func featureAddUtil(ctx *core.Context, f core.Feature) {
	client := ctx.Client
	features := client.Features

	client.Features = append(features, f)
}
