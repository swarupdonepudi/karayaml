# KaraYAML

*A declarative YAML layer for [Karabiner‑Elements](https://karabiner-elements.pqrs.org/) on macOS.*

KaraYAML turns the time‑consuming “click and JSON edit” workflow into a **single YAML file you can version‑control**.  
Define your shortcuts once, run a command, and Karabiner‑Elements reloads with a fresh config every time.

---

## Key features

* **YAML in, Karabiner JSON out** – no manual JSON editing
* **Caps Lock Hyper‑key launcher** – map ⌃⌥⌘⇧ + <Key> to `open -a "<App>"`
* **Git‑friendly** – keep `~/.karayaml/shortcuts.yaml` in your dotfiles
* **One‑command editing** – `karayaml edit` opens the file in VS Code (or `$EDITOR`) and validates duplicates
* **Safe init** – `karayaml init` generates a working Karabiner config if you do not have one
* **Pretty list view** – `karayaml list` prints a clean table of your mappings
* **Simple Go binary** – no dependencies beyond `go 1.24+`

---

## Installation

```bash
brew install swarupdonepudi/tap/karayaml
```

(Or build from source: `go install github.com/swarupdonepudi/karayaml@latest`.)

---

## Quick start

```bash
# 1. Generate the default Karabiner config
karayaml init

# 2. Create or edit your YAML shortcuts
karayaml edit      # opens ~/.karayaml/shortcuts.yaml in VS Code

# 3. Reload Karabiner automatically (done by karayaml)
# 4. Profit – Caps Lock + <key> now launches your apps
```

---

## YAML schema

```yaml
- key: a                  # single character or 0‑9
  app_file_path: /Applications/Slack.app
- key: 1
  app_file_path: /System/Applications/Calendar.app
```

`key`        Single alphanumeric key (a‑z or 0‑9).  
`app_file_path` Absolute path to a macOS `.app` bundle.

---

## Commands

| Command            | Description                                                                                              |
|--------------------|----------------------------------------------------------------------------------------------------------|
| `karayaml init`    | Generate a default `~/.config/karabiner/karabiner.json` with Hyper‑key and arrow helpers.                |
| `karayaml edit`    | Open `~/.karayaml/shortcuts.yaml` in VS Code (or set `$EDITOR`). Prevents duplicate keys before closing. |
| `karayaml list`    | Print a table of current shortcuts.                                                                      |
| `karayaml version` | Show the CLI version.                                                                                    |

---

## How it works

1. **Edit** – You maintain a single YAML file under `~/.karayaml/shortcuts.yaml`.
2. **Transform** – KaraYAML converts each entry into a Karabiner “complex modification” rule that listens for **Caps
   Lock + <Key>**.
3. **Reload** – The CLI writes to `~/.config/karabiner/karabiner.json`; Karabiner auto‑detects the change and reloads in
   place.

No system restarts, no GUI clicks.

---

## Building from source

```bash
git clone https://github.com/swarupdonepudi/karayaml
cd karayaml
make build          # Darwin ARM64 and AMD64 binaries in ./build
```

---

## Contributing

Issues and PRs are welcome. Please:

1. Open an issue first for large changes.
2. Run `make fmt vet test` before pushing.
3. Keep commit messages concise and conventional.

---

## License

MIT © Swarup Donepudi
