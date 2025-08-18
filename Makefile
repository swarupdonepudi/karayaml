version?=dev
name=karayaml
pkg=github.com/swarupdonepudi/karayaml
build_dir=build
LDFLAGS=-ldflags "-X ${pkg}/cmd/karayaml/root.VersionLabel=${version}"

build_cmd=go build -v ${LDFLAGS}

.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: build_darwin
build_darwin: vet
	GOOS=darwin ${build_cmd} -o ${build_dir}/${name}-darwin .

.PHONY: build
build: ${build_dir}/${name}

${build_dir}/${name}: deps vet
	GOOS=darwin GOARCH=amd64 ${build_cmd} -o ${build_dir}/${name}-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 ${build_cmd} -o ${build_dir}/${name}-darwin-arm64 .

.PHONY: test
test:
	go test -race -v -count=1 ./...

.PHONY: run
run: build
	${build_dir}/${name}

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf ${build_dir}

.PHONY: local
local: build_darwin
	rm -f ${HOME}/bin/${name}
	cp ./${build_dir}/${name}-darwin ${HOME}/bin/${name}
	chmod +x ${HOME}/bin/${name}

.PHONY: release-github
release-github:
	git tag ${version}
	git push origin ${version}
	gh release create ${version} \
		 --generate-notes \
        --title ${version} \
        build/karayaml-darwin-amd64 \
        build/karayaml-darwin-arm64

.PHONY: release
release: build test release-github

.PHONY: develop-site
develop-site:
	cd site && npm install --no-audit --no-fund
	cd site && npm run dev

.PHONY: preview-site
preview-site:
	cd site && npm install --no-audit --no-fund
	cd site && npm run build:serve
