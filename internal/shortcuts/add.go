package shortcuts

import (
	"github.com/pkg/errors"
)

// Add another shortcut to existing shortcuts, and then reloads the karabiner config.
func Add(key, file string) error {
	shortcuts, err := load()
	if err != nil {
		return errors.Wrapf(err, "failed to get list of shortcuts from config file")
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
