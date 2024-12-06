# Define paths
BUILD_DIR := build
CONTAINERD_DIR := lib/containerd
CONTAINERD_BIN := $(BUILD_DIR)/bin/containerd
CONTAINERD_BUILD_CMD := make -C $(CONTAINERD_DIR)
CONTAINERD_RUN_CMD := $(BUILD_DIR)/containerd

# Targets
.PHONY: all init build run clean

# Default target: initialize submodule and build containerd
all: init build

# Initialize submodules
init:
	@echo "Initializing submodules..."
	git submodule update --init --recursive

# Build containerd
build:
	@echo "Building containerd..."
	$(CONTAINERD_BUILD_CMD)
	# copy lib/containerd to BUILD_DIR
	@cp -r $(CONTAINERD_DIR)/bin $(BUILD_DIR)

# Run containerd
run:
	@echo "Running containerd..."
	$(CONTAINERD_RUN_CMD)

# Clean build artifacts
clean:
	@echo "Cleaning containerd build artifacts..."
	$(CONTAINERD_BUILD_CMD) clean
	@rm -rf $(BUILD_DIR)