/*
Copyright © 2023 Thomas Stringer <thomas@trstringer.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var list bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-example",
	Short: "Show example manifests for resources",
	Run: func(cmd *cobra.Command, args []string) {
		if list {
			resourceNames, err := getResources()
			if err != nil {
				fmt.Printf("Error getting resource names: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(strings.Join(resourceNames, "\n"))
			os.Exit(0)
		}

		if len(args) == 0 {
			fmt.Println("Pass in a resource to show or `--list`")
			os.Exit(1)
		}

		definitions, err := getResourcesByName(args)
		if err != nil {
			fmt.Printf("Error getting definitions: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf(strings.Join(definitions, "---\n"))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&list, "list", "l", false, "list available resources")
}
