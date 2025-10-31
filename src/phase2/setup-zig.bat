@echo off
REM Script to download and install Zig locally for this project (Windows)

set ZIG_VERSION=0.11.0
set ZIG_DIR=%USERPROFILE%\.local\zig
set ZIG_ZIP=zig-windows-x86_64-%ZIG_VERSION%.zip
set ZIG_URL=https://ziglang.org/download/%ZIG_VERSION%/%ZIG_ZIP%

echo Setting up Zig %ZIG_VERSION% for Windows...

REM Create directory
if not exist "%ZIG_DIR%" mkdir "%ZIG_DIR%"

REM Download Zig using PowerShell
echo Downloading Zig from %ZIG_URL%...
cd /d "%ZIG_DIR%"
powershell -Command "Invoke-WebRequest -Uri '%ZIG_URL%' -OutFile '%ZIG_ZIP%'"

REM Extract using PowerShell
echo Extracting Zig...
powershell -Command "Expand-Archive -Path '%ZIG_ZIP%' -DestinationPath '.' -Force"
del "%ZIG_ZIP%"

set ZIG_EXTRACTED_DIR=%ZIG_DIR%\zig-windows-x86_64-%ZIG_VERSION%

echo.
echo Zig has been installed to: %ZIG_EXTRACTED_DIR%
echo.
echo To use Zig, add it to your PATH:
echo   set PATH=%ZIG_EXTRACTED_DIR%;%%PATH%%
echo.
echo Or add it permanently through System Environment Variables
echo.
echo Verify installation with:
echo   zig version
echo.

pause
