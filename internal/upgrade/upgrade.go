package upgrade

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

// Run executes the upgrade command
func Run(currentVersion string, checkOnly bool, force bool) {
	if currentVersion == "" {
		currentVersion = "dev"
	}

	fmt.Println("→ Checking for updates...")

	latestVersion, err := GetLatestVersion()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n✗ Failed to check for updates: %v\n", err)
		fmt.Fprintln(os.Stderr, "  You can manually download from https://github.com/swarupdonepudi/karayaml/releases")
		os.Exit(1)
	}

	needsUpgrade := NeedsUpgrade(currentVersion, latestVersion)

	fmt.Println()
	if needsUpgrade {
		fmt.Printf("Current version: %s\n", currentVersion)
		fmt.Printf("Latest version:  %s\n", latestVersion)
	} else {
		fmt.Printf("Current version: %s\n", currentVersion)
		fmt.Printf("Latest version:  %s\n", latestVersion)
	}

	if !needsUpgrade && !force {
		fmt.Println()
		fmt.Printf("✓ karayaml is already up to date (%s)\n", currentVersion)
		return
	}

	if checkOnly {
		if needsUpgrade {
			fmt.Println()
			fmt.Printf("⚡ A new version (%s) is available!\n", latestVersion)
			fmt.Println()
			fmt.Println("Run 'karayaml upgrade' to update.")
		}
		return
	}

	if !needsUpgrade && force {
		fmt.Println()
		fmt.Println("→ Forcing upgrade...")
	}

	method := DetectUpgradeMethod()
	fmt.Printf("\n→ Upgrade method: %s\n", method.String())

	var upgradeErr error
	switch method {
	case MethodHomebrew:
		upgradeErr = UpgradeViaHomebrew()
	case MethodDirectDownload:
		upgradeErr = UpgradeViaDirect(latestVersion)
	}

	if upgradeErr != nil {
		handleUpgradeError(upgradeErr, latestVersion)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf("✓ Successfully upgraded to %s\n", latestVersion)

	if method == MethodDirectDownload {
		fmt.Println()
		fmt.Println("→ Note: You may need to restart your terminal for changes to take effect.")
	}
}

func handleUpgradeError(err error, latestVersion string) {
	fmt.Println()

	var permErr *PermissionError
	if errors.As(err, &permErr) {
		fmt.Fprintf(os.Stderr, "✗ %s\n", permErr.Error())
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Try running with sudo:")
		fmt.Fprintln(os.Stderr, "  sudo karayaml upgrade")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Or download manually:")
		goos, goarch := GetPlatformInfo()
		downloadURL := BuildDownloadURL(latestVersion, goos, goarch)
		fmt.Fprintf(os.Stderr, "  curl -LO %s\n", downloadURL)
		if goos != "windows" {
			fmt.Fprintln(os.Stderr, "  tar -xzf karayaml_*.tar.gz")
			fmt.Fprintln(os.Stderr, "  chmod +x karayaml")
			fmt.Fprintln(os.Stderr, "  sudo mv karayaml /usr/local/bin/")
		}
		return
	}

	fmt.Fprintf(os.Stderr, "✗ Upgrade failed: %v\n", err)
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "You can manually download the latest version from:")
	fmt.Fprintln(os.Stderr, "  https://github.com/swarupdonepudi/karayaml/releases")

	goos, goarch := GetPlatformInfo()
	downloadURL := BuildDownloadURL(latestVersion, goos, goarch)

	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Or download directly:")
	if runtime.GOOS == "windows" {
		fmt.Fprintf(os.Stderr, "  Invoke-WebRequest -Uri \"%s\" -OutFile \"karayaml.zip\"\n", downloadURL)
	} else {
		fmt.Fprintf(os.Stderr, "  curl -LO %s\n", downloadURL)
		fmt.Fprintln(os.Stderr, "  tar -xzf karayaml_*.tar.gz")
		fmt.Fprintln(os.Stderr, "  chmod +x karayaml")
		if runtime.GOOS == "darwin" {
			fmt.Fprintln(os.Stderr, "  xattr -dr com.apple.quarantine karayaml")
		}
		fmt.Fprintln(os.Stderr, "  sudo mv karayaml /usr/local/bin/")
	}
}
