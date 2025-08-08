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
    if err := karabinerconfig.Setup(); err != nil {
        log.Fatalf("failed to setup karabiner config")
        return
    }
	log.Info("success!")
}
