-- ArcgisHubWorldCountriesGeneralized SDK error

local ArcgisHubWorldCountriesGeneralizedError = {}
ArcgisHubWorldCountriesGeneralizedError.__index = ArcgisHubWorldCountriesGeneralizedError


function ArcgisHubWorldCountriesGeneralizedError.new(code, msg, ctx)
  local self = setmetatable({}, ArcgisHubWorldCountriesGeneralizedError)
  self.is_sdk_error = true
  self.sdk = "ArcgisHubWorldCountriesGeneralized"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function ArcgisHubWorldCountriesGeneralizedError:error()
  return self.msg
end


function ArcgisHubWorldCountriesGeneralizedError:__tostring()
  return self.msg
end


return ArcgisHubWorldCountriesGeneralizedError
