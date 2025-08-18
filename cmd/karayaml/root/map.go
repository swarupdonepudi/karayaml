package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
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
		log.Fatalf("invalid key: must be one of the supported keys; see internal/shortcuts/key_board_key.go")
		return
	}

	if err := shortcuts.Add(key, file); err != nil {
		log.Fatalf("%v", err)
	}
}
