.PHONY: install clean

GOFLAGS ?= $(GOFLAGS:)
GO=go

DIST_ROOT=dist
DIST_FOLDER_NAME=mattermost-load-test
DIST_PATH=$(DIST_ROOT)/$(DIST_FOLDER_NAME)


all: install

.installdeps:
	glide cache-clear
	glide update
	touch .installdeps

install: .installdeps
	$(GO) install ./cmd/mcreate
	$(GO) install ./cmd/mmanage
	$(GO) install ./cmd/loadtest

package: install
	rm -rf $(DIST_ROOT)
	mkdir -p $(DIST_PATH)/bin

	cp $(GOPATH)/bin/mcreate $(DIST_PATH)/bin
	cp $(GOPATH)/bin/mmanage $(DIST_PATH)/bin
	cp $(GOPATH)/bin/loadtest $(DIST_PATH)/bin
	cp loadtestconfig.json $(DIST_PATH)
	cp setup.sh $(DIST_PATH)/bin
	cp run.sh $(DIST_PATH)/bin
	cp README.md $(DIST_PATH)
	
	tar -C $(DIST_ROOT) -czf $(DIST_PATH).tar.gz $(DIST_FOLDER_NAME)

setup:
	./setup.sh

run:
	./run.sh

clean:
	rm -f errors.log cache.db stats.log status.log
	rm -f ./cmd/mmange/mmange
	rm -f ./cmd/mcreate/mcreate
	rm -f ./cmd/loadtest/loadtest
	rm -r .installdeps
	rm -rf $(DIST_ROOT)
