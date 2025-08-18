package shortcuts

type FileOpenShortcut struct {
	File string      `yaml:"file" json:"file"`
	Key  KeyBoardKey `yaml:"key" json:"key"`
}
