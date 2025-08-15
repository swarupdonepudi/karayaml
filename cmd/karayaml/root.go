package planton

import (
	"fmt"
	"os"
	"strings"

	"github.com/swarupdonepudi/karayaml/cmd/karayaml/root"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:     "karayaml",
	Version: root.VersionLabel,
	Short:   "YAML‑powered shortcut launcher for Karabiner‑Elements on macOS",
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "set log level to debug")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableSuggestions = true

	// Ensure version flags behave consistently across variants
	rootCmd.SetVersionTemplate("{{.Version}}\n")
	rootCmd.InitDefaultVersionFlag()
	if vf := rootCmd.Flags().Lookup("version"); vf != nil {
		vf.Shorthand = "v"
	}
	cobra.OnInitialize(func() {
		if debug {
			log.SetLevel(log.DebugLevel)
			log.Debug("running in debug mode")
		}
	})

	rootCmd.AddCommand(
		root.Map,
		root.Edit,
		root.Init,
		root.List,
		root.Version,
	)
}

func Execute() {
	// Normalize legacy single-dash long flag: "-version" -> "--version"
	os.Args = normalizeLegacyVersionArgs(os.Args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// normalizeLegacyVersionArgs converts "-version" and "-version=<value>" to their
// GNU-style equivalents so users can run `karayaml -version` without errors.
func normalizeLegacyVersionArgs(args []string) []string {
	for i, a := range args {
		if a == "-version" {
			args[i] = "--version"
		} else if strings.HasPrefix(a, "-version=") {
			// Preserve any explicit assignment, e.g. -version=true
			args[i] = "--" + a[1:]
		}
	}
	return args
}
