package shortcuts

import (
	"github.com/pkg/errors"
)

// Add another shortcut to existing shortcuts, and then reloads the karabiner config.
// The target is an app or file, or a URL when it carries a scheme ("https://…").
func Add(key, target string) error {
	shortcut := NewShortcut(key, target)
	if errs := Validate([]*FileOpenShortcut{shortcut}); len(errs) > 0 {
		return errs[0]
	}
	// Ensure shortcuts config exists; start with empty if missing
	shortcuts := make([]*FileOpenShortcut, 0)
	if cfgPath, err := getShortcutConfigFilePath(); err == nil {
		if IsFileExists(cfgPath) {
			loaded, loadErr := load()
			if loadErr != nil {
				return errors.Wrapf(loadErr, "failed to get list of shortcuts from config file")
			}
			shortcuts = loaded
		}
	}
	shortcuts = append(shortcuts, shortcut)
	// First write to the YAML config so ~/.kara.yaml reflects the new mapping
	if err := Write(shortcuts); err != nil {
		return errors.Wrap(err, "failed to write shortcuts to yaml config")
	}
	// Then apply to Karabiner config
	if err := save(shortcuts); err != nil {
		return errors.Wrap(err, "failed to save list of shortcuts")
	}
	return nil
}
