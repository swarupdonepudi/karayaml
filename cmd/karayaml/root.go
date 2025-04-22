package planton

import (
	"fmt"
	"github.com/swarupdonepudi/karayaml/cmd/karayaml/root"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:   "karayaml",
	Short: "YAML‑powered shortcut launcher for Karabiner‑Elements on macOS",
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "set log level to debug")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableSuggestions = true
	cobra.OnInitialize(func() {
		if debug {
			log.SetLevel(log.DebugLevel)
			log.Debug("running in debug mode")
		}
	})

	rootCmd.AddCommand(
		root.Add,
		root.Edit,
		root.Init,
		root.List,
		root.Version,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
