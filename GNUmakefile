GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
PKG_NAME=enf

default: build

build:
	go install

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)

# Currently required by tf-deploy compile
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: build fmt fmtcheck
