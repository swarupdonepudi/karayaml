package defaulteditor

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// NewEditorCommand returns an *exec.Cmd that opens the provided file in an editor
// and blocks until the editor exits. Resolution order:
// 1) $VISUAL
// 2) $EDITOR
// 3) VS Code `code --wait`
// 4) macOS TextEdit via AppleScript waiting for the document window to close
// 5) fallback to `vi`
func NewEditorCommand(filePath string) *exec.Cmd {
	// Prefer VISUAL, then EDITOR
	editor := strings.TrimSpace(os.Getenv("VISUAL"))
	if editor == "" {
		editor = strings.TrimSpace(os.Getenv("EDITOR"))
	}

	// Helper to detect VS Code launcher
	isCode := func(name string) bool {
		base := name
		if idx := strings.LastIndex(name, "/"); idx != -1 {
			base = name[idx+1:]
		}
		return base == "code"
	}

	build := func(name string, extraArgs []string) *exec.Cmd {
		parts := strings.Fields(name)
		bin := parts[0]
		args := append(parts[1:], extraArgs...)
		args = append(args, filePath)
		return exec.Command(bin, args...)
	}

	if editor != "" {
		extra := []string{}
		if isCode(editor) {
			// Ensure we wait for the file to close
			// Simple check to avoid duplicating --wait
			if !strings.Contains(" "+editor+" ", " --wait ") {
				extra = append(extra, "--wait")
			}
		}
		// If binary exists in PATH or user provided explicit path, use it
		if fields := strings.Fields(editor); len(fields) > 0 {
			if _, err := exec.LookPath(fields[0]); err == nil || strings.Contains(fields[0], "/") {
				return build(editor, extra)
			}
		}
	}

	// Try VS Code if available
	if _, err := exec.LookPath("code"); err == nil {
		return exec.Command("code", "--wait", filePath)
	}

	// macOS: TextEdit with window-close (Cmd+W) detection.
	// `open -W` waits for the entire app to quit (Cmd+Q). Instead, run a small
	// AppleScript that opens the file and blocks until that specific document is closed.
	if runtime.GOOS == "darwin" {
		script := `on run argv
set thePath to POSIX file (item 1 of argv)
tell application "TextEdit"
	activate
	set theDoc to open thePath
	repeat while (exists theDoc)
		delay 0.2
	end repeat
end tell
end run`
		return exec.Command("osascript", "-e", script, filePath)
	}

	// Final fallback
	return exec.Command("vi", filePath)
}
