package cmd

import (
	"fmt"
	"os"

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
		"module",
	)
	rootCMD.Flags().StringVarP(
		&service, "service", "s", "",
		"service",
	)
	rootCMD.Flags().StringVarP(
		&path, "path", "p", "",
		"path",
	)
}

var (
	rootCMD = &cobra.Command{
		Use:   "use",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run hugo...")
			fmt.Printf("%s\n", module)
			fmt.Printf("%s\n", service)
		},
	}
)

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
