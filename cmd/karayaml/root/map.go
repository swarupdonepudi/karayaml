package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
)

var Map = &cobra.Command{
	Use:   "map <key> <file>",
	Short: "map a key to open an app",
	Args:  cobra.ExactArgs(2),
	Run:   mapHandler,
}

func mapHandler(cmd *cobra.Command, args []string) {
	key := args[0]
	file := args[1]

	if len(key) != 1 {
		log.Fatalf("key must only be one character")
		return
	}

	if err := shortcuts.Add(key, file); err != nil {
		log.Fatalf("%v", err)
	}
}


