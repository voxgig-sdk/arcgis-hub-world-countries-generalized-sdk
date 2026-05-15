<?php
declare(strict_types=1);

// ArcgisHubWorldCountriesGeneralized SDK utility: feature_add

class ArcgisHubWorldCountriesGeneralizedFeatureAdd
{
    public static function call(ArcgisHubWorldCountriesGeneralizedContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
