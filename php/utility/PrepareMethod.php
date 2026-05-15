<?php
declare(strict_types=1);

// ArcgisHubWorldCountriesGeneralized SDK utility: prepare_method

class ArcgisHubWorldCountriesGeneralizedPrepareMethod
{
    private const METHOD_MAP = [
        'create' => 'POST',
        'update' => 'PUT',
        'load' => 'GET',
        'list' => 'GET',
        'remove' => 'DELETE',
        'patch' => 'PATCH',
    ];

    public static function call(ArcgisHubWorldCountriesGeneralizedContext $ctx): string
    {
        return self::METHOD_MAP[$ctx->op->name] ?? 'GET';
    }
}
