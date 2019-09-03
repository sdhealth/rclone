// +build linux darwin freebsd

package mount

import (
	"testing"

	"github.com/sdhealth/rclone/cmd/mountlib/mounttest"
)

func TestMount(t *testing.T) {
	mounttest.RunTests(t, mount)
}
