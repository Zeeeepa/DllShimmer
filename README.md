# DllShimmer (WORK IN PROGRESS!)

Compile everything:

```bash

x86_64-w64-mingw32-g++ -shared -o version.dll-shim shim.cpp version.def -static-libstdc++ -static-libgcc
```

Tips:

1. Zawsze najpierw zrób dynamic linking.

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

TODO:

- Is dynamic loading also relative to EXE or callee DLL?

## Options

### `-i / --input <file>` [required]

The original DLL that you want to backdoor.

### `-o / --output <dir>` [required]

The path to the directory where DllShimmer will save all generated files.

### `-x / --original-path <path | file>` [required]

In case of dynamic linking (default) provide the path where the proxy DLL will find the original DLL on the target system.

In the case of static linking (`--static`), specify only the name of the original DLL. It will be searched for according to the default loading order on Windows.

### `-m / --mutex` [optional]

Enabling this option will add a mutex to the source file, which prevents your backdoor from being executed more than once during a single program run. All original functions will continue to work normally.

### `--static` [optional]

TBD;

## Troubleshooting

### 1. Strange loader error (126) while loading original DLL

Sometimes, your proxy DLL displays an error when loading the original DLL, and the error code is 126, even though you theoretically specified the correct relative path in the `--proxy` parameter. Why isn't it working?!?

DLLs are searched for in the `Current Directory`. In 98% of cases, this is simply the location of the main EXE file, but there are programs (mostly old legacy ones) that arbitrarily change the `Current Directory` using, for example, `SetCurrentDirectoryW()`. The main program is aware of this change, so it loads your proxy DLL correctly, but you are unaware of this and try to load the original DLL relatively, while the program searches for it in the changed `Current Directory`.

This rule applies to both static and dynamic loading of the original DLL. Unfortunately, with static linking, this problem is much harder to detect because we don't have debug information. System loader just fails and it's over. This is why I always recommend using the default dynamic linking first.

In the case of dynamic linking, we have two options:

1. Adjust the path in `--proxy` to the new `Current Directory` situation.
2. Change the `Current Directory` dynamically to search for DLLs where we want.

In case of static linking, we really only have one option:

1. Move the original DLL to the `Current Directory`.
