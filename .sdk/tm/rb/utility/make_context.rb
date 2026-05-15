# ArcgisHubWorldCountriesGeneralized SDK utility: make_context
require_relative '../core/context'
module ArcgisHubWorldCountriesGeneralizedUtilities
  MakeContext = ->(ctxmap, basectx) {
    ArcgisHubWorldCountriesGeneralizedContext.new(ctxmap, basectx)
  }
end
