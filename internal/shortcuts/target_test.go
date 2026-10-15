package shortcuts

import (
	"strings"
	"testing"
)

func TestShellQuote(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"plain", "Safari", `'Safari'`},
		{"spaces", "/Applications/Google Chrome.app", `'/Applications/Google Chrome.app'`},
		{"single quote", "Bob's App", `'Bob'\''s App'`},
		{"only quotes", "''", `''\'''\'''`},
		{"empty", "", `''`},
		{"url with query", "https://example.com/?a=1&b=2", `'https://example.com/?a=1&b=2'`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shellQuote(tt.in); got != tt.want {
				t.Errorf("shellQuote(%q) = %s, want %s", tt.in, got, tt.want)
			}
		})
	}
}

func TestShellCommandAndTarget(t *testing.T) {
	tests := []struct {
		name       string
		shortcut   FileOpenShortcut
		wantCmd    string
		wantTarget string
	}{
		{"app", FileOpenShortcut{Key: "s", File: "Safari"}, `open -a 'Safari'`, "Safari"},
		{"app path", FileOpenShortcut{Key: "a", File: "/Applications/Slack.app"}, `open -a '/Applications/Slack.app'`, "/Applications/Slack.app"},
		{"app with quote", FileOpenShortcut{Key: "b", File: "Bob's App"}, `open -a 'Bob'\''s App'`, "Bob's App"},
		{"https url", FileOpenShortcut{Key: "g", URL: "https://github.com"}, `open 'https://github.com'`, "https://github.com"},
		{"custom scheme", FileOpenShortcut{Key: "g", URL: "nightcrawler://github"}, `open 'nightcrawler://github'`, "nightcrawler://github"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.shortcut.ShellCommand(); got != tt.wantCmd {
				t.Errorf("ShellCommand() = %s, want %s", got, tt.wantCmd)
			}
			if got := tt.shortcut.Target(); got != tt.wantTarget {
				t.Errorf("Target() = %s, want %s", got, tt.wantTarget)
			}
		})
	}
}

func TestNewShortcut(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		wantFile string
		wantURL  string
	}{
		{"app name", "Safari", "Safari", ""},
		{"app path", "/Applications/Slack.app", "/Applications/Slack.app", ""},
		{"https url", "https://github.com/pulls", "", "https://github.com/pulls"},
		{"custom scheme", "nightcrawler://github", "", "nightcrawler://github"},
		{"colon without slashes is a file", "Notes: work", "Notes: work", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShortcut("g", tt.target)
			if s.File != tt.wantFile || s.URL != tt.wantURL {
				t.Errorf("NewShortcut(%q) = file %q url %q, want file %q url %q", tt.target, s.File, s.URL, tt.wantFile, tt.wantURL)
			}
			if s.Key != "g" {
				t.Errorf("NewShortcut key = %q, want g", s.Key)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		shortcuts []*FileOpenShortcut
		wantErrs  []string
	}{
		{"app", []*FileOpenShortcut{{Key: "s", File: "Safari"}}, nil},
		{"url", []*FileOpenShortcut{{Key: "g", URL: "nightcrawler://github"}}, nil},
		{"mailto url", []*FileOpenShortcut{{Key: "m", URL: "mailto:someone@example.com"}}, nil},
		{"uppercase key", []*FileOpenShortcut{{Key: "G", File: "Safari"}}, nil},
		{"neither set is skipped, not invalid", []*FileOpenShortcut{{Key: "x"}}, nil},
		{"both set", []*FileOpenShortcut{{Key: "g", File: "Safari", URL: "https://github.com"}}, []string{"key 'g' sets both file and url"}},
		{"url without scheme", []*FileOpenShortcut{{Key: "g", URL: "github.com"}}, []string{"key 'g' has url 'github.com', which needs a scheme"}},
		{"unsupported key", []*FileOpenShortcut{{Key: "tab", File: "Safari"}}, []string{"key 'tab' is not a supported key"}},
		{
			"every problem named",
			[]*FileOpenShortcut{
				{Key: "s", File: "Safari"},
				{Key: "g", URL: "github.com"},
				{Key: "tab", File: "Mail", URL: "https://mail.example.com"},
			},
			[]string{"key 'g' has url", "key 'tab' is not a supported key", "key 'tab' sets both file and url"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := Validate(tt.shortcuts)
			if len(errs) != len(tt.wantErrs) {
				t.Fatalf("Validate() returned %d errors %v, want %d", len(errs), errs, len(tt.wantErrs))
			}
			for i, want := range tt.wantErrs {
				if !strings.Contains(errs[i].Error(), want) {
					t.Errorf("error %d = %q, want it to contain %q", i, errs[i], want)
				}
			}
		})
	}
}
