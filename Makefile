# Makefile

# Define the Go application entry point file
ENTRY_FILE := main.go

# Define the output directory for the compiled binaries
OUTPUT_DIR := build

# List of target platforms and architectures
# https://go.dev/doc/install/source#environment
TARGETS := linux/amd64 linux/arm64 darwin/arm64 darwin/amd64 windows/amd64 windows/arm64 windows/386

# Default target: build binaries for all platforms
all: $(TARGETS)

# Build binary for a specific target
$(TARGETS):
	@echo "Building for $@"
	@mkdir -p $(OUTPUT_DIR)/$(word 1,$(subst /, ,$@))
	#GOOS=$(word 1,$(subst /, ,$@)) GOARCH=$(word 2,$(subst /, ,$@)) go build -o $(OUTPUT_DIR)/$(word 1,$(subst /, ,$@))/app_$(subst /,_,$@)$(if $(filter windows,$(GOOS)),.exe)
	GOOS=$(firstword $(subst /, ,$@)) GOARCH=$(word 2,$(subst /, ,$@)) go build -o $(OUTPUT_DIR)/$(firstword $(subst /, ,$@))/app_$(subst /,_,$@)$(if $(filter windows,$(firstword $(subst /, ,$@))),.exe)

# Clean the compiled binaries
clean:
	@rm -rf $(OUTPUT_DIR)
	@echo "Cleaned up."

# Help target: display make targets and their descriptions
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  all       Build binaries for all platforms"
	@echo "  clean     Clean up compiled binaries"
	@echo "  help      Display this help message"

.PHONY: all clean help
