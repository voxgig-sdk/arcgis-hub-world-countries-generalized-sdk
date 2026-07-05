<?php
declare(strict_types=1);

// Typed models for the ArcgisHubWorldCountriesGeneralized SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Feature entity data model. */
class Feature
{
    public ?array $attribute = null;
    public ?array $geometry = null;
}

/** Request payload for Feature#list. */
class FeatureListMatch
{
    public ?array $attribute = null;
    public ?array $geometry = null;
}

/** Metadata entity data model. */
class Metadata
{
    public ?string $alia = null;
    public ?int $length = null;
    public ?string $name = null;
    public ?string $type = null;
}

/** Request payload for Metadata#list. */
class MetadataListMatch
{
    public ?string $alia = null;
    public ?int $length = null;
    public ?string $name = null;
    public ?string $type = null;
}

