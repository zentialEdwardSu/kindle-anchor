PLATFORMS = linux/arm/7 linux/arm64 linux/amd64 windows/amd64

all: $(PLATFORMS)
	@echo "Build completed!"

$(PLATFORMS):
	$(eval platform_split = $(subst /, ,$@))
	$(eval GOOS = $(word 1,$(platform_split)))
	$(eval GOARCH = $(word 2,$(platform_split)))
	$(eval GOARM = $(word 3,$(platform_split)))

	$(eval output_name = anchordav-$(GOOS)-$(GOARCH))
	@if [ "$(GOARCH)" = "arm" ]; then \
		output_name=$(output_name)v7; \
	fi
	@if [ "$(GOOS)" = "windows" ]; then \
		output_name=$(output_name).exe; \
	fi

	@echo "Building for $(GOOS)/$(GOARCH)$(if $(GOARM),/$(GOARM))..."
	env GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -o ./kindle_anchor/$(output_name)
	@if [ $$? -ne 0 ]; then \
		echo 'An error has occurred! Aborting the build process...'; \
		exit 1; \
	fi

.PHONY: all $(PLATFORMS)
