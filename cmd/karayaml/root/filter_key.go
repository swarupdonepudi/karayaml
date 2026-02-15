package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
)

var filterKeyExact bool

var FilterKey = &cobra.Command{
	Use:   "by-key <query>",
	Short: "filter shortcuts by key (fuzzy substring match by default, use --exact for exact match)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		matches, err := shortcuts.FilterByKey(query, filterKeyExact)
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
	FilterKey.Flags().BoolVar(&filterKeyExact, "exact", false, "require exact match instead of substring match")
}
