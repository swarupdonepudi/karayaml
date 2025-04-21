package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionLabel = "dev"

var Version = &cobra.Command{
	Use:   "version",
	Short: "check the version of the cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s", VersionLabel))
	},
}
