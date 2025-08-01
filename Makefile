#
# Copyright (c) 2019 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Determine the operating system
ifeq ($(OS), Windows_NT)
	UNAME := windows
else
	UNAME := $(shell uname|tr '[:upper:]' '[:lower:]')
	ifeq ($(UNAME),darwin)
	else ifeq ($(UNAME),linux)
	else
        # Unsupported OS
        # this block uses spaces instead of tabs so that we can do proper indentation enforcement
        $(error This Makefile only supports Linux, macOS (Darwin), and Windows.)
	endif
endif

ifneq (, $(shell command -v shasum))
SHA256CMD := shasum -a 256 --check
else ifneq (, $(shell command -v sha256sum))
SHA256CMD := sha256sum --check
else
$(error "please install 'shasum' or 'sha256sum'")
endif

goimports_version:=v0.4.0

.PHONY: check
check: metamodel
	./metamodel_generator/metamodel check --model=model

.PHONY: verify
verify: check verify-clientapi verify-openapi

update: clientapi openapi

.PHONY: openapi
openapi: metamodel
	./metamodel_generator/metamodel generate openapi --model=model --output=openapi

verify-openapi: metamodel
	$(eval TMPDIR := $(shell mktemp -d))
	./metamodel_generator/metamodel generate openapi --model=model --output=$(TMPDIR)
	diff -r $(TMPDIR)/ openapi/
	rm -rf $(TMPDIR)

clientapi: metamodel goimports-install
	./metamodel_generator/metamodel generate go \
		--model=model \
		--generators=types,builders,json \
		--base=github.com/openshift-online/ocm-api-model/clientapi \
		--output=clientapi
	pushd clientapi; go mod tidy; popd

verify-clientapi: metamodel goimports-install
	$(eval TMPDIR := $(shell mktemp -d))
	cp clientapi/go.mod $(TMPDIR)
	cp clientapi/go.sum $(TMPDIR)
	cp clientapi/LICENSE.txt $(TMPDIR)
	cp -r clientapi/dependencymagnet $(TMPDIR)
	./metamodel_generator/metamodel generate go \
		--model=model \
		--generators=types,builders,json \
		--base=github.com/openshift-online/ocm-api-model/clientapi \
		--output=$(TMPDIR)
	diff -r $(TMPDIR)/ clientapi/
	rm -rf $(TMPDIR)

metamodel:
	 $(MAKE) -C metamodel_generator clean
	 $(MAKE) -C metamodel_generator build

# Enforce indentation by tabs. License contains 2 spaces, so reject 3+.
lint:
	find -name '*.model' -print0 | xargs -0 sh -c '! egrep --with-filename --line-number "^(   |\t+ )" "$$@"'

.PHONY: clean
clean:
	rm -rf \
		./metamodel_generator/metamodel \
		metamodel \
		openapi \
		clientapi/accesstransparency \
        clientapi/accountsmgmt \
        clientapi/addonsmgmt \
        clientapi/arohcp \
        clientapi/authorizations \
        clientapi/clustersmgmt \
        clientapi/helpers \
        clientapi/jobqueue \
        clientapi/osdfleetmgmt \
        clientapi/servicelogs \
        clientapi/servicemgmt \
        clientapi/statusboard \
        clientapi/webrca \
		$(NULL)

.PHONY: validate-version
validate-version:
ifndef VERSION
	$(error VERSION is required. Usage: make tag VERSION=v0.0.426)
endif
	@# Validate that we're on the main branch
	@CURRENT_BRANCH=$$(git rev-parse --abbrev-ref HEAD); \
	if [ "$$CURRENT_BRANCH" != "main" ]; then \
		echo "Error: Tags must be created from the main branch. Currently on $$CURRENT_BRANCH"; \
		exit 1; \
	fi
	@# Validate version format (vX.Y.Z)
	@if ! echo "$(VERSION)" | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' > /dev/null; then \
		echo "Error: VERSION must follow the format vX.Y.Z (e.g., v0.0.426)"; \
		exit 1; \
	fi
	@# Get the latest tag and compare versions
	@LATEST_TAG=$$(git tag -l 'v0.*.*' | sort -V | tail -1); \
	if [ -n "$$LATEST_TAG" ]; then \
		if ! printf '%s\n%s\n' "$$LATEST_TAG" "$(VERSION)" | sort -V -C; then \
			echo "Error: VERSION $(VERSION) must be higher than the latest tag $$LATEST_TAG"; \
			exit 1; \
		fi; \
		if [ "$$LATEST_TAG" = "$(VERSION)" ]; then \
			echo "Error: VERSION $(VERSION) is the same as the latest tag $$LATEST_TAG"; \
			exit 1; \
		fi; \
	fi

MODULES = model clientapi metamodel_generator
.PHONY: tag-and-push-modules
tag-and-push-modules: validate-version
	@for module in $(MODULES); do \
		git tag -a -m 'Release $(VERSION)' $$module/$(VERSION); \
		git push upstream $$module/$(VERSION); \
	done 

.PHONY: release
release: tag-and-push-modules
	git tag -a -m 'Release $(VERSION)' $(VERSION) 
	git push upstream $(VERSION)

.PHONY: goimports-install
goimports-install:
	@GOBIN=$(LOCAL_BIN_PATH) go install golang.org/x/tools/cmd/goimports@$(goimports_version)

