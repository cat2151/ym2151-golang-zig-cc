#!/bin/bash
# Script to download and install Zig locally for this project

set -e

ZIG_VERSION="0.11.0"
ZIG_DIR="$HOME/.local/zig"
ZIG_TARBALL="zig-linux-x86_64-${ZIG_VERSION}.tar.xz"
ZIG_URL="https://ziglang.org/download/${ZIG_VERSION}/${ZIG_TARBALL}"

echo "Setting up Zig ${ZIG_VERSION}..."

# Create directory
mkdir -p "$ZIG_DIR"

# Download Zig
echo "Downloading Zig from ${ZIG_URL}..."
cd "$ZIG_DIR"

if command -v wget &> /dev/null; then
    wget -q --show-progress "${ZIG_URL}" -O "${ZIG_TARBALL}"
elif command -v curl &> /dev/null; then
    curl -L "${ZIG_URL}" -o "${ZIG_TARBALL}"
else
    echo "Error: wget or curl is required to download Zig"
    exit 1
fi

# Extract
echo "Extracting Zig..."
tar xf "${ZIG_TARBALL}"
rm "${ZIG_TARBALL}"

# Add to PATH
ZIG_EXTRACTED_DIR="${ZIG_DIR}/zig-linux-x86_64-${ZIG_VERSION}"

echo ""
echo "Zig has been installed to: ${ZIG_EXTRACTED_DIR}"
echo ""
echo "To use Zig, add it to your PATH:"
echo "  export PATH=\"${ZIG_EXTRACTED_DIR}:\$PATH\""
echo ""
echo "Or add this line to your ~/.bashrc or ~/.profile"
echo ""
echo "Verify installation with:"
echo "  zig version"
