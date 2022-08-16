# Pinned Versions
SNYK_DESKTOP_VERSION=1.827.0
SNYK_USER_VERSION=1.827.0
SNYK_OLD_VERSION=1.382.1
# Digest of the 1.827.0 snyk/snyk:docker image
SNYK_IMAGE_DIGEST=sha256:f9291a5310e3952369eeb8cd1c2a25f0c9fc930a3ccc88e1ea20956ad86b75a4
GO_VERSION=1.17.5
CLI_VERSION=20.10.11
ALPINE_VERSION=3.15.0
GOLANGCI_LINT_VERSION=v1.27.0-alpine
GOTESTSUM_VERSION=1.8.1
LTAG_VERSION=v0.2.3

GOOS ?= $(shell go env GOOS)
BINARY_EXT=
ifeq ($(GOOS),windows)
	BINARY_EXT=.exe
endif
PLATFORM_BINARY?=docker-scan_$(GOOS)_amd64$(BINARY_EXT)
BINARY=docker-scan$(BINARY_EXT)
