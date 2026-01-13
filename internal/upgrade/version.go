package upgrade

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	githubReleasesURL = "https://api.github.com/repos/swarupdonepudi/karayaml/releases"
	httpTimeout       = 10 * time.Second
)

// GitHubRelease represents a GitHub release from the API
type GitHubRelease struct {
	TagName    string `json:"tag_name"`
	Draft      bool   `json:"draft"`
	Prerelease bool   `json:"prerelease"`
}

// GetLatestVersion fetches the latest version from GitHub releases
func GetLatestVersion() (string, error) {
	client := &http.Client{Timeout: httpTimeout}

	req, err := http.NewRequest("GET", githubReleasesURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch releases: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch releases: HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var releases []GitHubRelease
	if err := json.Unmarshal(body, &releases); err != nil {
		return "", fmt.Errorf("failed to parse releases: %w", err)
	}

	if len(releases) == 0 {
		return "", fmt.Errorf("no releases found")
	}

	// Find the highest version (simple string comparison works for semver)
	var highestTag string
	for _, release := range releases {
		if release.Draft || release.Prerelease {
			continue
		}
		if highestTag == "" || compareVersions(release.TagName, highestTag) > 0 {
			highestTag = release.TagName
		}
	}

	if highestTag == "" {
		return "", fmt.Errorf("no valid releases found")
	}

	return highestTag, nil
}

// compareVersions compares two semver strings, returns >0 if a > b
func compareVersions(a, b string) int {
	aParts := parseVersion(a)
	bParts := parseVersion(b)

	for i := 0; i < 3; i++ {
		if aParts[i] > bParts[i] {
			return 1
		}
		if aParts[i] < bParts[i] {
			return -1
		}
	}
	return 0
}

// parseVersion extracts major, minor, patch from a version string
func parseVersion(v string) [3]int {
	v = strings.TrimPrefix(v, "v")
	parts := strings.Split(v, ".")
	var result [3]int
	for i := 0; i < len(parts) && i < 3; i++ {
		fmt.Sscanf(parts[i], "%d", &result[i])
	}
	return result
}

// NeedsUpgrade returns true if latestVersion is newer than currentVersion
func NeedsUpgrade(currentVersion, latestVersion string) bool {
	if currentVersion == latestVersion {
		return false
	}
	current := strings.TrimPrefix(currentVersion, "v")
	if current == "dev" || current == "" {
		return true
	}
	return compareVersions(latestVersion, currentVersion) > 0
}
