// Test Mega filesystem interface
package mega_test

import (
	"testing"

	"github.com/sdhealth/rclone/backend/mega"
	"github.com/sdhealth/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestMega:",
		NilObject:  (*mega.Object)(nil),
	})
}
