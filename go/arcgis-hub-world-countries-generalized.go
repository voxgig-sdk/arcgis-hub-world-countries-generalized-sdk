package voxgigarcgishubworldcountriesgeneralizedsdk

import (
	"github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go/core"
	"github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go/entity"
	"github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go/feature"
	_ "github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go/utility"
)

// Type aliases preserve external API.
type ArcgisHubWorldCountriesGeneralizedSDK = core.ArcgisHubWorldCountriesGeneralizedSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type ArcgisHubWorldCountriesGeneralizedEntity = core.ArcgisHubWorldCountriesGeneralizedEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type ArcgisHubWorldCountriesGeneralizedError = core.ArcgisHubWorldCountriesGeneralizedError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewFeatureEntityFunc = func(client *core.ArcgisHubWorldCountriesGeneralizedSDK, entopts map[string]any) core.ArcgisHubWorldCountriesGeneralizedEntity {
		return entity.NewFeatureEntity(client, entopts)
	}
	core.NewMetadataEntityFunc = func(client *core.ArcgisHubWorldCountriesGeneralizedSDK, entopts map[string]any) core.ArcgisHubWorldCountriesGeneralizedEntity {
		return entity.NewMetadataEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewArcgisHubWorldCountriesGeneralizedSDK = core.NewArcgisHubWorldCountriesGeneralizedSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewArcgisHubWorldCountriesGeneralizedSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *ArcgisHubWorldCountriesGeneralizedSDK  { return NewArcgisHubWorldCountriesGeneralizedSDK(nil) }
func Test() *ArcgisHubWorldCountriesGeneralizedSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
