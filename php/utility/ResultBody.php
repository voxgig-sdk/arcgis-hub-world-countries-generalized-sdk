<?php
declare(strict_types=1);

// ArcgisHubWorldCountriesGeneralized SDK utility: result_body

class ArcgisHubWorldCountriesGeneralizedResultBody
{
    public static function call(ArcgisHubWorldCountriesGeneralizedContext $ctx): ?ArcgisHubWorldCountriesGeneralizedResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
