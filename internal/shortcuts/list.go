package shortcuts

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
	"os"
)

// List returns the list of keyboard shortcuts as they exist in the YAML file.
// The order returned reflects the file order (prior to any display sorting).
func List() ([]*FileOpenShortcut, error) {
	shortcuts, err := load()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get keyboard shortcuts")
	}
	return shortcuts, nil
}

// PrintList prints the provided shortcuts to the console as three tables:
// Function Keys, Numbers, and Letters. Within each category, entries are
// ordered stably so duplicate keys preserve their YAML order.
func PrintList(shortcuts []*FileOpenShortcut) {
	ordered := make([]*FileOpenShortcut, len(shortcuts))
	copy(ordered, shortcuts)
	sortShortcutsStable(ordered)

	functionKeys := make([]*FileOpenShortcut, 0)
	numbers := make([]*FileOpenShortcut, 0)
	letters := make([]*FileOpenShortcut, 0)
	for _, s := range ordered {
		c, _, _ := classifyKey(string(s.Key))
		switch c {
		case 0:
			functionKeys = append(functionKeys, s)
		case 1:
			numbers = append(numbers, s)
		default:
			letters = append(letters, s)
		}
	}

	if len(functionKeys) > 0 {
		rows := make([]table.Row, 0, len(functionKeys))
		for _, s := range functionKeys {
			rows = append(rows, table.Row{s.Key, s.File})
		}
		printTableWithTitle("Function Keys", rows)
	}
	if len(numbers) > 0 {
		rows := make([]table.Row, 0, len(numbers))
		for _, s := range numbers {
			rows = append(rows, table.Row{s.Key, s.File})
		}
		printTableWithTitle("Numbers", rows)
	}
	if len(letters) > 0 {
		rows := make([]table.Row, 0, len(letters))
		for _, s := range letters {
			rows = append(rows, table.Row{s.Key, s.File})
		}
		printTableWithTitle("Letters", rows)
	}
}

func printTable(header table.Row, rows []table.Row) {
	println("")
	t := getDefaultTableWriter(header, rows)
	t.Render()
	println("")
}

func printTableWithTitle(title string, rows []table.Row) {
	println("")
	t := getDefaultTableWriter(nil, rows)
	if title != "" {
		t.SetTitle(title)
	}
	t.Render()
	println("")
}

func getDefaultTableWriter(header table.Row, rows []table.Row) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if header != nil {
		t.AppendHeader(header)
	}
	for _, r := range rows {
		t.AppendRow(r)
		t.AppendSeparator()
	}
	return t
}
