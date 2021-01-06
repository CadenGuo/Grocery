# Go version to require, run go version to see what version you are using
GO_VERSION := "go1.15"
GO_VERSION ?= $(GO_VERSION)

# Note about difference between = and :=
# := means to not evaluate everytime it's expanded.
# = means to evaluate variable everytime it's expanded
# Reference: <http://www.gnu.org/software/make/manual/html_node/Flavors.html#Flavors>
VERSIONCMD = "`git symbolic-ref HEAD | cut -b 12-`-`git rev-parse HEAD`"
VERSION := $(shell echo $(VERSIONCMD))

# Later versions of git supports --count argument to rev-list subcommand
# but not the version on RHEL6 so for now we will just use wc-l
COMMIT_COUNT_CMD = "`git rev-list HEAD | wc -l`"
GIT_BRANCH_CMD = "`git symbolic-ref HEAD | awk -F / '{print $$3}'`"

COMMIT_COUNT := $(shell echo $(COMMIT_COUNT_CMD))
GIT_BRANCH   := $(shell echo $(GIT_BRANCH_CMD))
DATE := $(shell echo `date +%FT%T%z`)

# if there are any changes not committed, modify the version
CHANGES := $(shell echo `git status --porcelain | wc -l`)
ifneq ($(strip $(CHANGES)), 0)
    VERSION := dirty-build-$(VERSION)
endif

# Not to inherit from the outer LDFLAGS
LDFLAGS =
REMOVESYMBOL := -w -s
ifeq (true, $(DEBUG))
    REMOVESYMBOL = ""
endif
LDFLAGS += -X grocery/pkg/version.version=$(VERSION) -X grocery/pkg/version.date=$(DATE) $(REMOVESYMBOL)

BUILD_DIR := $(CURDIR)/build

export GOBIN := $(BUILD_DIR)/bin

.PHONY: all clean test lint build

all: clean test build

clean:
	rm -rf $(BUILD_DIR)

lint: install-linter
	go fmt $(shell go list ./...)
	golangci-lint run --timeout=5m \
		./cmd/... \
		./internal/... \
		./pkg/... \
		./tests/...

install-linter:
ifneq ('', '$(shell which golangci-lint)')
	@echo 'golangci-lint is already installed'
else
	brew install golangci-lint
	brew upgrade golangci-lint
endif

test:
	make check-go-version
	go test -v $(CURDIR)/...

build:
	make check-go-version
	make build-dirs
	go install -ldflags "$(LDFLAGS)" grocery/cmd/...

check-go-version:
	@if ! go version | grep "$(GO_VERSION)" >/dev/null; then \
        printf "Wrong go version: "; \
        go version; \
        echo "Requires go version: $(GO_VERSION)"; \
        exit 2; \
    fi

build-dirs:
	@for dir in $(BUILD_DIR)/{bin,}; do \
        test -d $$dir || mkdir -p $$dir; \
    done


lint: install-linter
	go fmt $(shell go list ./...)
	golangci-lint run --timeout=5m \
		./cmd/... \
		./internal/... \
		./pkg/... \

