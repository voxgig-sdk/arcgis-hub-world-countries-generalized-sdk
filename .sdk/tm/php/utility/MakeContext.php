<?php
declare(strict_types=1);

// ArcgisHubWorldCountriesGeneralized SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class ArcgisHubWorldCountriesGeneralizedMakeContext
{
    public static function call(array $ctxmap, ?ArcgisHubWorldCountriesGeneralizedContext $basectx): ArcgisHubWorldCountriesGeneralizedContext
    {
        return new ArcgisHubWorldCountriesGeneralizedContext($ctxmap, $basectx);
    }
}
