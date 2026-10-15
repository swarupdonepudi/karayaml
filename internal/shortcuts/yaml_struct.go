package shortcuts

// FileOpenShortcut maps a key to what it opens: an app or file (File), or a URL (URL).
// Exactly one of File and URL is set; an entry with neither is skipped on reload.
type FileOpenShortcut struct {
	File string      `yaml:"file,omitempty" json:"file,omitempty"`
	URL  string      `yaml:"url,omitempty" json:"url,omitempty"`
	Key  KeyBoardKey `yaml:"key" json:"key"`
}
