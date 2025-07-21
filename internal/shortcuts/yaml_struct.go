package shortcuts

type FileOpenShortcut struct {
	Key  KeyBoardKey `yaml:"key" json:"key"`
	File string      `yaml:"file" json:"file"`
}
