package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
	"os"
	"strings"
)

var Map = &cobra.Command{
	Use:   "map <key> <file>",
	Short: "map a key to open an app",
	Args:  cobra.ExactArgs(2),
	Run:   mapHandler,
}

func mapHandler(cmd *cobra.Command, args []string) {
	key := strings.ToLower(args[0])
	file := args[1]

	if !shortcuts.IsValidKeyBoardKey(key) {
		fmt.Fprintln(os.Stderr, "invalid key: must be one of the supported keys; see internal/shortcuts/key_board_key.go")
		os.Exit(1)
	}

	if err := shortcuts.Add(key, file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
