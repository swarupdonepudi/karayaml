# Site Refresh, Filter Subcommand Rename, and UX Polish

**Date**: February 15, 2026

## Summary

Renamed the filter subcommands from `filter-key`/`filter-app` to `by-key`/`by-app` (with `by-file` as an alias) for a more intuitive CLI. Refreshed the KaraYAML landing site with a real GitHub star badge, light default theme, improved Karabiner-Elements comparison section, pronoun cleanup, and visual polish for the hero keyboard shortcut demo.

## Problem Statement

Several areas needed attention after the initial filter command implementation:

### Pain Points

- `karayaml filter filter-key f` reads awkwardly -- "filter filter-key" is redundant
- The site used a generic lucide GitHub icon instead of the real GitHub mark, and did not show the star count
- Dark theme was the default, but light theme is more conventional for landing pages
- The "How we compare" section framed Karabiner-Elements negatively with red X icons, when KaraYAML is built on top of it and deeply grateful for it
- Several FAQ answers used pronouns ("it", "it's") instead of naming "KaraYAML" directly
- The hero visual's key and app icon components were nearly invisible in light theme due to very light backgrounds

## Solution

### CLI: Subcommand Rename

Renamed the filter subcommands for natural readability:

```bash
# Before (redundant)
karayaml filter filter-key f
karayaml filter filter-app auth

# After (clean)
karayaml filter by-key f
karayaml filter by-app auth    # alias: by-file
```

### Key Components

- **`cmd/karayaml/root/filter_key.go`** -- `Use` changed from `filter-key` to `by-key`
- **`cmd/karayaml/root/filter_app.go`** -- `Use` changed from `filter-app` to `by-app`, added `Aliases: []string{"by-file"}`

### Site: GitHub Star Badge

- Created `site/src/components/ui/GitHubStarBadge.jsx` -- fetches live star count from GitHub API, renders real GitHub SVG icon with star count
- Exported `GitHubIcon` for reuse across the site
- Replaced lucide `Github` icon in nav bar (now shows star badge), hero section, and footer

### Site: Default Light Theme

- `Layout.jsx` -- changed `isDark` default from `true` to `false`; users who previously saved a preference keep it

### Site: Comparison Section Rewrite

Completely rewrote the "How we compare" section:

- Heading: "Built on Karabiner-Elements" (was "How we compare")
- Added a gratitude callout with heart icon crediting the Karabiner-Elements team
- Changed from adversarial "cons vs pros" to complementary "Karabiner-Elements (the foundation)" + "KaraYAML adds (the developer experience)" -- both columns use green checkmarks
- Aligned the gratitude card and comparison cards inside a single `max-w-4xl` container so their edges match perfectly
- Arrow overlaid at the exact center between cards on desktop, shown inline on mobile

### Site: Pronoun Cleanup

Replaced pronouns with "KaraYAML" in:

- FAQ: "It will launch..." -> "KaraYAML launches..."
- FAQ: "It creates/updates..." -> "KaraYAML creates and updates..."
- FAQ: "It's not needed..." -> "KaraYAML is not available..."
- FAQ: "it's released" -> "KaraYAML is released"
- FeatureCards: "use it on any Mac" -> "use KaraYAML on any Mac"

### Site: Hero Visual Polish

- `Key` component: `bg-slate-100 dark:bg-slate-800` with `border-slate-300 dark:border-slate-600` -- visible in both themes
- `AppIcon` component: same treatment, replacing the near-invisible `bg-gradient-to-br from-primary/10 to-secondary/10`

### Site: Misc

- Updated copyright year from 2025 to 2026
- Footer links now point to real GitHub URL and scroll to site sections
- Updated CLI Reference to include all current commands (filter by-key/by-app, upgrade)
- Added `--exact` and `--debug` flags to the flags section
- Updated Quickstart step 4 to show `search`, `filter by-key`, `filter by-app`
- Added FAQ entry: "How do I find a specific shortcut?"

## Benefits

- More intuitive CLI: `filter by-key` reads naturally as a sentence
- Professional site with real GitHub presence and live star count
- Grateful, accurate positioning relative to Karabiner-Elements
- Consistent brand voice: "KaraYAML" named everywhere, no ambiguous pronouns
- Hero visual looks polished in both light and dark themes

## Impact

**Users**: Cleaner CLI experience and a site that accurately reflects KaraYAML's current capabilities.

**Community**: The comparison section now honors Karabiner-Elements instead of positioning against it, which better reflects the project's values.

**Files changed**: 10 modified, 1 new (GitHubStarBadge component).

## Related Work

Follows the initial filter command and alias implementation from the earlier changelog entry (`2026-02-15-060549-command-aliases-and-filter-command.md`).

---

**Status**: Production Ready
**Timeline**: Single session
