# frozen_string_literal: true

# Typed models for the ArcgisHubWorldCountriesGeneralized SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Feature entity data model.
#
# @!attribute [rw] attribute
#   @return [Hash, nil]
#
# @!attribute [rw] geometry
#   @return [Hash, nil]
Feature = Struct.new(
  :attribute,
  :geometry,
  keyword_init: true
)

# Match filter for Feature#list (any subset of Feature fields).
#
# @!attribute [rw] attribute
#   @return [Hash, nil]
#
# @!attribute [rw] geometry
#   @return [Hash, nil]
FeatureListMatch = Struct.new(
  :attribute,
  :geometry,
  keyword_init: true
)

# Metadata entity data model.
#
# @!attribute [rw] alia
#   @return [String, nil]
#
# @!attribute [rw] length
#   @return [Integer, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
Metadata = Struct.new(
  :alia,
  :length,
  :name,
  :type,
  keyword_init: true
)

# Match filter for Metadata#list (any subset of Metadata fields).
#
# @!attribute [rw] alia
#   @return [String, nil]
#
# @!attribute [rw] length
#   @return [Integer, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
MetadataListMatch = Struct.new(
  :alia,
  :length,
  :name,
  :type,
  keyword_init: true
)

