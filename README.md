# DllShimmer (WORK IN PROGRESS!)

Compile everything:

```bash

x86_64-w64-mingw32-g++ -shared -o version.dll-shim shim.cpp version.def -static-libstdc++ -static-libgcc
```

Features:

- Two linking methods: dynamic (LoadLibraryA) and static (via IAT).
- All functions implemented in the original DLL can be backdoored.
- Backdoored functions work as original, program doesn't crash.
- Forwarded functions are forwarded as original.
- Both MSVC (`#pragma comment`) and GCC forwarding (`.def file`) are supported.

Limitations:

1. Probably it doesn't work with floating-point parameters because they require different registers to be used in ABI but I might be wrong, TODO: check
2. There are some huge obfuscated DLLs with weird name mangling and tricks (e.g. Qt framework DLL). I don't recommend to use them as a backdoor base. Just use some normal DLL with 10-30 exported functions and it's going to work perfectly.
3. It supports only x86-64.
