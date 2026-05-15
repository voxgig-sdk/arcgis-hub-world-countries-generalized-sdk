<?php
declare(strict_types=1);

// ArcgisHubWorldCountriesGeneralized SDK utility: prepare_body

class ArcgisHubWorldCountriesGeneralizedPrepareBody
{
    public static function call(ArcgisHubWorldCountriesGeneralizedContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
