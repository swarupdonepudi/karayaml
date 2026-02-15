package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/karayaml/internal/shortcuts"
	"os"
)

var Find = &cobra.Command{
	Use:     "find <query>",
	Aliases: []string{"search"},
	Short:   "search ~/.kara.yaml for apps whose name contains the query (case-insensitive)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		matches, err := shortcuts.Find(query)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// readability: blank line before printing result summary/table
		fmt.Println("")
		switch len(matches) {
		case 0:
			fmt.Println("No matches found")
		case 1:
			fmt.Println(shortcuts.FormatSingleMatchMessage(matches[0]))
			shortcuts.PrintMatches(matches)
		default:
			fmt.Println(shortcuts.FormatMultiMatchMessage())
			shortcuts.PrintMatches(matches)
		}
	},
}
