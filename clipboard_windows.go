//go:build windows
// +build windows

package main

/*
#define UNICODE
#include <windows.h>

// 声明 HiddenWndProc（定义在 clipboard.c 内）
LRESULT CALLBACK HiddenWndProc(HWND hwnd, UINT msg, WPARAM wParam, LPARAM lParam);
*/
import "C"

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"

	"github.com/atotto/clipboard"
	"golang.org/x/sys/windows/registry"
)

const runKeyPath = `Software\Microsoft\Windows\CurrentVersion\Run`
const runValueName = "ClipboardBridge"

var hwndListener C.HWND

// utf16FromString 为 C 安全的 UTF16 分配
func utf16FromString(s string) *uint16 {
	utf := syscall.StringToUTF16(s)
	size := C.size_t(len(utf)) * C.size_t(unsafe.Sizeof(utf[0]))

	ptr := C.malloc(size)
	C.memcpy(ptr, unsafe.Pointer(&utf[0]), size)

	return (*uint16)(ptr)
}

//export GoClipboardChanged
func GoClipboardChanged() {
	text, _ := clipboard.ReadAll()

	clipboardMu.Lock()
	defer clipboardMu.Unlock()

	if text != lastClipboard {
		lastClipboard = text
		lastUpdated = time.Now().Unix()
		logInfo(fmt.Sprintf("检测到本地剪贴板更新，内容长度: %d 字节", len(text)))
		logDebug(fmt.Sprintf("剪贴板内容详情: %s", getPreview(text)))
	}
}

// initClipboardListener 初始化剪贴板监听 (Windows)
func initClipboardListener() {
	class := utf16FromString("ClipboardHiddenWindow")

	var wc C.WNDCLASSW
	wc.lpfnWndProc = C.WNDPROC(C.HiddenWndProc)
	wc.lpszClassName = (*C.WCHAR)(unsafe.Pointer(class))

	if C.RegisterClassW(&wc) == 0 {
		logInfo("注册监听窗口类失败")
		return
	}

	hwndListener = C.CreateWindowExW(
		0,
		wc.lpszClassName,
		wc.lpszClassName,
		0,
		0, 0, 0, 0,
		nil, nil, nil, nil,
	)

	if hwndListener == nil {
		logInfo("创建隐藏监听窗口失败")
		return
	}

	C.AddClipboardFormatListener(hwndListener)

	logInfo("剪贴板监听已启动")
}

// openLogFile 使用记事本打开日志文件 (Windows)
func openLogFile() {
	logPath := filepath.Join(exeDir(), "clipboard_bridge.log")
	cmd := exec.Command("notepad.exe", logPath)
	err := cmd.Start()
	if err != nil {
		logError("打开日志文件失败", err)
	} else {
		logInfo("已打开日志文件")
	}
}

// updateAutostart 更新开机自启 (Windows)
func updateAutostart() {
	exePath, err := os.Executable()
	if err != nil {
		logError("获取 exe 路径失败", err)
		return
	}

	k, _, err := registry.CreateKey(registry.CURRENT_USER, runKeyPath, registry.SET_VALUE)
	if err != nil {
		logError("打开 Run 注册表失败", err)
		return
	}
	defer k.Close()

	cfgMu.RLock()
	autoStart := cfg.AutoStart
	cfgMu.RUnlock()

	if autoStart {
		if err := k.SetStringValue(runValueName, exePath); err != nil {
			logError("写入开机自启失败", err)
		} else {
			logInfo("✓ 已设置开机自启")
		}
	} else {
		_ = k.DeleteValue(runValueName)
		logInfo("✓ 已取消开机自启")
	}

	// 保存配置到文件
	saveConfig()
}

// ensureFirewallRule 确保防火墙规则存在 (Windows)
func ensureFirewallRule() {
	cfgMu.RLock()
	autoFirewall := cfg.AutoFirewall
	port := cfg.Port
	cfgMu.RUnlock()

	if !autoFirewall {
		return
	}

	cmd := exec.Command("netsh", "advfirewall", "firewall",
		"add", "rule",
		"name=ClipboardBridge",
		"dir=in",
		"action=allow",
		"protocol=TCP",
		fmt.Sprintf("localport=%d", port),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		logInfo("添加防火墙规则失败（可能需要管理员权限或规则已存在）")
	} else {
		logInfo("已尝试添加防火墙规则")
	}
}

