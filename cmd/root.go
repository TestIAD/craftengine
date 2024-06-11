package cmd

import (
	"fmt"
	"os"

	"github.com/TestIAD/craftengine/internal"
	"github.com/spf13/cobra"
)

var (
	module  string
	service string
	path    string
)

func init() {
	rootCMD.Flags().StringVarP(
		&module, "module", "m", "console",
		"module for console or admin",
	)
	rootCMD.Flags().StringVarP(
		&service, "service", "s", "",
		"what service with a upper camel case",
	)
	rootCMD.Flags().StringVarP(
		&path, "path", "p", "",
		"app path",
	)
}

var (
	rootCMD = &cobra.Command{
		Use:   "craft",
		Short: "craft is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.`,
		Run: func(cmd *cobra.Command, args []string) {
			internal.Parse(module, service, path)
		},
	}
)

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
