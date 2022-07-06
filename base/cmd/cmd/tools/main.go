package main

import (
	"context"
	"fmt"
	"os"

	"github.com/iotexproject/w3bstream/base/cmd/cmd/tools/gen"
	"github.com/iotexproject/w3bstream/base/cmd/cmd/tools/hook"
	"github.com/iotexproject/w3bstream/base/cmd/version"

	"github.com/go-courier/logr"
	"github.com/spf13/cobra"
)

var verbose = false

var cmd = &cobra.Command{
	Use:     "tools",
	Version: version.Version,
}

func init() {
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", verbose, "")

	cmd.AddCommand(gen.CmdGen)
	cmd.AddCommand(hook.CmdHook)
}

func main() {
	ctx := logr.WithLogger(context.Background(), logr.StdLogger())

	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
