package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewFeatureEntityFunc func(client *ArcgisHubWorldCountriesGeneralizedSDK, entopts map[string]any) ArcgisHubWorldCountriesGeneralizedEntity

var NewMetadataEntityFunc func(client *ArcgisHubWorldCountriesGeneralizedSDK, entopts map[string]any) ArcgisHubWorldCountriesGeneralizedEntity

