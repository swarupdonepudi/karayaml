package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/karabinerconfig"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "initialize keyboard shortcuts",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	c, err := karabinerconfig.GetDefault()
	if err != nil {
		log.Fatalf("failed to get karabiner config")
		return
	}
	if err := c.Save(); err != nil {
		log.Fatalf("failed to save config with shortcuts")
		return
	}
	log.Info("success!")
}
