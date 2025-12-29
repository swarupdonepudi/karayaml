# ── project metadata ────────────────────────────────────────────────────────────
name        := karayaml
pkg         := github.com/swarupdonepudi/karayaml
build_dir   := build
LDFLAGS     := -ldflags "-X $(pkg)/cmd/karayaml/root.VersionLabel=$$(git describe --tags --always --dirty)"

# ── helper vars ────────────────────────────────────────────────────────────────
build_cmd   := go build -v $(LDFLAGS)

# bump: major, minor, or patch (default)
bump ?= patch

# ── quality / housekeeping ─────────────────────────────────────────────────────
.PHONY: deps vet fmt test clean
deps:          ## download & tidy modules
	go mod download
	go mod tidy

vet:           ## go vet
	go vet ./...

fmt:           ## go fmt
	go fmt ./...

test: vet      ## run tests with race detector
	go test -race -v -count=1 ./...

clean:         ## remove build artifacts
	rm -rf $(build_dir)

# ── build ─────────────────────────────────────────────────────────────────────
.PHONY: build build_darwin
build_darwin: vet fmt
	GOOS=darwin $(build_cmd) -o $(build_dir)/$(name)-darwin .

build: deps vet fmt ## build CLI binaries
	GOOS=darwin GOARCH=amd64 $(build_cmd) -o $(build_dir)/$(name)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 $(build_cmd) -o $(build_dir)/$(name)-darwin-arm64 .

# ── local utility ──────────────────────────────────────────────────────────────
.PHONY: snapshot local
snapshot: deps ## build a local snapshot using GoReleaser
	goreleaser release --snapshot --clean --skip=publish

local: build_darwin ## build and install binary to ~/bin
	rm -f $(HOME)/bin/$(name)
	cp ./$(build_dir)/$(name)-darwin $(HOME)/bin/$(name)
	chmod +x $(HOME)/bin/$(name)

# ── release tagging ────────────────────────────────────────────────────────────
.PHONY: release build-check next-version
build-check:   ## quick compile to verify build
	go build -o /dev/null .

next-version:  ## show what the next version would be
	@latest=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	major=$$(echo $$latest | sed 's/v//' | cut -d. -f1); \
	minor=$$(echo $$latest | sed 's/v//' | cut -d. -f2); \
	patch=$$(echo $$latest | sed 's/v//' | cut -d. -f3); \
	case "$(bump)" in \
		major) major=$$((major + 1)); minor=0; patch=0 ;; \
		minor) minor=$$((minor + 1)); patch=0 ;; \
		patch) patch=$$((patch + 1)) ;; \
		*) echo "Invalid bump type: $(bump). Use major, minor, or patch"; exit 1 ;; \
	esac; \
	echo "v$$major.$$minor.$$patch"

release: test build-check ## auto-bump version, tag & push (bump=major|minor|patch, default: patch)
	@latest=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	major=$$(echo $$latest | sed 's/v//' | cut -d. -f1); \
	minor=$$(echo $$latest | sed 's/v//' | cut -d. -f2); \
	patch=$$(echo $$latest | sed 's/v//' | cut -d. -f3); \
	case "$(bump)" in \
		major) major=$$((major + 1)); minor=0; patch=0 ;; \
		minor) minor=$$((minor + 1)); patch=0 ;; \
		patch) patch=$$((patch + 1)) ;; \
		*) echo "Invalid bump type: $(bump). Use major, minor, or patch"; exit 1 ;; \
	esac; \
	version="v$$major.$$minor.$$patch"; \
	echo "Current version: $$latest"; \
	echo "Releasing: $$version ($(bump) bump)"; \
	git tag -a $$version -m "$$version"; \
	git push origin $$version

# ── textedit testing ───────────────────────────────────────────────────────────
.PHONY: textedit-test text
textedit-test:
	VISUAL= EDITOR="osascript $(CURDIR)/scripts/textedit_wait.applescript" go run ./main.go edit

text: textedit-test

# ── website (site/) ────────────────────────────────────────────────────────────
.PHONY: develop-site preview-site dev-site build-site
develop-site:
	cd site && npm install --no-audit --no-fund
	cd site && npm run dev

preview-site:
	cd site && npm install --no-audit --no-fund
	cd site && npm run build:serve

dev-site:
	cd site && yarn dev

build-site:
	cd site && yarn build

# ── default target ─────────────────────────────────────────────────────────────
.DEFAULT_GOAL := test
