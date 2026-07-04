// Typed models for the ArcgisHubWorldCountriesGeneralized SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Feature {
  attribute?: Record<string, any>
  geometry?: Record<string, any>
}

export type FeatureListMatch = Partial<Feature>

export interface Metadata {
  alia?: string
  length?: number
  name?: string
  type?: string
}

export type MetadataListMatch = Partial<Metadata>

