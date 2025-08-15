# DllShimmer

Compile everything:

```bash

x86_64-w64-mingw32-g++ -shared -o version.dll-shim shim.cpp version.def -static-libstdc++ -static-libgcc

```

Features:

- All functions implemented in the original DLL can be backdoored.
- Backdoored functions work as original, program doesn't crash.
- Forwarded functions are forwarded as original.
- Both MSVC (`#pragma comment`) and GCC forwarding (`.def file`) are supported.

Caveats:

1. Probably it doesn't work with floating-point parameters because they require different registers from va_list.
2. Original ordinal numbers of exported and forwarded functions are not preserved
3. There are some huge obfuscated DLLs with weird name mangling and tricks (e.g. Qt framework DLL). I don't recommend to use them as a backdoor base. Just use some normal DLL with 10-30 exported functions and it's going to work perfectly.

## Nowa architektura

Najlepiej by było, gdyby original.dll lądował w IAT backdoor.dll ze wszystkimi funkcjami. Wtedy mamy wszystko dostępne od razu, bez dodatkowego używania WinAPI.

Pojawia się wtedy problem, że nie można importować i eksportować tych samych symboli.

- Importuj oryginalne