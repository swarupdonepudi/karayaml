package shortcuts

import (
	"github.com/pkg/errors"
)

// Add another shortcut to existing shortcuts, and then reloads the karabiner config.
func Add(key, file string) error {
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
	shortcuts = append(shortcuts, &FileOpenShortcut{
		Key:  KeyBoardKey(key),
		File: file,
	})
	if err := save(shortcuts); err != nil {
		return errors.Wrap(err, "failed to save list of shortcuts")
	}
	return nil
}
