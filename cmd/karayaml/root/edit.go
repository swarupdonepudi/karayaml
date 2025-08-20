package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
	"os"
)

var Edit = &cobra.Command{
	Use:   "edit",
	Short: "edit keyboard shortcuts to open apps",
	Run:   editHandler,
}

func editHandler(cmd *cobra.Command, args []string) {
	if err := shortcuts.Edit(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
