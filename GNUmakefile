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

lint:
	@echo "==> Checking source code against linters..."
	@GOGC=30 golangci-lint run ./$(PKG_NAME)
	@tfproviderlint \
		-c 1 \
		-S001 \
		-S002 \
		-S003 \
		-S004 \
		-S005 \
		-S007 \
		-S008 \
		-S009 \
		-S010 \
		-S011 \
		-S012 \
		-S013 \
		-S014 \
		-S015 \
		-S016 \
		-S017 \
		-S019 \
		./$(PKG_NAME)

tools:
	GO111MODULE=on go install github.com/bflad/tfproviderlint/cmd/tfproviderlint
	GO111MODULE=on go install github.com/client9/misspell/cmd/misspell
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: build fmt fmtcheck lint tools