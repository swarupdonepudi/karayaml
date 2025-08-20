package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
	"os"
)

var Reload = &cobra.Command{
	Use:   "reload",
	Short: "reload ~/.kara.yaml into Karabiner and refresh the app",
	Run: func(cmd *cobra.Command, args []string) {
		if err := shortcuts.Reload(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("reloaded shortcuts into Karabiner")
	},
}
