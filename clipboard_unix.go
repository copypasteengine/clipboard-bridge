//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
)

// initClipboardListener 初始化剪贴板监听 (Unix/Linux/macOS)
// 注意：Unix 系统没有像 Windows 那样的系统级剪贴板监听机制
// 这里使用轮询方式检测剪贴板变化
func initClipboardListener() {
	logInfo("剪贴板监听已启动（轮询模式）")
	
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		
		for range ticker.C {
			text, err := clipboard.ReadAll()
			if err != nil {
				continue
			}
			
			clipboardMu.Lock()
			if text != lastClipboard {
				lastClipboard = text
				lastUpdated = time.Now().Unix()
				logDebug(fmt.Sprintf("检测到本地剪贴板更新，内容长度: %d 字节", len(text)))
			}
			clipboardMu.Unlock()
		}
	}()
}

// openLogFile 打开日志文件 (Unix/Linux/macOS)
func openLogFile() {
	logPath := filepath.Join(exeDir(), "clipboard_bridge.log")
	
	// 尝试使用系统默认编辑器打开
	var cmd *exec.Cmd
	switch {
	case fileExists("/usr/bin/xdg-open"): // Linux
		cmd = exec.Command("xdg-open", logPath)
	case fileExists("/usr/bin/open"): // macOS
		cmd = exec.Command("open", logPath)
	default:
		// 尝试常见的文本编辑器
		editors := []string{"gedit", "kate", "nano", "vim", "vi"}
		for _, editor := range editors {
			if _, err := exec.LookPath(editor); err == nil {
				cmd = exec.Command(editor, logPath)
				break
			}
		}
	}
	
	if cmd == nil {
		logError("未找到可用的文本编辑器", nil)
		return
	}
	
	err := cmd.Start()
	if err != nil {
		logError("打开日志文件失败", err)
	} else {
		logInfo("已打开日志文件")
	}
}

// fileExists 检查文件是否存在
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// updateAutostart 更新开机自启 (Unix/Linux/macOS)
func updateAutostart() {
	cfgMu.RLock()
	autoStart := cfg.AutoStart
	cfgMu.RUnlock()
	
	if autoStart {
		logInfo("开机自启功能需要手动配置（Unix/Linux/macOS）")
		logInfo("Linux: 将程序添加到 ~/.config/autostart/")
		logInfo("macOS: 使用 'launchd' 或系统设置中的登录项")
	} else {
		logInfo("已禁用开机自启")
	}
	
	// 保存配置到文件
	saveConfig()
}

// ensureFirewallRule 确保防火墙规则存在 (Unix/Linux/macOS)
func ensureFirewallRule() {
	cfgMu.RLock()
	autoFirewall := cfg.AutoFirewall
	cfgMu.RUnlock()
	
	if !autoFirewall {
		return
	}
	
	logInfo("防火墙配置需要手动操作（Unix/Linux/macOS）")
	logInfo("Linux: 使用 'ufw' 或 'iptables'")
	logInfo("macOS: 在系统偏好设置 → 安全性与隐私 → 防火墙")
}

