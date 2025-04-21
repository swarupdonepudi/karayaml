package shortcuts

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
	"os"
	"sort"
)

// List returns a map of keyboard shortcuts
func List() (map[KeyBoardKey]*AppShortcut, error) {
	shortcuts, err := load()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get keyboard shortcuts")
	}
	return convertToMap(shortcuts), nil
}

// PrintList prints the provided shortcuts map to the console
func PrintList(shortcuts map[KeyBoardKey]*AppShortcut) {
	header := table.Row{"key", "app path"}
	rows := make([]table.Row, 0)
	hotkeys := make([]string, 0)
	for k, _ := range shortcuts {
		hotkeys = append(hotkeys, string(k))
	}
	sort.Strings(hotkeys)
	for _, r := range hotkeys {
		rows = append(rows, table.Row{shortcuts[KeyBoardKey(r)].Key, shortcuts[KeyBoardKey(r)].AppFilePath})
	}
	printTable(header, rows)
}

func printTable(header table.Row, rows []table.Row) {
	println("")
	t := getDefaultTableWriter(header, rows)
	t.Render()
	println("")
}

// convertToMap converts the provided shortcuts into a map
func convertToMap(shortcuts []*AppShortcut) map[KeyBoardKey]*AppShortcut {
	shortcutsMap := make(map[KeyBoardKey]*AppShortcut, 0)
	for _, s := range shortcuts {
		shortcutsMap[s.Key] = s
	}
	return shortcutsMap
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
