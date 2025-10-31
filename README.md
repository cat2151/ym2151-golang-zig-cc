# ym2151-golang-zig-cc

YM2151 (OPM) FM sound chip emulation using Nuked-OPM with Go and C.

## Projects

### Phase 2: WAV Generator

Located in `src/phase2/`, this program generates a 440Hz 3-second WAV file using the Nuked-OPM emulator.

See [src/phase2/README.md](src/phase2/README.md) for details and build instructions.

## Requirements

- Go 1.20 or later
- Zig 0.11.0 or later (https://ziglang.org/download/)
  - **Note: GCC and mingw are not allowed. You must use zig cc as the C compiler.**
- Windows platform (builds also work on Linux for testing)

## License

This project uses Nuked-OPM which is licensed under GNU LGPL 2.1.