package core

func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "ArcgisHubWorldCountriesGeneralized",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://services.arcgis.com/P3ePLMYs2RVChkJx/arcgis/rest/services/World_Countries_Generalized/FeatureServer",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"feature": map[string]any{},
				"metadata": map[string]any{},
			},
		},
		"entity": map[string]any{
			"feature": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attribute",
						"req": false,
						"type": "`$OBJECT`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "geometry",
						"req": false,
						"type": "`$OBJECT`",
						"active": true,
						"index$": 1,
					},
				},
				"name": "feature",
				"op": map[string]any{
					"list": map[string]any{
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "json",
											"kind": "query",
											"name": "f",
											"orig": "f",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"kind": "query",
											"name": "geometry",
											"orig": "geometry",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"kind": "query",
											"name": "geometry_type",
											"orig": "geometry_type",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"kind": "query",
											"name": "order_by_field",
											"orig": "order_by_field",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "*",
											"kind": "query",
											"name": "out_field",
											"orig": "out_field",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "result_offset",
											"orig": "result_offset",
											"reqd": false,
											"type": "`$INTEGER`",
											"active": true,
										},
										map[string]any{
											"example": 1000,
											"kind": "query",
											"name": "result_record_count",
											"orig": "result_record_count",
											"reqd": false,
											"type": "`$INTEGER`",
											"active": true,
										},
										map[string]any{
											"example": true,
											"kind": "query",
											"name": "return_geometry",
											"orig": "return_geometry",
											"reqd": false,
											"type": "`$BOOLEAN`",
											"active": true,
										},
										map[string]any{
											"example": "esriSpatialRelIntersects",
											"kind": "query",
											"name": "spatial_rel",
											"orig": "spatial_rel",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "1=1",
											"kind": "query",
											"name": "where",
											"orig": "where",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
								"method": "GET",
								"orig": "/0/query",
								"parts": []any{
									"0",
									"query",
								},
								"select": map[string]any{
									"exist": []any{
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
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"metadata": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alia",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "length",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "type",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "metadata",
				"op": map[string]any{
					"list": map[string]any{
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "json",
											"kind": "query",
											"name": "f",
											"orig": "f",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
								"method": "GET",
								"orig": "/0",
								"parts": []any{
									"0",
								},
								"select": map[string]any{
									"exist": []any{
										"f",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
