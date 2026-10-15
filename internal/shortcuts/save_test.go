package shortcuts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// karabinerCommands reads the generated Karabiner config and returns each rule's
// shell command by its description.
func karabinerCommands(t *testing.T, home string) map[string]string {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(home, ".config", "karabiner", "karabiner.json"))
	if err != nil {
		t.Fatal(err)
	}
	var cfg struct {
		Profiles []struct {
			ComplexModifications struct {
				Rules []struct {
					Description  string `json:"description"`
					Manipulators []struct {
						To []struct {
							ShellCommand string `json:"shell_command"`
						} `json:"to"`
					} `json:"manipulators"`
				} `json:"rules"`
			} `json:"complex_modifications"`
		} `json:"profiles"`
	}
	if err := json.Unmarshal(raw, &cfg); err != nil {
		t.Fatal(err)
	}
	commands := make(map[string]string)
	for _, r := range cfg.Profiles[0].ComplexModifications.Rules {
		for _, m := range r.Manipulators {
			for _, to := range m.To {
				if to.ShellCommand != "" {
					commands[r.Description] = to.ShellCommand
				}
			}
		}
	}
	return commands
}

func TestSave(t *testing.T) {
	tests := []struct {
		name      string
		shortcuts []*FileOpenShortcut
		want      map[string]string
		wantErr   string
	}{
		{
			name: "apps and urls",
			shortcuts: []*FileOpenShortcut{
				{Key: "s", File: "Safari"},
				{Key: "b", File: "Bob's App"},
				{Key: "g", URL: "nightcrawler://github"},
			},
			want: map[string]string{
				"open Safari":                `open -a 'Safari'`,
				"open Bob's App":             `open -a 'Bob'\''s App'`,
				"open nightcrawler://github": `open 'nightcrawler://github'`,
			},
		},
		{
			name:      "entry with neither is skipped",
			shortcuts: []*FileOpenShortcut{{Key: "s", File: "Safari"}, {Key: "x"}},
			want:      map[string]string{"open Safari": `open -a 'Safari'`},
		},
		{
			name:      "placeholder after a mapping keeps the mapping",
			shortcuts: []*FileOpenShortcut{{Key: "s", File: "Safari"}, {Key: "s"}},
			want:      map[string]string{"open Safari": `open -a 'Safari'`},
		},
		{
			name:      "invalid entries block the write",
			shortcuts: []*FileOpenShortcut{{Key: "s", File: "Safari"}, {Key: "g", URL: "github.com"}, {Key: "tab", File: "Mail"}},
			wantErr:   "key 'g' has url 'github.com'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			home := useTempHome(t)
			if err := os.MkdirAll(filepath.Join(home, ".config", "karabiner"), 0755); err != nil {
				t.Fatal(err)
			}
			err := save(tt.shortcuts)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) || !strings.Contains(err.Error(), "key 'tab'") {
					t.Fatalf("save() error = %v, want it to name %q and key 'tab'", err, tt.wantErr)
				}
				if _, statErr := os.Stat(filepath.Join(home, ".config", "karabiner", "karabiner.json")); !os.IsNotExist(statErr) {
					t.Errorf("save() wrote the Karabiner config despite invalid shortcuts")
				}
				return
			}
			if err != nil {
				t.Fatalf("save() error = %v", err)
			}
			got := karabinerCommands(t, home)
			if len(got) != len(tt.want) {
				t.Fatalf("save() wrote %d commands %v, want %d", len(got), got, len(tt.want))
			}
			for desc, cmd := range tt.want {
				if got[desc] != cmd {
					t.Errorf("rule %q = %s, want %s", desc, got[desc], cmd)
				}
			}
		})
	}
}
