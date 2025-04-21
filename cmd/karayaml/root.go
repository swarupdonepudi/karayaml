package planton

import (
	"fmt"
	"github.com/swarupdonepudi/karayaml/cmd/karayaml/root"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debug bool
var traceApiRequests bool

var rootCmd = &cobra.Command{
	Use:   "karayaml",
	Short: "yaml interface for karabiner-elements",
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, string("debug"), false, "set log level to debug")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableSuggestions = true
	cobra.OnInitialize(func() {
		if debug {
			log.SetLevel(log.DebugLevel)
			log.Debug("running in debug mode")
		}
	})

	rootCmd.AddCommand(root.Version)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
