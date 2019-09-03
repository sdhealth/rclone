package rmdir

import (
	"context"

	"github.com/sdhealth/rclone/cmd"
	"github.com/sdhealth/rclone/fs/operations"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(commandDefintion)
}

var commandDefintion = &cobra.Command{
	Use:   "rmdir remote:path",
	Short: `Remove the path if empty.`,
	Long: `
Remove the path.  Note that you can't remove a path with
objects in it, use purge for that.`,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 1, command, args)
		fdst := cmd.NewFsDir(args)
		cmd.Run(true, false, command, func() error {
			return operations.Rmdir(context.Background(), fdst, "")
		})
	},
}
