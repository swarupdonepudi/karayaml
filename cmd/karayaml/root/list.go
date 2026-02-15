package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
	"os"
)

var List = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list shortcuts to open apps",
	Run:   listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	list, err := shortcuts.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	shortcuts.PrintList(list)
}
