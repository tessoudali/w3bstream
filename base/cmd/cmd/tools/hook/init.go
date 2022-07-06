package hook

import (
	"github.com/iotexproject/w3bstream/base/cmd/internal/githooks"
	"github.com/spf13/cobra"
)

func init() {
	CmdHook.AddCommand(cmdHookInit)
}

var cmdHookInit = &cobra.Command{
	Use:   "init",
	Short: "git hook init",
	Run: func(cmd *cobra.Command, args []string) {
		githooks.Init()
	},
}
