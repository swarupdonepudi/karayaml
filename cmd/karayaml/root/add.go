package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
)

var Add = &cobra.Command{
	Use:   "add",
	Short: "add a keyboard shortcut to open an app",
	Run:   addHandler,
}

func init() {
	Add.Flags().String("key", "", "key for the shortcut")
	Add.Flags().String("file", "", "path of the file to open")
}

func addHandler(cmd *cobra.Command, args []string) {
	key, err := cmd.Flags().GetString("key")
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	if len(key) > 1 {
		log.Fatalf("key must only be one character")
		return
	}

	file, err := cmd.Flags().GetString("file")
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	if err := shortcuts.Add(key, file); err != nil {
		log.Fatalf("%v", err)
	}
}
