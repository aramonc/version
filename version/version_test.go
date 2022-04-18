package version_test

import (
	"context"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aramonc/version/version"
)

func TestCreatesVersionWithDefault(t *testing.T) {
	expect := version.Version{App: "development", Commit: "development", GoVersion: runtime.Version()}

	actual := version.New()

	assert.Equal(t, expect, *actual)
}

func TestCreatesVersion(t *testing.T) {
	version.AppVersion = "test1"
	version.CommitHash = "testabc"

	expect := version.Version{App: "test1", Commit: "testabc", GoVersion: runtime.Version()}

	actual := version.New()

	assert.Equal(t, expect, *actual)
}

func TestCreatesVersionInContext(t *testing.T) {
	version.AppVersion = "test1"
	version.CommitHash = "testabc"

	expect := version.Version{App: "test1", Commit: "testabc", GoVersion: runtime.Version()}

	actual := version.WithContext(context.Background())

	v, isVersion := actual.Value(version.Key).(*version.Version)

	assert.True(t, isVersion)
	assert.Equal(t, expect, *v)
}

func TestCreatesPrettyPrintString(t *testing.T) {
	v := &version.Version{
		App:       "test1",
		Commit:    "testabc",
		GoVersion: "go1.18",
	}

	expected := "App: test1\nCommit: testabc\nGo: go1.18"

	assert.Equal(t, expected, v.String())
}

func TestCreatesDataMap(t *testing.T) {
	v := &version.Version{
		App:       "test1",
		Commit:    "testabc",
		GoVersion: "go1.18",
	}

	expected := map[string]string{
		"App":    "test1",
		"Commit": "testabc",
		"Go":     "go1.18",
	}

	assert.Equal(t, expected, v.Fields())
}
