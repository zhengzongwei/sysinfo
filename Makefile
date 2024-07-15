# Makefile

# Define the Go application entry point file
ENTRY_FILE := main.go

# Define the output directory for the compiled binaries
OUTPUT_DIR := build

# List of target platforms and architectures

# Specific platforms
MAC_TARGETS := darwin/arm64 darwin/amd64
LINUX_TARGETS := linux/amd64 linux/arm64
WINDOWS_TARGETS := windows/amd64 windows/arm64 windows/386
TARGETS := $(MAC_TARGETS) $(LINUX_TARGETS) $(WINDOWS_TARGETS)
# Default target: build binaries for all platforms
all: $(TARGETS)
# Build binaries for macOS platforms
mac: $(MAC_TARGETS)
# Build binaries for Linux platforms
linux: $(LINUX_TARGETS)
# Build binaries for Windows platforms
windows: $(WINDOWS_TARGETS)

$(TARGETS):
	@echo "Building for $@"
	@mkdir -p $(OUTPUT_DIR)/$(word 1,$(subst /, ,$@))
	GOOS=$(firstword $(subst /, ,$@)) GOARCH=$(word 2,$(subst /, ,$@)) go build -o $(OUTPUT_DIR)/$(firstword $(subst /, ,$@))/app_$(subst /,_,$@)$(if $(filter windows,$(firstword $(subst /, ,$@))),.exe) $(ENTRY_FILE)
# Clean the compiled binaries
clean:
	@rm -rf $(OUTPUT_DIR)
	@echo "Cleaned up."

# Help target: display make targets and their descriptions
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  all       Build binaries for all platforms"
	@echo "  mac       Build binaries for macOS platforms"
	@echo "  linux     Build binaries for Linux platforms"
	@echo "  windows   Build binaries for Windows platforms"
	@echo "  clean     Clean up compiled binaries"
	@echo "  help      Display this help message"

.PHONY: all mac linux windows clean help
