package shortcuts

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
	"os"
	"strings"
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

// PrintMatches prints a flat table of the provided shortcuts (key and file),
// preserving stable ordering for readability.
func PrintMatches(matches []*FileOpenShortcut) {
	if len(matches) == 0 {
		return
	}
	ordered := make([]*FileOpenShortcut, len(matches))
	copy(ordered, matches)
	sortShortcutsStable(ordered)

	rows := make([]table.Row, 0, len(ordered))
	for _, s := range ordered {
		rows = append(rows, table.Row{s.Key, s.File})
	}
	printTableWithTitle("Matches", rows)
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

// Find returns all shortcuts where the file name contains the query
// (case-insensitive). It does not mutate ordering.
func Find(query string) ([]*FileOpenShortcut, error) {
	list, err := List()
	if err != nil {
		return nil, err
	}
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return []*FileOpenShortcut{}, nil
	}
	matches := make([]*FileOpenShortcut, 0)
	for _, s := range list {
		if strings.Contains(strings.ToLower(s.File), q) {
			matches = append(matches, s)
		}
	}
	return matches, nil
}

// FilterByKey returns shortcuts whose Key matches the query.
// If exact is true, only exact (case-insensitive) matches are returned.
// Otherwise, keys that contain the query as a substring are returned.
func FilterByKey(query string, exact bool) ([]*FileOpenShortcut, error) {
	list, err := List()
	if err != nil {
		return nil, err
	}
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return []*FileOpenShortcut{}, nil
	}
	matches := make([]*FileOpenShortcut, 0)
	for _, s := range list {
		k := strings.ToLower(string(s.Key))
		if exact {
			if k == q {
				matches = append(matches, s)
			}
		} else {
			if strings.Contains(k, q) {
				matches = append(matches, s)
			}
		}
	}
	return matches, nil
}

// FilterByApp returns shortcuts whose File path matches the query.
// If exact is true, only exact (case-insensitive) matches are returned.
// Otherwise, files that contain the query as a substring are returned.
func FilterByApp(query string, exact bool) ([]*FileOpenShortcut, error) {
	list, err := List()
	if err != nil {
		return nil, err
	}
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return []*FileOpenShortcut{}, nil
	}
	matches := make([]*FileOpenShortcut, 0)
	for _, s := range list {
		f := strings.ToLower(s.File)
		if exact {
			if f == q {
				matches = append(matches, s)
			}
		} else {
			if strings.Contains(f, q) {
				matches = append(matches, s)
			}
		}
	}
	return matches, nil
}

// FormatSingleMatchMessage returns a friendly sentence for a single match.
func FormatSingleMatchMessage(s *FileOpenShortcut) string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf("Yeah, I found one match: key '%s' is mapped to '%s'", s.Key, s.File)
}

// FormatMultiMatchMessage returns a header sentence for multiple matches.
func FormatMultiMatchMessage() string {
	return "We have found multiple matches and here is the list"
}
