package shortcuts

import (
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/swarupdonepudi/karayaml/internal/defaulteditor"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	FileName = "shortcuts.yaml"
)

// editYaml opens config file in vs-code and waits until the file is closed.
func editYaml() error {
	configFile, err := getShortcutConfigFilePath()
	if err != nil {
		return errors.Wrapf(err, "failed to get shortcut config file path")
	}
    // Ensure file exists with an empty list if missing
    if !IsFileExists(configFile) {
        if err := Write([]*FileOpenShortcut{}); err != nil {
            return errors.Wrapf(err, "failed to initialize shortcut config file")
        }
    }
	for {
		duplicates := make([]string, 0)

		cmd := exec.Command(defaulteditor.DefaultEditor, "--wait", configFile)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()

		appShortcuts, err := load()
		if err != nil {
			return errors.Wrapf(err, "failed to list shortcuts")
		}
		duplicates = getDuplicates(appShortcuts)
		if len(duplicates) == 0 {
			break
		} else {
			log.Errorf("fix %v keys which may have either empty or duplicate shortcut mappings", duplicates)
		}
	}
	return nil
}

// getDuplicates iterates through the shortcuts and considers shortcuts that have empty app name to be a duplicate entry.
// note: the value could also be left empty intentionally.
func getDuplicates(shortcuts []*FileOpenShortcut) []string {
	duplicates := make([]string, 0)
	for _, s := range shortcuts {
		if s.File != "" {
			continue
		}
		duplicates = append(duplicates, string(s.Key))
	}
	return duplicates
}

// Write writes provided shortcuts to shortcuts config file
func Write(shortcuts []*FileOpenShortcut) error {
	shortcutConfigFilePath, err := getShortcutConfigFilePath()
	if err != nil {
		return errors.Wrapf(err, "failed to get config file path")
	}
	if !IsDirExists(filepath.Dir(shortcutConfigFilePath)) {
        if err := os.MkdirAll(filepath.Dir(shortcutConfigFilePath), 0755); err != nil {
			return errors.Wrapf(err, "failed to create %s dir", filepath.Dir(shortcutConfigFilePath))
		}
	}
	defaultShortcutsYamlBytes, err := yaml.Marshal(shortcuts)
	if err != nil {
		return errors.Wrapf(err, "failed to initialize")
	}
    if err := os.WriteFile(shortcutConfigFilePath, defaultShortcutsYamlBytes, 0644); err != nil {
		return errors.Wrapf(err, "failed to write %s file", shortcutConfigFilePath)
	}
	return nil
}

// load returns list of app shortcuts from the shortcuts config file
func load() ([]*FileOpenShortcut, error) {
	appShortcuts := make([]*FileOpenShortcut, 0)
	shortcutConfigFile, err := getShortcutConfigFilePath()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get extra file path")
	}
	shortcutsFileContent, err := os.ReadFile(shortcutConfigFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s file", shortcutConfigFile)
	}
	if err := yaml.Unmarshal(shortcutsFileContent, &appShortcuts); err != nil {
		return nil, errors.Wrapf(err, "failed to yaml unmarshal app shortcuts config file")
	}
	return appShortcuts, nil
}

// getShortcutConfigFilePath returns the location of the keyboard shortcut configs
func getShortcutConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get home dir")
	}
	return filepath.Join(homeDir, ".karayaml", FileName), nil
}

// IsFileExists check if a file exists
func IsFileExists(f string) bool {
	if f == "" {
		return false
	}
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// IsDirExists check if a directory exists
func IsDirExists(d string) bool {
	if d == "" {
		return false
	}
	info, err := os.Stat(d)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
