#include <stdio.h>
#include <iostream>
// Put your imports here...

// #------------------------------------------------------------------#
// |                        "DON'T TOUCH" ZONE                        |
// |                         (auto generated)                         |
// #------------------------------------------------------------------#
#define CreateDecoder CreateDecoderOriginal
#define CreateEncoder CreateEncoderOriginal
#define CreateObject CreateObjectOriginal
#define GetHandlerProperty GetHandlerPropertyOriginal
#define GetHandlerProperty2 GetHandlerProperty2Original
#define GetHashers GetHashersOriginal
#define GetIsArc GetIsArcOriginal
#define GetMethodProperty GetMethodPropertyOriginal
#define GetModuleProp GetModulePropOriginal
#define GetNumberOfFormats GetNumberOfFormatsOriginal
#define GetNumberOfMethods GetNumberOfMethodsOriginal
#define SetCaseSensitive SetCaseSensitiveOriginal
#define SetCodecs SetCodecsOriginal
#define SetLargePageMode SetLargePageModeOriginal
#include <windows.h>
#undef CreateDecoder
#undef CreateEncoder
#undef CreateObject
#undef GetHandlerProperty
#undef GetHandlerProperty2
#undef GetHashers
#undef GetIsArc
#undef GetMethodProperty
#undef GetModuleProp
#undef GetNumberOfFormats
#undef GetNumberOfMethods
#undef SetCaseSensitive
#undef SetCodecs
#undef SetLargePageMode


#define MUTEX(name) \
    (CreateMutexA(NULL, TRUE, name) && GetLastError() != ERROR_ALREADY_EXISTS)

#define ARGS_COUNT 12

typedef uint64_t (*Func12)(
    uint64_t, uint64_t, uint64_t, uint64_t,
    uint64_t, uint64_t, uint64_t, uint64_t,
    uint64_t, uint64_t, uint64_t, uint64_t
);

#define PROXY_FUNCTION(function)                                                 \
    va_list ap;                                                                  \
    va_start(ap, arg1);                                                          \
    uint64_t args[ARGS_COUNT];                                                   \
    args[0] = arg1;                                                              \
                                                                                 \
    for (int i = 1; i < ARGS_COUNT ; i++) {                                      \
        args[i] = va_arg(ap, uint64_t);                                          \
    }                                                                            \
                                                                                 \
    va_end(ap);                                                                  \
                                                                                 \
    HMODULE hModule = LoadLibraryA("7z2.dll");                             \
    if (hModule == NULL) {                                                       \
        printf("[!] 7z.dll: LoadLibraryA(7z2.dll) failed\n");      \
        printf("\tError code: %lu\n", GetLastError());                           \
    }                                                                            \
                                                                                 \
    Func12 pFunction = (Func12) GetProcAddress(hModule, function);               \
    if (pFunction == NULL) {                                                     \
        printf(                                                                  \
            "[!] 7z.dll: GetProcAddress(%s, 7z2.dll) failed\n",     \
            function );                                                          \
        printf("\tError code: %lu\n", GetLastError());                           \
                                                                                 \
    }                                                                            \
                                                                                 \
    return pFunction(args[0], args[1], args[2], args[3], args[4], args[5],       \
                        args[6], args[7], args[8], args[9], args[10], args[11]); \

// ---- Forwarded functions ------------------------------------------

// #------------------------------------------------------------------#
// |                    END OF "DON'T TOUCH" ZONE                     |
// #------------------------------------------------------------------#

extern "C" __declspec(dllexport) UINT64 CreateDecoder(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: CreateDecoder called\n");
    #endif
    
    if (MUTEX("Global\\CreateDecoder__0")) {
        // Put your code here...
    }

    PROXY_FUNCTION("CreateDecoder");
}

extern "C" __declspec(dllexport) UINT64 CreateEncoder(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: CreateEncoder called\n");
    #endif
    
    if (MUTEX("Global\\CreateEncoder__1")) {
        // Put your code here...
    }

    PROXY_FUNCTION("CreateEncoder");
}

extern "C" __declspec(dllexport) UINT64 CreateObject(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: CreateObject called\n");
    #endif
    
    if (MUTEX("Global\\CreateObject__2")) {
        // Put your code here...
    }

    PROXY_FUNCTION("CreateObject");
}

extern "C" __declspec(dllexport) UINT64 GetHandlerProperty(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: GetHandlerProperty called\n");
    #endif
    
    if (MUTEX("Global\\GetHandlerProperty__3")) {
        // Put your code here...
    }

    PROXY_FUNCTION("GetHandlerProperty");
}

extern "C" __declspec(dllexport) UINT64 GetHandlerProperty2(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: GetHandlerProperty2 called\n");
    #endif
    
    if (MUTEX("Global\\GetHandlerProperty2__4")) {
        // Put your code here...
    }

    PROXY_FUNCTION("GetHandlerProperty2");
}

extern "C" __declspec(dllexport) UINT64 GetHashers(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: GetHashers called\n");
    #endif
    
    if (MUTEX("Global\\GetHashers__5")) {
        // Put your code here...
    }

    PROXY_FUNCTION("GetHashers");
}

extern "C" __declspec(dllexport) UINT64 GetIsArc(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: GetIsArc called\n");
    #endif
    
    if (MUTEX("Global\\GetIsArc__6")) {
        // Put your code here...
    }

    PROXY_FUNCTION("GetIsArc");
}

extern "C" __declspec(dllexport) UINT64 GetMethodProperty(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: GetMethodProperty called\n");
    #endif
    
    if (MUTEX("Global\\GetMethodProperty__7")) {
        // Put your code here...
    }

    PROXY_FUNCTION("GetMethodProperty");
}

extern "C" __declspec(dllexport) UINT64 GetModuleProp(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: GetModuleProp called\n");
    #endif
    
    if (MUTEX("Global\\GetModuleProp__8")) {
        // Put your code here...
    }

    PROXY_FUNCTION("GetModuleProp");
}

extern "C" __declspec(dllexport) UINT64 GetNumberOfFormats(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: GetNumberOfFormats called\n");
    #endif
    
    if (MUTEX("Global\\GetNumberOfFormats__9")) {
        // Put your code here...
    }

    PROXY_FUNCTION("GetNumberOfFormats");
}

extern "C" __declspec(dllexport) UINT64 GetNumberOfMethods(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: GetNumberOfMethods called\n");
    #endif
    
    if (MUTEX("Global\\GetNumberOfMethods__10")) {
        // Put your code here...
    }

    PROXY_FUNCTION("GetNumberOfMethods");
}

extern "C" __declspec(dllexport) UINT64 SetCaseSensitive(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: SetCaseSensitive called\n");
    #endif
    
    if (MUTEX("Global\\SetCaseSensitive__11")) {
        // Put your code here...
    }

    PROXY_FUNCTION("SetCaseSensitive");
}

extern "C" __declspec(dllexport) UINT64 SetCodecs(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: SetCodecs called\n");
    #endif
    
    if (MUTEX("Global\\SetCodecs__12")) {
        // Put your code here...
    }

    PROXY_FUNCTION("SetCodecs");
}

extern "C" __declspec(dllexport) UINT64 SetLargePageMode(UINT64 arg1, ...) {
    #ifdef DEBUG
        printf("[+] 7z.dll: SetLargePageMode called\n");
    #endif
    
    if (MUTEX("Global\\SetLargePageMode__13")) {
        // Put your code here...
    }

    PROXY_FUNCTION("SetLargePageMode");
}


BOOL WINAPI DllMain(HINSTANCE hinstDLL, DWORD fdwReason, LPVOID lpvReserved) { 
    
    switch (fdwReason) {
    case DLL_PROCESS_ATTACH: {
        #ifdef DEBUG
            printf("[+] 7z.dll: DLL_PROCESS_ATTACH event\n");
        #endif
    }
    case DLL_THREAD_ATTACH:
    case DLL_THREAD_DETACH:
    case DLL_PROCESS_DETACH:
        break;
    }

    return TRUE;
}