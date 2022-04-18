package version

import (
	"context"
	"fmt"
	"runtime"
)

// AppVersion variable to set the value of the application version at build time
var AppVersion string

// CommitHash variable to set the value of the commit used at build time
var CommitHash string

// Key is a constant to be used with contexts
const Key = "appVersion"

func init() {
	if AppVersion == "" {
		AppVersion = "development"
	}

	if CommitHash == "" {
		CommitHash = "development"
	}
}

// Version holds application version information
type Version struct {
	App       string
	GoVersion string
	Commit    string
}

// New returns a pointer to the version object
func New() *Version {
	return createVersion()
}

// WithContext returns a new context with the version object as a variable
func WithContext(ctx context.Context) context.Context {
	v := createVersion()

	return context.WithValue(ctx, Key, v)
}

func createVersion() *Version {
	return &Version{App: AppVersion, Commit: CommitHash, GoVersion: runtime.Version()}
}

// String generates a pretty print string with the version data
func (v *Version) String() string {
	return fmt.Sprintf("App: %s\nCommit: %s\nGo: %s", v.App, v.Commit, v.GoVersion)
}

// Fields returns a map of version data
func (v *Version) Fields() map[string]string {
	return map[string]string{
		"App":    v.App,
		"Commit": v.Commit,
		"Go":     v.GoVersion,
	}
}
