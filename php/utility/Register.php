<?php
declare(strict_types=1);

// ArcgisHubWorldCountriesGeneralized SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

ArcgisHubWorldCountriesGeneralizedUtility::setRegistrar(function (ArcgisHubWorldCountriesGeneralizedUtility $u): void {
    $u->clean = [ArcgisHubWorldCountriesGeneralizedClean::class, 'call'];
    $u->done = [ArcgisHubWorldCountriesGeneralizedDone::class, 'call'];
    $u->make_error = [ArcgisHubWorldCountriesGeneralizedMakeError::class, 'call'];
    $u->feature_add = [ArcgisHubWorldCountriesGeneralizedFeatureAdd::class, 'call'];
    $u->feature_hook = [ArcgisHubWorldCountriesGeneralizedFeatureHook::class, 'call'];
    $u->feature_init = [ArcgisHubWorldCountriesGeneralizedFeatureInit::class, 'call'];
    $u->fetcher = [ArcgisHubWorldCountriesGeneralizedFetcher::class, 'call'];
    $u->make_fetch_def = [ArcgisHubWorldCountriesGeneralizedMakeFetchDef::class, 'call'];
    $u->make_context = [ArcgisHubWorldCountriesGeneralizedMakeContext::class, 'call'];
    $u->make_options = [ArcgisHubWorldCountriesGeneralizedMakeOptions::class, 'call'];
    $u->make_request = [ArcgisHubWorldCountriesGeneralizedMakeRequest::class, 'call'];
    $u->make_response = [ArcgisHubWorldCountriesGeneralizedMakeResponse::class, 'call'];
    $u->make_result = [ArcgisHubWorldCountriesGeneralizedMakeResult::class, 'call'];
    $u->make_point = [ArcgisHubWorldCountriesGeneralizedMakePoint::class, 'call'];
    $u->make_spec = [ArcgisHubWorldCountriesGeneralizedMakeSpec::class, 'call'];
    $u->make_url = [ArcgisHubWorldCountriesGeneralizedMakeUrl::class, 'call'];
    $u->param = [ArcgisHubWorldCountriesGeneralizedParam::class, 'call'];
    $u->prepare_auth = [ArcgisHubWorldCountriesGeneralizedPrepareAuth::class, 'call'];
    $u->prepare_body = [ArcgisHubWorldCountriesGeneralizedPrepareBody::class, 'call'];
    $u->prepare_headers = [ArcgisHubWorldCountriesGeneralizedPrepareHeaders::class, 'call'];
    $u->prepare_method = [ArcgisHubWorldCountriesGeneralizedPrepareMethod::class, 'call'];
    $u->prepare_params = [ArcgisHubWorldCountriesGeneralizedPrepareParams::class, 'call'];
    $u->prepare_path = [ArcgisHubWorldCountriesGeneralizedPreparePath::class, 'call'];
    $u->prepare_query = [ArcgisHubWorldCountriesGeneralizedPrepareQuery::class, 'call'];
    $u->result_basic = [ArcgisHubWorldCountriesGeneralizedResultBasic::class, 'call'];
    $u->result_body = [ArcgisHubWorldCountriesGeneralizedResultBody::class, 'call'];
    $u->result_headers = [ArcgisHubWorldCountriesGeneralizedResultHeaders::class, 'call'];
    $u->transform_request = [ArcgisHubWorldCountriesGeneralizedTransformRequest::class, 'call'];
    $u->transform_response = [ArcgisHubWorldCountriesGeneralizedTransformResponse::class, 'call'];
});
