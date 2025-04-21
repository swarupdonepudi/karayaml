package shortcuts

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/swarupdonepudi/karayaml/internal/karabinerconfig"
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
	shortcutMap := convertToMap(shortcuts)
	c, err := karabinerconfig.GetDefault()
	if err != nil {
		return errors.Wrapf(err, "failed to get karabiner config")
	}
	if err := addShortcuts(c, shortcutMap); err != nil {
		return errors.Wrapf(err, "failed to add shortcuts to open apps")
	}
	if err := c.Save(); err != nil {
		return errors.Wrapf(err, "failed to save config with shortcuts")
	}
	return nil
}

// addShortcuts adds the provided map of shortcuts to karabiner config
func addShortcuts(c *karabinerconfig.Config, shortcuts map[KeyBoardKey]*AppShortcut) error {
	defaultProfile, err := c.GetDefaultProfile()
	if err != nil {
		return errors.Wrapf(err, "failed to get default profile")
	}
	rules := make([]*karabinerconfig.ComplexModRule, 0)
	for _, s := range shortcuts {
		rules = append(rules, &karabinerconfig.ComplexModRule{
			Description: fmt.Sprintf("open %s app", s.AppFilePath),
			Manipulators: []*karabinerconfig.ComplexModRuleManipulator{
				{
					From: &karabinerconfig.ComplexModRuleManipulatorFrom{
						KeyCode: string(s.Key),
						Modifiers: &karabinerconfig.ComplexModRuleManipulatorFromModifiers{
							Mandatory: karabinerconfig.CapsLockModifierKeys,
						},
					},
					To: []*karabinerconfig.ComplexModRuleManipulatorTo{
						{
							ShellCommand: fmt.Sprintf("open -a '%s'", s.AppFilePath),
						},
					},
					Type: karabinerconfig.ComplexModRuleManipulatorTypeBasic,
				},
			},
		})
	}
	defaultProfile.ComplexMod.Rules = append(defaultProfile.ComplexMod.Rules, rules...)
	return nil
}
