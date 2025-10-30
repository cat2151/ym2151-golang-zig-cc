@echo off
echo Building Nuked-OPM library...
cd lib
gcc -c -I. nuked-opm\opm.c -o nuked-opm\opm.o
gcc -c -I. opm_wrapper.c -o opm_wrapper.o
ar rcs libopm.a nuked-opm\opm.o opm_wrapper.o
cd ..

echo Building Go program...
go build -o phase2.exe

echo Build complete! Run phase2.exe to generate WAV file.
