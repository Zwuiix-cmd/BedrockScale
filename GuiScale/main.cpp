#include <windows.h>
#include <tlhelp32.h>
#include <tchar.h>
#include <iostream>

using namespace std;

// Retrieves process ID from process name
DWORD GetProcessID(const char* processName)
{
    HANDLE snapshot = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);
    PROCESSENTRY32 processEntry;
    processEntry.dwSize = sizeof(processEntry);

    if (snapshot == INVALID_HANDLE_VALUE)
        return 0;

    if (Process32First(snapshot, &processEntry))
    {
        do
        {
            char convertedName[MAX_PATH];
            WideCharToMultiByte(CP_ACP, 0, processEntry.szExeFile, -1, convertedName, MAX_PATH, NULL, NULL);

            if (strcmp(convertedName, processName) == 0)
            {
                CloseHandle(snapshot);
                return processEntry.th32ProcessID;
            }
        } while (Process32Next(snapshot, &processEntry));
    }
    CloseHandle(snapshot);
    return 0;
}

// Retrieves module base address from module name and process ID
uintptr_t GetModuleBaseAddress(const char* moduleName, DWORD processID)
{
    HANDLE snapshot = CreateToolhelp32Snapshot(TH32CS_SNAPMODULE, processID);
    MODULEENTRY32 moduleEntry;
    moduleEntry.dwSize = sizeof(moduleEntry);

    if (snapshot == INVALID_HANDLE_VALUE)
        return 0;

    if (Module32First(snapshot, &moduleEntry))
    {
        do
        {
            // Converts WCHAR to char
            char convertedName[MAX_MODULE_NAME32];
            WideCharToMultiByte(CP_ACP, 0, moduleEntry.szModule, -1, convertedName, MAX_MODULE_NAME32, NULL, NULL);

            if (strcmp(convertedName, moduleName) == 0)
                break;
        } while (Module32Next(snapshot, &moduleEntry));
    }
    CloseHandle(snapshot);
    return (uintptr_t)moduleEntry.modBaseAddr;
}

// Written to process memory
void WriteToMemory(uintptr_t memoryAddress, float value)
{
    const char* targetProcess = "Minecraft.Windows.exe";
    DWORD processID = GetProcessID(targetProcess);
    if (processID == 0) {
        cout << "Failed to attach process.";
        return;
    }
    cout << "Successfully attach process." << endl;
    cout << endl << "Pid: " << processID << endl;

    uintptr_t moduleBaseAddress = GetModuleBaseAddress(targetProcess, processID);
    HANDLE processHandle = OpenProcess(PROCESS_ALL_ACCESS, NULL, processID);
    uintptr_t writeAddress = moduleBaseAddress + memoryAddress;

    cout << "Memory: " << targetProcess << "+0x" << hex << memoryAddress << endl;

    DWORD oldProtect;
    VirtualProtectEx(processHandle, (LPVOID)(writeAddress), sizeof(value), PAGE_EXECUTE_READWRITE, &oldProtect);
    WriteProcessMemory(processHandle, (LPVOID)(writeAddress), &value, sizeof(value), nullptr);

    VirtualProtectEx(processHandle, (LPVOID)(writeAddress), sizeof(value), oldProtect, &oldProtect);
    CloseHandle(processHandle);

    cout << endl << "Memory writing (" << targetProcess << "+0x" << hex << memoryAddress << ")" << " to " << value << endl;
    cout << "You can now close the program." << endl;
}

int main()
{
    // "1.20.41"
    WriteToMemory(0x492A4A8, 7);
    while (true) {}
    return 0;
}
