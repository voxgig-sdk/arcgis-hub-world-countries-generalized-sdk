package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go"
	"github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go/core"

	vs "github.com/voxgig-sdk/arcgis-hub-world-countries-generalized-sdk/go/utility/struct"
)

func TestFeatureEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Feature(nil)
		if ent == nil {
			t.Fatal("expected non-nil FeatureEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := featureBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "feature." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set ARCGISHUBWORLDCOUNTRIESGENERALIZED_TEST_FEATURE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		featureRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.feature", setup.data)))
		var featureRef01Data map[string]any
		if len(featureRef01DataRaw) > 0 {
			featureRef01Data = core.ToMapAny(featureRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = featureRef01Data

		// LIST
		featureRef01Ent := client.Feature(nil)
		featureRef01Match := map[string]any{}

		featureRef01ListResult, err := featureRef01Ent.List(featureRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, featureRef01ListOk := featureRef01ListResult.([]any)
		if !featureRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", featureRef01ListResult)
		}

	})
}

func featureBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "feature", "FeatureTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read feature test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse feature test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"feature01", "feature02", "feature03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("ARCGISHUBWORLDCOUNTRIESGENERALIZED_TEST_FEATURE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"ARCGISHUBWORLDCOUNTRIESGENERALIZED_TEST_FEATURE_ENTID": idmap,
		"ARCGISHUBWORLDCOUNTRIESGENERALIZED_TEST_LIVE":      "FALSE",
		"ARCGISHUBWORLDCOUNTRIESGENERALIZED_TEST_EXPLAIN":   "FALSE",
		"ARCGISHUBWORLDCOUNTRIESGENERALIZED_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["ARCGISHUBWORLDCOUNTRIESGENERALIZED_TEST_FEATURE_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["ARCGISHUBWORLDCOUNTRIESGENERALIZED_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["ARCGISHUBWORLDCOUNTRIESGENERALIZED_APIKEY"],
			},
			extra,
		})
		client = sdk.NewArcgisHubWorldCountriesGeneralizedSDK(core.ToMapAny(mergedOpts))
	}

	live := env["ARCGISHUBWORLDCOUNTRIESGENERALIZED_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["ARCGISHUBWORLDCOUNTRIESGENERALIZED_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
