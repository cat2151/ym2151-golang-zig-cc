# Zig CC Requirement

This project **requires** Zig as the C compiler. GCC and mingw are **not allowed**.

## Why Zig CC?

The project uses `zig cc` as a cross-platform C compiler that provides:
- Consistent compilation across platforms
- Better cross-compilation support
- No need for mingw on Windows

## Installation

### Quick Setup

We provide helper scripts to download and install Zig locally:

**Linux/Mac:**
```bash
./setup-zig.sh
export PATH="$HOME/.local/zig/zig-linux-x86_64-0.11.0:$PATH"
```

**Windows:**
```cmd
setup-zig.bat
set PATH=%USERPROFILE%\.local\zig\zig-windows-x86_64-0.11.0;%PATH%
```

### Manual Installation

1. Download Zig 0.11.0 or later from: https://ziglang.org/download/
2. Extract the archive
3. Add the zig binary to your PATH

### Verification

```bash
zig version
```

Should output something like: `0.11.0`

## Building

Once Zig is installed, use the build scripts:

```bash
./build.sh       # Linux/Mac
build.bat        # Windows
```

The build scripts will:
1. Check that zig is available
2. Compile C sources using `zig cc`
3. Build the Go program with CGo using `zig cc`

## Troubleshooting

### "zig: command not found" or "zig is not recognized"

Make sure Zig is in your PATH. After installation, you may need to:
- Open a new terminal/command prompt
- Re-source your shell configuration: `source ~/.bashrc`
- Verify with: `which zig` (Linux/Mac) or `where zig` (Windows)

### Build errors

If you see C compilation errors, ensure you're using Zig 0.11.0 or later:
```bash
zig version
```

## Note for CI/CD

In CI/CD environments (like GitHub Actions), you can use the `goto-bus-stop/setup-zig` action:

```yaml
- uses: goto-bus-stop/setup-zig@v2
  with:
    version: 0.11.0
```
