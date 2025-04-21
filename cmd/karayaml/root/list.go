package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "list shortcuts to open apps",
	Run:   listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	list, err := shortcuts.List()
	if err != nil {
		log.Fatalf("%v", err)
	}
	shortcuts.PrintList(list)
}
