package shortcuts

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/swarupdonepudi/karayaml/internal/karabinerconfig"
)

// save adds provided list of shortcuts to karabiner config
func save(shortcuts []*FileOpenShortcut) error {
	shortcutMap := convertToMap(shortcuts)

	c, err := karabinerconfig.GetDefault()
	if err != nil {
		return errors.Wrapf(err, "failed to get karabiner config")
	}

	defaultProfile, err := c.GetDefaultProfile()
	if err != nil {
		return errors.Wrapf(err, "failed to get default profile")
	}

	rules := make([]*karabinerconfig.ComplexModRule, 0)

	for _, s := range shortcutMap {
		rules = append(rules, &karabinerconfig.ComplexModRule{
			Description: fmt.Sprintf("open %s app", s.File),
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
							ShellCommand: fmt.Sprintf("open -a '%s'", s.File),
						},
					},
					Type: karabinerconfig.ComplexModRuleManipulatorTypeBasic,
				},
			},
		})
	}

	defaultProfile.ComplexMod.Rules = append(defaultProfile.ComplexMod.Rules, rules...)

	if err := c.Save(); err != nil {
		return errors.Wrapf(err, "failed to save config with shortcuts")
	}

	return nil
}

// Reload reads shortcuts from the YAML file and applies them to Karabiner.
func Reload() error {
	loaded, err := load()
	if err != nil {
		return errors.Wrapf(err, "failed to get list of shortcuts from config file")
	}
	if err := save(loaded); err != nil {
		return errors.Wrap(err, "failed to save list of shortcuts")
	}
	return nil
}
