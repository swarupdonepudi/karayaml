# Command Aliases and Filter Command

**Date**: February 15, 2026

## Summary

Added `search` as an alias for the `find` command, `ls` as an alias for the `list` command, and introduced a new `filter` parent command with `filter-key` and `filter-app` subcommands. The filter subcommands support both fuzzy substring matching (default) and exact matching via an `--exact` flag. This improves CLI discoverability and gives users more targeted ways to look up their shortcuts.

## Problem Statement

Users were trying natural command names that didn't exist -- `karayaml search auth` and `karayaml ls` both returned "unknown command" errors. Additionally, the existing `find` command only searches by app/file name. There was no way to filter shortcuts by their mapped key (e.g., "show me everything mapped to keys starting with `f`").

### Pain Points

- `search` is a natural synonym for `find` but was not recognized
- `ls` is a universal shorthand for listing but was not recognized
- No way to filter shortcuts by key -- only by app name via `find`
- No way to do exact-match lookups (e.g., find only the shortcut mapped to exactly key `f`, not `f1`, `f2`, etc.)

## Solution

Leveraged Cobra's built-in `Aliases` field for the two alias additions -- zero friction, no additional wiring needed. For filtering, introduced a `filter` parent command with two subcommands (`filter-key` and `filter-app`), each accepting a positional query argument and an `--exact` flag.

### Key Components

- **`cmd/karayaml/root/find.go`** -- Added `Aliases: []string{"search"}`
- **`cmd/karayaml/root/list.go`** -- Added `Aliases: []string{"ls"}`
- **`cmd/karayaml/root/filter.go`** -- New parent command that registers `filter-key` and `filter-app`
- **`cmd/karayaml/root/filter_key.go`** -- Filters shortcuts by key with `--exact` flag
- **`cmd/karayaml/root/filter_app.go`** -- Filters shortcuts by app/file path with `--exact` flag
- **`internal/shortcuts/list.go`** -- New `FilterByKey()` and `FilterByApp()` functions
- **`cmd/karayaml/root.go`** -- Registered `root.Filter` in the root command

## Implementation Details

**Aliases** use Cobra's native `Aliases` field, which automatically registers the alias and shows it in help output:

```go
var Find = &cobra.Command{
    Use:     "find <query>",
    Aliases: []string{"search"},
    // ...
}
```

**Filter functions** follow the same pattern as the existing `Find()` function. Both `FilterByKey` and `FilterByApp` accept a `query` string and an `exact` bool. When `exact` is false, they do case-insensitive substring matching (`strings.Contains`). When true, they do case-insensitive equality (`==`).

**Filter commands** each maintain their own `--exact` flag via a package-level bool variable, registered in an `init()` function. Both reuse the existing `PrintMatches()` function for consistent table output.

## Benefits

- Users can now type `karayaml search` or `karayaml ls` as they'd naturally expect
- Shortcuts can be filtered by key, which was previously impossible
- Exact matching enables precise lookups when the fuzzy default is too broad
- Consistent UX: filter output uses the same table format as `find`
- No breaking changes to existing commands

## Impact

**Users**: Two fewer "unknown command" errors for common aliases. New filtering capability for power users managing many shortcuts.

**Developers**: Clean pattern for adding future filter dimensions (e.g., filter by key category). The `FilterByKey`/`FilterByApp` functions are simple and composable.

**Files changed**: 5 modified, 3 new (63 insertions across modified files + ~120 lines in new files).

## Related Work

This is the first feature addition to karayaml's CLI command set beyond the original `find`, `list`, `map`, `edit`, `init`, `reload`, and `upgrade` commands. The `filter` command pattern could be extended in the future to support filtering by key category (function keys, numbers, letters).

---

**Status**: Production Ready
**Timeline**: Single session
