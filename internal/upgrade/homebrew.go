package upgrade

import (
	"fmt"
	"os"
	"os/exec"
)

// UpgradeViaHomebrew upgrades the CLI using Homebrew cask
func UpgradeViaHomebrew() error {
	fmt.Println()
	fmt.Println("→ Updating Homebrew...")

	updateCmd := exec.Command("brew", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("failed to update Homebrew: %w", err)
	}

	fmt.Println()
	fmt.Println("→ Upgrading karayaml...")

	upgradeCmd := exec.Command("brew", "upgrade", "--cask", "karayaml")
	upgradeCmd.Stdout = os.Stdout
	upgradeCmd.Stderr = os.Stderr
	if err := upgradeCmd.Run(); err != nil {
		return fmt.Errorf("Homebrew upgrade failed: %w", err)
	}

	return nil
}
