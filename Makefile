all: help

## help: Show this help message
.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## generate: Generate the protobuf files
.PHONY: generate
generate:
	buf generate

## build: Build the protobuf files
.PHONY: build
build:
	buf build

## breaking: Check for breaking changes
.PHONY: breaking
breaking:
	buf breaking --against-registry

## lint: Lint the protobuf files
.PHONY: lint
lint:
	buf lint

## format: Format the protobuf files
.PHONY: format
format:
	buf format

## push: Push the protobuf files
.PHONY: push
push:
	buf push

## version-major: Increment the major version based on the latest git tag
.PHONY: version-major
version-major:
	@current_version=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	version_without_v=$$(echo $$current_version | sed 's/^v//'); \
	major=$$(echo $$version_without_v | cut -d. -f1); \
	minor=0; \
	patch=0; \
	new_version="v$$((major+1)).$$minor.$$patch"; \
	echo "Current version: $$current_version"; \
	echo "New version: $$new_version"; \
	git tag -a $$new_version -m "Release $$new_version"; \
	git push --follow-tags; \
	echo "Successfully created and pushed tag $$new_version"

## version-minor: Increment the minor version based on the latest git tag
.PHONY: version-minor
version-minor:
	@current_version=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	version_without_v=$$(echo $$current_version | sed 's/^v//'); \
	major=$$(echo $$version_without_v | cut -d. -f1); \
	minor=$$(echo $$version_without_v | cut -d. -f2); \
	patch=0; \
	new_version="v$$major.$$((minor+1)).$$patch"; \
	echo "Current version: $$current_version"; \
	echo "New version: $$new_version"; \
	git tag -a $$new_version -m "Release $$new_version"; \
	git push --follow-tags; \
	echo "Successfully created and pushed tag $$new_version"

## version-patch: Increment the patch version based on the latest git tag
.PHONY: version-patch
version-patch:
	@current_version=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	version_without_v=$$(echo $$current_version | sed 's/^v//'); \
	major=$$(echo $$version_without_v | cut -d. -f1); \
	minor=$$(echo $$version_without_v | cut -d. -f2); \
	patch=$$(echo $$version_without_v | cut -d. -f3); \
	new_version="v$$major.$$minor.$$((patch+1))"; \
	echo "Current version: $$current_version"; \
	echo "New version: $$new_version"; \
	git tag -a $$new_version -m "Release $$new_version"; \
	git push --follow-tags; \
	echo "Successfully created and pushed tag $$new_version"
