package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
)

var Edit = &cobra.Command{
	Use:   "edit",
	Short: "edit keyboard shortcuts to open apps",
	Run:   editHandler,
}

func editHandler(cmd *cobra.Command, args []string) {
	if err := shortcuts.Edit(); err != nil {
		log.Fatalf("%v", err)
	}
}
