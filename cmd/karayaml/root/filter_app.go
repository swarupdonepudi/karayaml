package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
)

var filterAppExact bool

var FilterApp = &cobra.Command{
	Use:     "by-app <query>",
	Aliases: []string{"by-file"},
	Short:   "filter shortcuts by app/file path (fuzzy substring match by default, use --exact for exact match)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		matches, err := shortcuts.FilterByApp(query, filterAppExact)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("")
		if len(matches) == 0 {
			fmt.Println("No matches found")
		} else {
			shortcuts.PrintMatches(matches)
		}
	},
}

func init() {
	FilterApp.Flags().BoolVar(&filterAppExact, "exact", false, "require exact match instead of substring match")
}
