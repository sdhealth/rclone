// Test AzureBlob filesystem interface

// +build !plan9,!solaris

package azureblob

import (
	"testing"

	"github.com/sdhealth/rclone/fs"
	"github.com/sdhealth/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName:  "TestAzureBlob:",
		NilObject:   (*Object)(nil),
		TiersToTest: []string{"Hot", "Cool"},
		ChunkedUpload: fstests.ChunkedUploadConfig{
			MaxChunkSize: maxChunkSize,
		},
	})
}

func (f *Fs) SetUploadChunkSize(cs fs.SizeSuffix) (fs.SizeSuffix, error) {
	return f.setUploadChunkSize(cs)
}

func (f *Fs) SetUploadCutoff(cs fs.SizeSuffix) (fs.SizeSuffix, error) {
	return f.setUploadCutoff(cs)
}

var (
	_ fstests.SetUploadChunkSizer = (*Fs)(nil)
	_ fstests.SetUploadCutoffer   = (*Fs)(nil)
)
