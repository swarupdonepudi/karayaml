package shortcuts

import (
	"github.com/pkg/errors"
)

// Edit opens config file in vs-code and waits until the file is closed and then reloads the karabiner config.
func Edit() error {
	if err := editYaml(); err != nil {
		return errors.Wrapf(err, "failed to edit the config file")
	}
	shortcuts, err := load()
	if err != nil {
		return errors.Wrapf(err, "failed to get list of shortcuts from config file")
	}
	// Persist sorted YAML while preserving duplicate key order
	if err := Write(shortcuts); err != nil {
		return errors.Wrapf(err, "failed to write sorted shortcuts to yaml config")
	}
	if err := save(shortcuts); err != nil {
		return errors.Wrap(err, "failed to save list of shortcuts")
	}
	return nil
}
