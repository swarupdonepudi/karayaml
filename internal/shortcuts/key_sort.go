package shortcuts

import (
	"sort"
	"strconv"
	"strings"
)

// classifyKey returns a category rank, a numeric value rank (when available),
// and the normalized key string for lexicographic fallback.
// Category order: 0=function keys, 1=numbers, 2=letters/others.
func classifyKey(key string) (int, int, string) {
	k := strings.ToLower(key)
	// Function keys: f1..f12
	if strings.HasPrefix(k, "f") {
		if n, err := strconv.Atoi(strings.TrimPrefix(k, "f")); err == nil {
			return 0, n, k
		}
	}
	// Numbers: 0..10
	if n, err := strconv.Atoi(k); err == nil {
		return 1, n, k
	}
	// Letters/others
	return 2, 0, k
}

// sortHotkeyStringsStable sorts keys in-place by function keys, then numbers, then letters.
func sortHotkeyStringsStable(keys []string) {
	sort.SliceStable(keys, func(i, j int) bool {
		ci, vi, si := classifyKey(keys[i])
		cj, vj, sj := classifyKey(keys[j])
		if ci != cj {
			return ci < cj
		}
		// For function keys and numbers, sort by numeric value
		if ci == 0 || ci == 1 {
			if vi != vj {
				return vi < vj
			}
			return si < sj
		}
		// For letters/others, sort lexicographically
		return si < sj
	})
}

// sortShortcutsStable sorts shortcuts in-place using the same ordering. It is
// stable, so entries with identical keys preserve their original relative order.
func sortShortcutsStable(shortcuts []*FileOpenShortcut) {
	sort.SliceStable(shortcuts, func(i, j int) bool {
		ci, vi, si := classifyKey(string(shortcuts[i].Key))
		cj, vj, sj := classifyKey(string(shortcuts[j].Key))
		if ci != cj {
			return ci < cj
		}
		if ci == 0 || ci == 1 {
			if vi != vj {
				return vi < vj
			}
			return si < sj
		}
		return si < sj
	})
}
