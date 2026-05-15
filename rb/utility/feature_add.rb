# ArcgisHubWorldCountriesGeneralized SDK utility: feature_add
module ArcgisHubWorldCountriesGeneralizedUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
