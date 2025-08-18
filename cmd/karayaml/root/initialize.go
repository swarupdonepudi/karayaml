package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/karabinerconfig"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "initialize keyboard shortcuts",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	if err := karabinerconfig.Setup(); err != nil {
		log.Fatalf("failed to setup karabiner config")
		return
	}

	if created, err := shortcuts.EnsureDefaultShortcuts(); err != nil {
		log.Fatalf("failed to ensure default shortcuts: %v", err)
		return
	} else if created {
		log.Info("created default shortcuts for Safari (s) and Mail (m)")
	}

	if err := shortcuts.Reload(); err != nil {
		log.Fatalf("failed to reload shortcuts: %v", err)
		return
	}

	log.Info("initialized karabiner config and shortcuts")
}
