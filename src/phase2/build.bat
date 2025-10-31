@echo off

REM Check for zig
where zig >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo Error: zig is required but not found in PATH
    echo Please install zig from https://ziglang.org/download/
    exit /b 1
)

echo Building Nuked-OPM library with zig cc...
cd lib
zig cc -c -I. nuked-opm\opm.c -o nuked-opm\opm.o
zig cc -c -I. opm_wrapper.c -o opm_wrapper.o
zig ar rcs libopm.a nuked-opm\opm.o opm_wrapper.o
cd ..

echo Building Go program with zig cc...
set CC=zig cc
set CXX=zig c++
go build -o phase2.exe

echo Build complete! Run phase2.exe to generate WAV file.
