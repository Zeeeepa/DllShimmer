#!/usr/bin/env bash

set -ueo pipefail

go run main.go -i '7z.dll' -p '7z2.dll' --def 'dll.def' -m > shim.cpp

x86_64-w64-mingw32-g++ -shared shim.cpp dll.def -o 7z.dll-shim -static-libstdc++ -static-libgcc -D DEBUG=1

mv 7z.dll-shim ~/vm/Windows\ 11/shared/7z.dll