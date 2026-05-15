package = "voxgig-sdk-arcgis-hub-world-countries-generalized"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk.git"
}
description = {
  summary = "ArcgisHubWorldCountriesGeneralized SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["arcgis-hub-world-countries-generalized_sdk"] = "arcgis-hub-world-countries-generalized_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
