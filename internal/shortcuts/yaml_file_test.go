package shortcuts

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// useTempHome points the config file at a fresh home directory for the test.
func useTempHome(t *testing.T) string {
	t.Helper()
	home := t.TempDir()
	t.Setenv("HOME", home)
	return home
}

func TestWriteRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		// in is a config file as karayaml writes it.
		in string
	}{
		{"file entries", "- file: /System/Applications/Calendar.app\n  key: \"1\"\n- file: /Applications/Slack.app\n  key: a\n- file: Bob's App\n  key: s\n"},
		{"url entry", "- key: g\n  url: nightcrawler://github\n"},
		{"mixed", "- file: Safari\n  key: s\n- key: z\n  url: https://example.com/a?b=c\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			home := useTempHome(t)
			path := filepath.Join(home, FileName)
			if err := os.WriteFile(path, []byte(tt.in), 0644); err != nil {
				t.Fatal(err)
			}
			loaded, err := load()
			if err != nil {
				t.Fatalf("load() error = %v", err)
			}
			if err := Write(loaded); err != nil {
				t.Fatalf("Write() error = %v", err)
			}
			out, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if string(out) != tt.in {
				t.Errorf("round trip changed the file:\n got: %q\nwant: %q", out, tt.in)
			}
		})
	}
}

func TestWriteURLEntryHasNoFileKey(t *testing.T) {
	home := useTempHome(t)
	if err := Write([]*FileOpenShortcut{NewShortcut("g", "nightcrawler://github")}); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	out, err := os.ReadFile(filepath.Join(home, FileName))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(out), "file:") {
		t.Errorf("url entry wrote a file key:\n%s", out)
	}
	if want := "- key: g\n  url: nightcrawler://github\n"; string(out) != want {
		t.Errorf("Write() = %q, want %q", out, want)
	}
}
