package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/karabinerconfig"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
	"os"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "initialize keyboard shortcuts",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	if err := karabinerconfig.Setup(); err != nil {
		fmt.Fprintln(os.Stderr, "failed to setup karabiner config")
		os.Exit(1)
	}

	if created, err := shortcuts.EnsureDefaultShortcuts(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to ensure default shortcuts: %v\n", err)
		os.Exit(1)
	} else if created {
		fmt.Println("created default shortcuts for Safari (s) and Mail (m)")
	}

	if err := shortcuts.Reload(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to reload shortcuts: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("initialized karabiner config and shortcuts")
}
