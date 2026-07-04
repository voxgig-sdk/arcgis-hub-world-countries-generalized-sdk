# Typed models for the ArcgisHubWorldCountriesGeneralized SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Feature:
    attribute: Optional[dict] = None
    geometry: Optional[dict] = None


@dataclass
class FeatureListMatch:
    attribute: Optional[dict] = None
    geometry: Optional[dict] = None


@dataclass
class Metadata:
    alia: Optional[str] = None
    length: Optional[int] = None
    name: Optional[str] = None
    type: Optional[str] = None


@dataclass
class MetadataListMatch:
    alia: Optional[str] = None
    length: Optional[int] = None
    name: Optional[str] = None
    type: Optional[str] = None

