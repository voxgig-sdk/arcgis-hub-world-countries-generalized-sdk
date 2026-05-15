# ArcgisHubWorldCountriesGeneralized SDK feature factory

from feature.base_feature import ArcgisHubWorldCountriesGeneralizedBaseFeature
from feature.test_feature import ArcgisHubWorldCountriesGeneralizedTestFeature


def _make_feature(name):
    features = {
        "base": lambda: ArcgisHubWorldCountriesGeneralizedBaseFeature(),
        "test": lambda: ArcgisHubWorldCountriesGeneralizedTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
