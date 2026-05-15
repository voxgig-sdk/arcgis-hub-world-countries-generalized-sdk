# ProjectName SDK exists test

import pytest
from arcgishubworldcountriesgeneralized_sdk import ArcgisHubWorldCountriesGeneralizedSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = ArcgisHubWorldCountriesGeneralizedSDK.test(None, None)
        assert testsdk is not None
