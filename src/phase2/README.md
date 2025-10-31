# Phase 2: Nuked-OPM WAV Generator

This program generates a 3-second 440Hz (A4 note) WAV file using the Nuked-OPM YM2151 emulator.

## Features

- Uses Nuked-OPM cycle-accurate YM2151 emulator
- Generates a 440Hz tone for 3 seconds
- Implements 10ms cycle consumption after FM sound chip register writes
- Outputs stereo 16-bit PCM WAV file at 62.5kHz sample rate

## Building

### Prerequisites

- Go 1.20 or later
- Zig 0.11.0 or later (https://ziglang.org/download/)
  - **Note: GCC and mingw are not allowed. You must use zig cc as the C compiler.**

### Installing Zig

#### Option 1: Use the setup script (recommended for local development)
```bash
./setup-zig.sh       # On Linux/Mac
setup-zig.bat        # On Windows
```

Then add Zig to your PATH as instructed by the script.

#### Option 2: Manual installation
Download Zig from https://ziglang.org/download/ and extract it to a directory of your choice. Add the zig binary to your PATH.

**Linux/Mac:**
```bash
export PATH="/path/to/zig:$PATH"
```

**Windows:**
```cmd
set PATH=C:\path\to\zig;%PATH%
```

### Build Steps

1. Ensure zig is installed and in your PATH:
```bash
zig version
```

2. Use the build script:
```bash
./build.sh       # On Linux/Mac
build.bat        # On Windows
```

Or manually:

1. Compile the C library:
```bash
cd lib
zig cc -c -I. nuked-opm/opm.c -o nuked-opm/opm.o
zig cc -c -I. opm_wrapper.c -o opm_wrapper.o
zig ar rcs libopm.a nuked-opm/opm.o opm_wrapper.o
cd ..
```

2. Build the Go program:
```bash
CC="zig cc" CXX="zig c++" go build -o phase2.exe
```

## Running

```bash
./phase2.exe
```

This will generate `output.wav` in the current directory.

## Technical Details

### YM2151 Configuration

- Master clock: 4MHz
- Sample rate: 62.5kHz (4MHz / 64)
- Channel 0 configured with:
  - Key Code (KC): 0x4D for ~440Hz
  - Key Fraction (KF): 0x80
  - Algorithm: 7 (all operators as carriers)
  - Attack Rate: 31 (fastest)
  - Release Rate: 15 (fast release)

### Register Write Timing

After each FM register write operation, the program consumes 10ms worth of clock cycles to simulate real hardware timing constraints.

## License

- Nuked-OPM is licensed under GNU LGPL 2.1 (see lib/nuked-opm/LICENSE)
- This program code is part of the ym2151-golang-zig-cc project
