TAG      := 0.$(shell date +%Y%m%d).$(shell git describe --always)
VERSION	     ?= 3.3.0-dev
PUBLISH      := 0

ifdef $$VERSION
VERSION := $$VERSION
endif

ifeq ($(PUBLISH),0)
publish_option = --skip=publish
VERSION_NAME := $(VERSION)-$(TAG)
else
VERSION_NAME := v$(VERSION)
endif

releaser_options = --skip=validate

$(shell echo $(VERSION) > .version)
$(shell echo $(TAG) > .release)

GORELEASER := goreleaser --clean
.PHONY: release
release:
	git tag $(VERSION_NAME) || true
	$(GORELEASER) $(publish_option) $(releaser_options) --config .goreleaser.yaml
	$(GORELEASER) $(releaser_options) --config .goreleaser-docker.yaml

.PHONY: clean
clean:
	@rm -Rf release release-docker .version .release
