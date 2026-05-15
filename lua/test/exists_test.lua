-- ProjectName SDK exists test

local sdk = require("arcgis-hub-world-countries-generalized_sdk")

describe("ArcgisHubWorldCountriesGeneralizedSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
