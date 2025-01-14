// Syslog interface for non-Unix variants only

// +build windows nacl plan9

package log

import (
	"log"
	"runtime"
)

// Starts syslog if configured, returns true if it was started
func startSysLog() bool {
	log.Panicf("--syslog not supported on %s platform", runtime.GOOS)
	return false
}
