# KaraYAML

*A declarative YAML layer for [Karabiner‑Elements](https://karabiner-elements.pqrs.org/) on macOS.*

KaraYAML turns the time‑consuming “click and JSON edit” workflow into a **single YAML file you can version‑control**.  
Define your shortcuts once, run a command, and Karabiner‑Elements reloads with a fresh config every time.

---

## Website & docs

Looking for features, quick start, CLI commands, schema, or "how it works"? Visit the website for the full documentation:

- **Open the website**: [https://karayaml.dev](https://karayaml.dev)

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
