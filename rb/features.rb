# ArcgisHubWorldCountriesGeneralized SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module ArcgisHubWorldCountriesGeneralizedFeatures
  def self.make_feature(name)
    case name
    when "base"
      ArcgisHubWorldCountriesGeneralizedBaseFeature.new
    when "test"
      ArcgisHubWorldCountriesGeneralizedTestFeature.new
    else
      ArcgisHubWorldCountriesGeneralizedBaseFeature.new
    end
  end
end
