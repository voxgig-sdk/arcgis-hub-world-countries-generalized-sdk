<?php
declare(strict_types=1);

// ArcgisHubWorldCountriesGeneralized SDK exists test

require_once __DIR__ . '/../arcgishubworldcountriesgeneralized_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = ArcgisHubWorldCountriesGeneralizedSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
