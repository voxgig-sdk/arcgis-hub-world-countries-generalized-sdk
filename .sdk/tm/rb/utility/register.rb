# ArcgisHubWorldCountriesGeneralized SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

ArcgisHubWorldCountriesGeneralizedUtility.registrar = ->(u) {
  u.clean = ArcgisHubWorldCountriesGeneralizedUtilities::Clean
  u.done = ArcgisHubWorldCountriesGeneralizedUtilities::Done
  u.make_error = ArcgisHubWorldCountriesGeneralizedUtilities::MakeError
  u.feature_add = ArcgisHubWorldCountriesGeneralizedUtilities::FeatureAdd
  u.feature_hook = ArcgisHubWorldCountriesGeneralizedUtilities::FeatureHook
  u.feature_init = ArcgisHubWorldCountriesGeneralizedUtilities::FeatureInit
  u.fetcher = ArcgisHubWorldCountriesGeneralizedUtilities::Fetcher
  u.make_fetch_def = ArcgisHubWorldCountriesGeneralizedUtilities::MakeFetchDef
  u.make_context = ArcgisHubWorldCountriesGeneralizedUtilities::MakeContext
  u.make_options = ArcgisHubWorldCountriesGeneralizedUtilities::MakeOptions
  u.make_request = ArcgisHubWorldCountriesGeneralizedUtilities::MakeRequest
  u.make_response = ArcgisHubWorldCountriesGeneralizedUtilities::MakeResponse
  u.make_result = ArcgisHubWorldCountriesGeneralizedUtilities::MakeResult
  u.make_point = ArcgisHubWorldCountriesGeneralizedUtilities::MakePoint
  u.make_spec = ArcgisHubWorldCountriesGeneralizedUtilities::MakeSpec
  u.make_url = ArcgisHubWorldCountriesGeneralizedUtilities::MakeUrl
  u.param = ArcgisHubWorldCountriesGeneralizedUtilities::Param
  u.prepare_auth = ArcgisHubWorldCountriesGeneralizedUtilities::PrepareAuth
  u.prepare_body = ArcgisHubWorldCountriesGeneralizedUtilities::PrepareBody
  u.prepare_headers = ArcgisHubWorldCountriesGeneralizedUtilities::PrepareHeaders
  u.prepare_method = ArcgisHubWorldCountriesGeneralizedUtilities::PrepareMethod
  u.prepare_params = ArcgisHubWorldCountriesGeneralizedUtilities::PrepareParams
  u.prepare_path = ArcgisHubWorldCountriesGeneralizedUtilities::PreparePath
  u.prepare_query = ArcgisHubWorldCountriesGeneralizedUtilities::PrepareQuery
  u.result_basic = ArcgisHubWorldCountriesGeneralizedUtilities::ResultBasic
  u.result_body = ArcgisHubWorldCountriesGeneralizedUtilities::ResultBody
  u.result_headers = ArcgisHubWorldCountriesGeneralizedUtilities::ResultHeaders
  u.transform_request = ArcgisHubWorldCountriesGeneralizedUtilities::TransformRequest
  u.transform_response = ArcgisHubWorldCountriesGeneralizedUtilities::TransformResponse
}
