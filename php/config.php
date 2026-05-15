<?php
declare(strict_types=1);

// ArcgisHubWorldCountriesGeneralized SDK configuration

class ArcgisHubWorldCountriesGeneralizedConfig
{
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "ArcgisHubWorldCountriesGeneralized",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
        ],
            ],
            "options" => [
                "base" => "https://services.arcgis.com/P3ePLMYs2RVChkJx/arcgis/rest/services/World_Countries_Generalized/FeatureServer",
                "auth" => [
                    "prefix" => "Bearer",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "feature" => [],
                    "metadata" => [],
                ],
            ],
            "entity" => [
        'feature' => [
          'fields' => [
            [
              'name' => 'attribute',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'geometry',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 1,
            ],
          ],
          'name' => 'feature',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'json',
                        'kind' => 'query',
                        'name' => 'f',
                        'orig' => 'f',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'geometry',
                        'orig' => 'geometry',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'geometry_type',
                        'orig' => 'geometry_type',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'order_by_field',
                        'orig' => 'order_by_field',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '*',
                        'kind' => 'query',
                        'name' => 'out_field',
                        'orig' => 'out_field',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 0,
                        'kind' => 'query',
                        'name' => 'result_offset',
                        'orig' => 'result_offset',
                        'reqd' => false,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                      [
                        'example' => 1000,
                        'kind' => 'query',
                        'name' => 'result_record_count',
                        'orig' => 'result_record_count',
                        'reqd' => false,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                      [
                        'example' => true,
                        'kind' => 'query',
                        'name' => 'return_geometry',
                        'orig' => 'return_geometry',
                        'reqd' => false,
                        'type' => '`$BOOLEAN`',
                        'active' => true,
                      ],
                      [
                        'example' => 'esriSpatialRelIntersects',
                        'kind' => 'query',
                        'name' => 'spatial_rel',
                        'orig' => 'spatial_rel',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '1=1',
                        'kind' => 'query',
                        'name' => 'where',
                        'orig' => 'where',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/0/query',
                  'parts' => [
                    '0',
                    'query',
                  ],
                  'select' => [
                    'exist' => [
                      'f',
                      'geometry',
                      'geometry_type',
                      'order_by_field',
                      'out_field',
                      'result_offset',
                      'result_record_count',
                      'return_geometry',
                      'spatial_rel',
                      'where',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'metadata' => [
          'fields' => [
            [
              'name' => 'alia',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'length',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'name',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'type',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'metadata',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'json',
                        'kind' => 'query',
                        'name' => 'f',
                        'orig' => 'f',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/0',
                  'parts' => [
                    '0',
                  ],
                  'select' => [
                    'exist' => [
                      'f',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return ArcgisHubWorldCountriesGeneralizedFeatures::make_feature($name);
    }
}
