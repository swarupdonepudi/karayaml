package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
)

var Reload = &cobra.Command{
	Use:   "reload",
	Short: "reload ~/.kara.yaml into Karabiner and refresh the app",
	Run: func(cmd *cobra.Command, args []string) {
		if err := shortcuts.Reload(); err != nil {
			log.Fatalf("%v", err)
		}
		log.Info("reloaded shortcuts into Karabiner")
	},
}
