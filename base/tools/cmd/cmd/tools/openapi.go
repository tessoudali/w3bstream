package main

import (
	"github.com/go-courier/httptransport/generators/openapi"
	_ "github.com/go-courier/httptransport/validator/strfmt"
	"github.com/go-courier/packagesx"
	"github.com/iotexproject/w3bstream/base/tools/cmd/internal/generate"
	"github.com/spf13/cobra"
)

var cmdSwagger = &cobra.Command{
	Use:     "openapi",
	Aliases: []string{"swagger"},
	Short:   "scan current project and generate openapi.json",
	Run: func(cmd *cobra.Command, args []string) {
		generate.Run("openapi", func(pkg *packagesx.Package) generate.Generator {
			g := openapi.NewOpenAPIGenerator(pkg)
			g.Scan(cmd.Context())
			return g
		})
	},
}

func init() {
	cmd.AddCommand(cmdSwagger)
}
