package dbhashsum

import (
	"context"
	"os"

	"github.com/sdhealth/rclone/cmd"
	"github.com/sdhealth/rclone/fs/operations"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(commandDefintion)
}

var commandDefintion = &cobra.Command{
	Use:   "dbhashsum remote:path",
	Short: `Produces a Dropbox hash file for all the objects in the path.`,
	Long: `
Produces a Dropbox hash file for all the objects in the path.  The
hashes are calculated according to [Dropbox content hash
rules](https://www.dropbox.com/developers/reference/content-hash).
The output is in the same format as md5sum and sha1sum.
`,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 1, command, args)
		fsrc := cmd.NewFsSrc(args)
		cmd.Run(false, false, command, func() error {
			return operations.DropboxHashSum(context.Background(), fsrc, os.Stdout)
		})
	},
}
