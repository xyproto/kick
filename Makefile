.PHONY: all clean install static

DESTDIR ?=
PREFIX ?= /usr
UNAME_R ?= $(shell uname -r)

ifneq (,$(findstring arch,$(UNAME_R)))
# Arch Linux
LDFLAGS ?= -Wl,-O2,--sort-common,--as-needed,-z,relro,-z,now
BUILDFLAGS ?= -mod=vendor -buildmode=pie -trimpath -ldflags "-s -w -linkmode=external -extldflags $(LDFLAGS)"
else
# Default settings
BUILDFLAGS ?= -mod=vendor -trimpath
endif

# Build targets
all: kick mutator

kick:
	(cd cmd/kick && go build $(BUILDFLAGS) -o ../../kick)

mutator:
	(cd cmd/mutator && go build $(BUILDFLAGS) -o ../../mutator)

# Build statically (no dynamic libraries)
static:
	CGO_ENABLED=0 go build $(BUILDFLAGS) -ldflags "-s -w -extldflags '-static'" -o ../../kick ./cmd/kick
	CGO_ENABLED=0 go build $(BUILDFLAGS) -ldflags "-s -w -extldflags '-static'" -o ../../mutator ./cmd/mutator

# Installation target
install:
	install -Dm755 kick "$(DESTDIR)$(PREFIX)/bin/kick"
	install -Dm755 mutator "$(DESTDIR)$(PREFIX)/bin/mutator"

# Clean targets
clean:
	(cd cmd/kick && go clean)
	(cd cmd/mutator && go clean)
	rm -f kick mutator
