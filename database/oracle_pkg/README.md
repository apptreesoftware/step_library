## Building For windows

1. Install Cygwin
2. When on the Cygwin package page, make sure to choose the x86_64 GCC tools for install
3. Install Git bash
4. Build with
`env CC=x86_64-w64-mingw32-gcc go build -o main.exe`