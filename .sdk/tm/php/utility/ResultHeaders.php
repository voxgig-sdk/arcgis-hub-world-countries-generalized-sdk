<?php
declare(strict_types=1);

// ArcgisHubWorldCountriesGeneralized SDK utility: result_headers

class ArcgisHubWorldCountriesGeneralizedResultHeaders
{
    public static function call(ArcgisHubWorldCountriesGeneralizedContext $ctx): ?ArcgisHubWorldCountriesGeneralizedResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
