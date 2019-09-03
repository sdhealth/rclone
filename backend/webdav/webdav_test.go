// Test Webdav filesystem interface
package webdav_test

import (
	"testing"

	"github.com/sdhealth/rclone/backend/webdav"
	"github.com/sdhealth/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestWebdav:",
		NilObject:  (*webdav.Object)(nil),
	})
}
