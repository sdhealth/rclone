package reveal

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
	Use:   "reveal password",
	Short: `Reveal obscured password from rclone.conf`,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 1, command, args)
		cmd.Run(false, false, command, func() error {
			revealed, err := obscure.Reveal(args[0])
			if err != nil {
				return err
			}
			fmt.Println(revealed)
			return nil
		})
	},
	Hidden: true,
}
