package root

import (
	"github.com/spf13/cobra"
)

var Filter = &cobra.Command{
	Use:   "filter",
	Short: "filter shortcuts by key or app/file",
}

func init() {
	Filter.AddCommand(FilterKey, FilterApp)
}
