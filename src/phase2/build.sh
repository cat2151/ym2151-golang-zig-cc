#!/bin/bash
set -e

# Check for zig
if ! command -v zig &> /dev/null; then
    echo "Error: zig is required but not found in PATH"
    echo "Please install zig from https://ziglang.org/download/"
    exit 1
fi

echo "Building Nuked-OPM library with zig cc..."
cd lib
zig cc -c -I. nuked-opm/opm.c -o nuked-opm/opm.o
zig cc -c -I. opm_wrapper.c -o opm_wrapper.o
zig ar rcs libopm.a nuked-opm/opm.o opm_wrapper.o
cd ..

echo "Building Go program with zig cc..."
CC="zig cc" CXX="zig c++" go build -o phase2.exe

echo "Build complete! Run ./phase2.exe to generate WAV file."
