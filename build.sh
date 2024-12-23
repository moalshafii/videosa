#!/bin/bash

# Set the output directory for the builds
OUTPUT_DIR="bin"

# Create the output directory if it doesn't exist
mkdir -p $OUTPUT_DIR

# Function to build for a given OS and architecture
build_for_platform() {
  local os=$1
  local arch=$2
  local output_name="videosa-${os}-${arch}"

  echo "Building for ${os} ${arch}..."
  GOOS=$os GOARCH=$arch go build -o $OUTPUT_DIR/$output_name
  if [ $? -ne 0 ]; then
    echo "Build for ${os} ${arch} failed!"
    exit 1
  fi
}

# Build for Windows (64-bit)
build_for_platform windows amd64

# Build for Linux (64-bit)
build_for_platform linux amd64

# Build for macOS (64-bit)
build_for_platform darwin amd64

# Notify user of successful builds
echo "All builds completed successfully!"
