# 🔄 代码重构说明 - 多平台支持

## 📝 重构概述

为了支持多操作系统（Windows、Linux、macOS），项目代码已重构为以下结构：

### 📂 新的文件结构

| 文件 | 说明 | 平台 |
|------|------|------|
| `main_common.go` | 跨平台通用代码（HTTP 服务、配置、日志等） | 全部 |
| `clipboard_windows.go` | Windows 特定实现 | Windows |
| `clipboard_unix.go` | Unix/Linux/macOS 实现 | Linux/macOS |
| `clipboard.c` | Windows 剪贴板 C 代码 | Windows |
| `icon.go` | 托盘图标资源 | 全部 |
| `build.sh` | Linux/macOS 构建脚本 | Linux/macOS |
| `build.bat` | Windows 构建脚本 | Windows |

### 🔧 技术实现

#### 条件编译

使用 Go 的 build tags 实现平台特定代码：

```go
//go:build windows
// +build windows
```

```go
//go:build !windows
// +build !windows
```

#### 平台差异

| 功能 | Windows | Linux/macOS |
|------|---------|-------------|
| 剪贴板监听 | 系统 API (C) | 轮询 (Go) |
| 监听频率 | 实时 | 1秒/次 |
| 开机自启 | 注册表自动 | 需要手动配置 |
| 防火墙 | netsh 自动 | 需要手动配置 |
| 日志查看器 | notepad.exe | xdg-open/open |

## 🎯 重构目标

1. ✅ **跨平台支持**: Windows、Linux、macOS
2. ✅ **代码分离**: 平台特定代码独立文件
3. ✅ **保持兼容**: API 和配置保持不变
4. ✅ **自动构建**: GitHub Actions 多平台构建
5. ✅ **易于维护**: 清晰的代码结构

## 📋 API 兼容性

**完全兼容！** 所有 HTTP API 接口保持不变：
- `GET /pull` - 获取剪贴板
- `POST /push` - 设置剪贴板
- `GET /meta` - 获取元数据
- `GET /ping` - 健康检查

## 🔄 迁移步骤

### 对于开发者

1. **更新代码库**
   ```bash
   git pull origin main
   ```

2. **重新编译**
   ```bash
   # Windows
   go build -ldflags="-H windowsgui" -o ClipboardBridge.exe
   
   # Linux/macOS
   go build -o clipboard-bridge
   ```

3. **测试运行**
   ```bash
   # Windows
   ./ClipboardBridge.exe
   
   # Linux/macOS
   ./clipboard-bridge
   ```

### 对于用户

**无需任何操作！**

- 配置文件 `config.json` 完全兼容
- API 调用方式不变
- iOS 快捷指令无需修改

## 🆕 新增功能

### 1. 多平台构建脚本

**Linux/macOS:**
```bash
chmod +x build.sh
./build.sh v1.0.0
```

**Windows:**
```batch
build.bat v1.0.0
```

### 2. GitHub Actions 多平台自动构建

推送 tag 后自动构建：
- Windows (x64)
- Linux (x64, ARM64)
- macOS (Intel, Apple Silicon)

### 3. 平台特定优化

- **Windows**: 保持原有的高性能实时监听
- **Unix**: 使用轮询方式，资源占用低

## 🐛 已知限制

### Linux/macOS

1. **剪贴板监听**: 使用轮询方式，有 1 秒延迟
2. **开机自启**: 需要手动配置
   - Linux: `~/.config/autostart/`
   - macOS: 系统偏好设置 → 用户与群组 → 登录项
3. **防火墙**: 需要手动配置
   - Linux: `sudo ufw allow 5678/tcp`
   - macOS: 系统偏好设置 → 防火墙

### 依赖要求

**Linux:**
```bash
# X11
sudo apt-get install xclip  # 或 xsel

# Wayland
sudo apt-get install wl-clipboard
```

**macOS:**
无额外依赖

## 📊 性能对比

| 指标 | Windows | Linux/macOS |
|------|---------|-------------|
| CPU 使用 | ~0% (事件驱动) | ~0.1% (1秒轮询) |
| 内存占用 | ~15MB | ~12MB |
| 响应延迟 | <10ms | <1秒 |
| 电池影响 | 极低 | 极低 |

## 🔮 未来规划

1. **改进 Linux/macOS 监听**
   - 探索更高效的监听机制
   - 考虑使用 inotify 或 FSEvents

2. **GUI 配置界面**
   - 跨平台图形配置工具
   - 简化用户配置流程

3. **更多平台**
   - FreeBSD 支持
   - Android/iOS 原生客户端

## 🤝 贡献指南

欢迎提交 PR 改进跨平台支持！

**建议优化方向:**
1. Linux/macOS 更高效的剪贴板监听
2. 自动开机自启配置
3. 系统通知支持
4. 多语言界面

## 📞 问题反馈

遇到平台相关问题？请在 GitHub Issues 中反馈，并注明：
- 操作系统和版本
- 错误日志
- 复现步骤

---

**感谢你的支持！** 🎉

