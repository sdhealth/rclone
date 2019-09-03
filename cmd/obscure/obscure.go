package obscure

import (
	"fmt"

	"github.com/sdhealth/rclone/cmd"
	"github.com/sdhealth/rclone/fs/config/obscure"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(commandDefintion)
}

var commandDefintion = &cobra.Command{
	Use:   "obscure password",
	Short: `Obscure password for use in the rclone.conf`,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 1, command, args)
		cmd.Run(false, false, command, func() error {
			obscured := obscure.MustObscure(args[0])
			fmt.Println(obscured)
			return nil
		})
	},
}
