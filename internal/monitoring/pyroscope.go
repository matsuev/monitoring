package monitoring

import (
	"log"
	"runtime"

	"github.com/grafana/pyroscope-go"
)

// StartProfiling ...
func StartProfiling() {
	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	if _, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "simple.golang.app",
		// replace this with the address of pyroscope server
		ServerAddress: "http://127.0.0.1:4040",
		// you can disable logging by setting this to nil
		// Logger: pyroscope.StandardLogger,
		// Optional HTTP Basic authentication (Grafana Cloud)
		// BasicAuthUser:     "<User>",
		// BasicAuthPassword: "<Password>",
		// Optional Pyroscope tenant ID (only needed if using multi-tenancy). Not needed for Grafana Cloud.
		// TenantID:          "<TenantID>",
		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	}); err != nil {
		log.Println(err)
	}
}
