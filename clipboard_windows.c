#define UNICODE
#include <windows.h>

extern void GoClipboardChanged();

LRESULT CALLBACK HiddenWndProc(HWND hwnd, UINT msg, WPARAM wParam, LPARAM lParam) {
    switch (msg) {
        case WM_CLIPBOARDUPDATE:
            GoClipboardChanged();
            return 0;
    }
    return DefWindowProc(hwnd, msg, wParam, lParam);
}
