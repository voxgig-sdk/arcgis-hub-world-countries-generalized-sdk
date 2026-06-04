# ArcgisHubWorldCountriesGeneralized SDK configuration


def make_config():
    return {
        "main": {
            "name": "ArcgisHubWorldCountriesGeneralized",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://services.arcgis.com/P3ePLMYs2RVChkJx/arcgis/rest/services/World_Countries_Generalized/FeatureServer",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "feature": {},
                "metadata": {},
            },
        },
        "entity": {
      "feature": {
        "fields": [
          {
            "name": "attribute",
            "req": False,
            "type": "`$OBJECT`",
            "active": True,
            "index$": 0,
          },
          {
            "name": "geometry",
            "req": False,
            "type": "`$OBJECT`",
            "active": True,
            "index$": 1,
          },
        ],
        "name": "feature",
        "op": {
          "list": {
            "name": "list",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "example": "json",
                      "kind": "query",
                      "name": "f",
                      "orig": "f",
                      "reqd": False,
                      "type": "`$STRING`",
                      "active": True,
                    },
                    {
                      "kind": "query",
                      "name": "geometry",
                      "orig": "geometry",
                      "reqd": False,
                      "type": "`$STRING`",
                      "active": True,
                    },
                    {
                      "kind": "query",
                      "name": "geometry_type",
                      "orig": "geometry_type",
                      "reqd": False,
                      "type": "`$STRING`",
                      "active": True,
                    },
                    {
                      "kind": "query",
                      "name": "order_by_field",
                      "orig": "order_by_field",
                      "reqd": False,
                      "type": "`$STRING`",
                      "active": True,
                    },
                    {
                      "example": "*",
                      "kind": "query",
                      "name": "out_field",
                      "orig": "out_field",
                      "reqd": False,
                      "type": "`$STRING`",
                      "active": True,
                    },
                    {
                      "example": 0,
                      "kind": "query",
                      "name": "result_offset",
                      "orig": "result_offset",
                      "reqd": False,
                      "type": "`$INTEGER`",
                      "active": True,
                    },
                    {
                      "example": 1000,
                      "kind": "query",
                      "name": "result_record_count",
                      "orig": "result_record_count",
                      "reqd": False,
                      "type": "`$INTEGER`",
                      "active": True,
                    },
                    {
                      "example": True,
                      "kind": "query",
                      "name": "return_geometry",
                      "orig": "return_geometry",
                      "reqd": False,
                      "type": "`$BOOLEAN`",
                      "active": True,
                    },
                    {
                      "example": "esriSpatialRelIntersects",
                      "kind": "query",
                      "name": "spatial_rel",
                      "orig": "spatial_rel",
                      "reqd": False,
                      "type": "`$STRING`",
                      "active": True,
                    },
                    {
                      "example": "1=1",
                      "kind": "query",
                      "name": "where",
                      "orig": "where",
                      "reqd": False,
                      "type": "`$STRING`",
                      "active": True,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/0/query",
                "parts": [
                  "0",
                  "query",
                ],
                "select": {
                  "exist": [
                    "f",
                    "geometry",
                    "geometry_type",
                    "order_by_field",
                    "out_field",
                    "result_offset",
                    "result_record_count",
                    "return_geometry",
                    "spatial_rel",
                    "where",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "active": True,
                "index$": 0,
              },
            ],
            "input": "data",
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "metadata": {
        "fields": [
          {
            "name": "alia",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 0,
          },
          {
            "name": "length",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 1,
          },
          {
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 2,
          },
          {
            "name": "type",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 3,
          },
        ],
        "name": "metadata",
        "op": {
          "list": {
            "name": "list",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "example": "json",
                      "kind": "query",
                      "name": "f",
                      "orig": "f",
                      "reqd": False,
                      "type": "`$STRING`",
                      "active": True,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/0",
                "parts": [
                  "0",
                ],
                "select": {
                  "exist": [
                    "f",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "active": True,
                "index$": 0,
              },
            ],
            "input": "data",
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
