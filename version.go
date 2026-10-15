package aiven

import (
	"runtime/debug"
	"sync"
)

const modulePath = "github.com/aiven/go-client-codegen"

// Version returns the module version from build info, or "v0.0.0" for local builds.
func Version() string { return version() }

var version = sync.OnceValue(func() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "v0.0.0"
	}
	for _, m := range info.Deps {
		if m.Path == modulePath {
			return m.Version
		}
	}
	return "v0.0.0"
})
