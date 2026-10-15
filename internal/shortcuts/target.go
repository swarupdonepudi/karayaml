package shortcuts

import (
	"fmt"
	"net/url"
	"strings"
)

// NewShortcut returns a shortcut for the key that opens the target: a URL when the
// target carries a scheme separator ("https://…", "slack://…"), an app or file otherwise.
func NewShortcut(key, target string) *FileOpenShortcut {
	s := &FileOpenShortcut{Key: KeyBoardKey(key)}
	if strings.Contains(target, "://") {
		s.URL = target
	} else {
		s.File = target
	}
	return s
}

// Target returns what the shortcut opens: the URL when set, otherwise the app or file.
func (s *FileOpenShortcut) Target() string {
	if s.URL != "" {
		return s.URL
	}
	return s.File
}

// ShellCommand returns the command Karabiner runs for the shortcut: `open -a` for an
// app or file, plain `open` for a URL, which macOS hands to the URL's registered app.
func (s *FileOpenShortcut) ShellCommand() string {
	if s.URL != "" {
		return "open " + shellQuote(s.URL)
	}
	return "open -a " + shellQuote(s.File)
}

// shellQuote wraps the value in single quotes for /bin/sh, escaping any single quote
// inside it, so a name like "Bob's App" stays one argument.
func shellQuote(v string) string {
	return "'" + strings.ReplaceAll(v, "'", `'\''`) + "'"
}

// Validate returns one error per invalid shortcut: an unsupported key, both file and
// url set, or a url without a scheme. An entry with neither file nor url is not an
// error; it is skipped when the Karabiner config is written.
func Validate(shortcuts []*FileOpenShortcut) []error {
	errs := make([]error, 0)
	for _, s := range shortcuts {
		if !IsValidKeyBoardKey(string(s.Key)) {
			errs = append(errs, fmt.Errorf("key '%s' is not a supported key", s.Key))
		}
		if s.File != "" && s.URL != "" {
			errs = append(errs, fmt.Errorf("key '%s' sets both file and url; keep one", s.Key))
		}
		if s.URL != "" {
			if u, err := url.Parse(s.URL); err != nil || u.Scheme == "" {
				errs = append(errs, fmt.Errorf("key '%s' has url '%s', which needs a scheme such as https://", s.Key, s.URL))
			}
		}
	}
	return errs
}

// formatErrors joins validation errors into one indented list for display.
func formatErrors(errs []error) string {
	lines := make([]string, 0, len(errs))
	for _, e := range errs {
		lines = append(lines, "  "+e.Error())
	}
	return strings.Join(lines, "\n")
}
