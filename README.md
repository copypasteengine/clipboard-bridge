# 📋 Clipboard Bridge - 剪贴板桥接服务

一个轻量级的跨平台剪贴板同步服务,通过 HTTP API 实现跨设备剪贴板共享。支持 Windows、Linux 和 macOS 系统,特别适合与移动设备（iOS/Android）之间同步剪贴板内容。

## 🖥️ 支持的平台

| 平台 | 架构 | 剪贴板监听 | 开机自启 | 防火墙配置 |
|------|------|------------|----------|------------|
| Windows | x64 | ✅ 系统级 | ✅ 自动 | ✅ 自动 |
| Linux | x64 | ⚡ 轮询 | 📝 手动 | 📝 手动 |
| macOS | Apple Silicon (M1/M2/M3) | ⚡ 轮询 | 📝 手动 | 📝 手动 |

> **macOS Intel 用户**: 可以使用 Rosetta 2 运行 ARM64 版本（性能几乎无差异），或[从源码编译](#从源码编译)

## ✨ 主要功能

- 🔄 **实时剪贴板监控** - 自动检测本地剪贴板变化
- 🌐 **HTTP API** - 提供 REST 风格的接口访问剪贴板
- 🔒 **Token 认证** - 可选的安全令牌验证
- 🚀 **开机自启** - 支持 Windows 开机自动启动
- 🛡️ **防火墙管理** - 自动配置 Windows 防火墙规则
- 📱 **系统托盘** - 友好的托盘图标和菜单管理
- 📝 **日志记录** - 完整的操作日志,支持多级别(debug/info/error)
- 🔐 **并发安全** - 多线程安全设计

## 📦 安装和使用

### 1. 下载和运行

**从 GitHub Release 下载:**

1. 访问 [Releases 页面](https://github.com/YOUR_USERNAME/clipboard-bridge/releases)
2. 根据你的系统下载对应的文件:
   - **Windows x64**: `ClipboardBridge-windows-amd64.zip`
   - **Linux x64**: `ClipboardBridge-linux-amd64.tar.gz`
   - **macOS (M1/M2/M3)**: `ClipboardBridge-macos-arm64.tar.gz`

> **注意**: macOS Intel 用户需要使用 Rosetta 2 运行 ARM64 版本，或[从源码编译](#从源码编译)

**Windows 安装:**
```powershell
# 1. 解压 zip 文件
# 2. 双击运行 ClipboardBridge.exe
# 3. 查看系统托盘图标
```

**Linux / macOS 安装:**
```bash
# 1. 解压文件
tar -xzf ClipboardBridge-*.tar.gz

# 2. 添加执行权限
chmod +x clipboard-bridge

# 3. 运行程序
./clipboard-bridge

# 4. (可选) 移动到系统路径
sudo mv clipboard-bridge /usr/local/bin/
```

> **注意**: 
> - Windows: 首次运行可能需要允许防火墙访问
> - Linux/macOS: 需要安装 X11/Wayland 剪贴板支持

### 2. 配置文件

程序首次运行会在同目录下创建 `config.json`,可以根据需要修改配置。

**示例配置:**

```json
{
  "port": 5678,
  "token": "",
  "auto_start": true,
  "auto_firewall": true,
  "log_level": "info"
}
```

**配置说明:**

| 配置项 | 说明 | 可选值 |
|--------|------|--------|
| `port` | 服务监听端口 | 1024-65535 |
| `token` | API 访问令牌,为空则不验证 | 任意字符串 |
| `auto_start` | 是否开机自启 | true/false |
| `auto_firewall` | 是否自动配置防火墙 | true/false |
| `log_level` | 日志级别 | debug/info/error |

### 3. 系统托盘菜单

右键点击托盘图标:
- 📡 **服务地址** - 显示当前访问地址
- 💻 **本机地址** - 本机测试地址
- 🚀 **开机自启** - 切换开机自启状态
- ▶️ **启动/停止服务** - 手动控制服务
- 📄 **打开日志文件** - 查看运行日志
- ❌ **退出** - 退出程序

## 🔌 API 接口

### 获取剪贴板内容 (Pull)

```http
GET /pull
X-Auth-Token: your-token
```

**响应:**
```
Hello World
```

### 设置剪贴板内容 (Push)

```http
POST /push
X-Auth-Token: your-token
Content-Type: application/x-www-form-urlencoded

text=Hello World
```

或直接发送请求体:

```http
POST /push
X-Auth-Token: your-token

Hello World
```

**响应:**
```
OK
```

### 获取剪贴板元数据 (Meta)

```http
GET /meta
X-Auth-Token: your-token
```

**响应:**
```json
{
  "text": "Hello World",
  "updated": 1701234567
}
```

### 健康检查 (Ping)

```http
GET /ping
X-Auth-Token: your-token
```

**响应:**
```
PONG
```

## 📱 iOS 快捷指令配置

### 功能说明

通过 iOS 快捷指令实现智能剪贴板同步:
- ✅ iOS 空 + Windows 有内容 → 自动从 Windows 同步
- ✅ Windows 空 + iOS 有内容 → 自动推送到 Windows
- ✅ 两边都空 → 提示无内容
- ✅ 两边都有且相同 → 提示已同步
- ✅ 两边都有但不同 → 弹菜单让你选择

### 配置前准备

需要准备的信息:

```
家里 Wi-Fi 名称: MyHomeWiFi
家里 Windows IP: 192.168.1.100

公司 Wi-Fi 名称: OfficeWiFi  
公司 Windows IP: 10.0.8.50

端口号: 5678
Token: (如果设置了就填,没设置就留空)
```

### 快速配置步骤

1. **获取网络信息**
   ```
   获取网络详细信息 (Wi-Fi)
   从输入获取值 → 字典值 → 键: SSID
   设定变量 → CurrentSSID
   ```

2. **判断网络环境**
   ```
   如果 CurrentSSID 等于 "MyHomeWiFi"
      文本 → 192.168.1.100
      设定变量 → WindowsIP
   否则如果 CurrentSSID 等于 "OfficeWiFi"
      文本 → 10.0.8.50
      设定变量 → WindowsIP
   否则
      显示通知 → "当前 Wi-Fi 未配置,无法同步"
      停止快捷指令
   ```

3. **设置连接参数**
   ```
   文本 → 5678
   设定变量 → Port
   
   文本 → your-token
   设定变量 → Token
   ```

4. **获取 Windows 剪贴板**
   ```
   文本 → http://[WindowsIP]:[Port]/meta
   获取 URL 内容
      方法: GET
      标头: X-Auth-Token = [Token]
   
   从输入获取值 → 字典值 → 键: text
   设定变量 → WinText
   ```

5. **获取 iOS 剪贴板**
   ```
   获取剪贴板
   设定变量 → IOSText
   ```

6. **智能同步逻辑**
   ```
   如果 WinText 等于 IOSText
      显示通知 → "✓ 剪贴板已同步"
      停止快捷指令
   
   如果 IOSText 为空
      如果 WinText 为空
         显示通知 → "两边都是空的"
         停止快捷指令
      否则
         设定剪贴板 → WinText
         显示通知 → "⬇️ 已从 Windows 同步"
         停止快捷指令
   
   如果 WinText 为空
      文本 → http://[WindowsIP]:[Port]/push
      获取 URL 内容
         方法: POST
         请求体: 表单
         字段: text = [IOSText]
         标头: X-Auth-Token = [Token]
      显示通知 → "⬆️ 已上传到 Windows"
      停止快捷指令
   
   从菜单中选择
      提示: "剪贴板内容不同,请选择:"
      选项1: 使用 Windows 内容
         设定剪贴板 → WinText
         显示通知 → "⬇️ 已使用 Windows 内容"
      选项2: 使用 iOS 内容
         (发送 POST 请求到 /push)
         显示通知 → "⬆️ 已使用 iOS 内容"
      选项3: 取消
         显示通知 → "已取消同步"
   ```

### 使用方法

**方式 1: 直接运行**
- 打开"快捷指令" App
- 点击"剪贴板同步"

**方式 2: Siri 语音**
- "嘿 Siri,运行剪贴板同步"

**方式 3: 主屏幕小组件**
- 长按主屏幕 → 添加小组件 → 选择快捷指令

**方式 4: 自动化**
- 打开"快捷指令" → "自动化"
- 创建个人自动化 → 连接到 Wi-Fi
- 添加操作: 运行快捷指令 → 剪贴板同步
- 关闭"运行前询问"

## 🛠️ 从源码编译

如果你想自己编译程序:

### 环境要求

- Go 1.20 或更高版本
- GCC 编译器（Windows 需要 MinGW-w64）

### 编译步骤

```bash
# 1. 克隆仓库
git clone https://github.com/YOUR_USERNAME/clipboard-bridge.git
cd clipboard-bridge

# 2. 安装依赖
go mod download

# 3. 编译

# Windows (无窗口模式)
go build -ldflags="-H windowsgui" -o ClipboardBridge.exe

# Linux
go build -o clipboard-bridge

# macOS
go build -o clipboard-bridge

# 交叉编译 (在 Linux/macOS 上编译 Windows 版本)
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags="-H windowsgui" -o ClipboardBridge.exe

# 4. 运行
# Windows: ./ClipboardBridge.exe
# Linux/macOS: ./clipboard-bridge
```

### 自动构建

本项目配置了 GitHub Actions 自动构建:

1. 创建新的 Git 标签: `git tag v1.0.0`
2. 推送标签: `git push origin v1.0.0`
3. GitHub Actions 会自动构建并创建 Release
4. 可以在 Releases 页面下载构建好的二进制文件

## 🔧 平台特性说明

### Windows
- ✅ **系统级剪贴板监听**: 使用 Windows API 实时监听剪贴板变化
- ✅ **自动开机自启**: 通过注册表自动配置
- ✅ **防火墙自动配置**: 自动添加防火墙规则（可能需要管理员权限）
- ✅ **系统托盘**: 友好的托盘图标和菜单

### Linux
- ⚡ **轮询监听**: 每秒检查一次剪贴板变化
- 📝 **手动开机自启**: 需要添加到 `~/.config/autostart/` 或使用 systemd
- 📝 **手动防火墙**: 使用 `ufw` 或 `iptables` 配置
- ✅ **系统托盘**: 支持（需要桌面环境）
- 📦 **依赖**: 需要 `xclip` 或 `xsel`（X11）或 `wl-clipboard`（Wayland）

### macOS
- ⚡ **轮询监听**: 每秒检查一次剪贴板变化
- 📝 **手动开机自启**: 通过系统偏好设置 → 用户与群组 → 登录项
- 📝 **手动防火墙**: 通过系统偏好设置 → 安全性与隐私 → 防火墙
- ✅ **系统托盘**: 支持菜单栏图标

## 🔒 安全说明

1. **Token 认证**: 建议设置 Token 以防止未授权访问
2. **局域网访问**: 服务默认监听所有网络接口,建议仅在可信网络使用
3. **防火墙**: 
   - Windows: 程序会尝试自动添加防火墙规则
   - Linux: 手动运行 `sudo ufw allow 5678/tcp`
   - macOS: 在系统设置中允许入站连接
4. **HTTPS**: 当前版本使用 HTTP,如需加密传输建议配置反向代理

## 📝 日志文件

日志文件位置: `clipboard_bridge.log` (程序同目录)

**日志级别说明:**
- `error`: 仅记录错误信息
- `info`: 记录关键操作和错误 (默认)
- `debug`: 记录所有详细信息,包括请求内容

**示例日志:**
```
[2024-12-05 10:30:15] [INFO] 程序启动,日志文件: D:\clipboard\clipboard_bridge.log
[2024-12-05 10:30:15] [INFO] 配置加载完成: Port=5678, Token=ab****yz, AutoStart=true
[2024-12-05 10:30:15] [INFO] 剪贴板监听已启动
[2024-12-05 10:30:15] [INFO] 🚀 剪贴板服务已启动
[2024-12-05 10:30:15] [INFO]    外部访问: http://192.168.1.100:5678
[2024-12-05 10:30:15] [INFO]    本机访问: http://localhost:5678
[2024-12-05 10:31:20] [INFO] 收到 Push 请求 (来自 192.168.1.200:54321)，内容长度: 15 字节
[2024-12-05 10:31:20] [INFO] ✓ 成功写入剪贴板，内容长度: 15 字节
```

## 🐛 常见问题

### Q1: 无法连接到服务
**A:** 检查以下几点:
- 服务是否已启动 (查看托盘图标)
- 端口是否被占用
- 防火墙是否允许该端口
- IP 地址是否正确

### Q2: Token 验证失败
**A:** 确保:
- 请求头 `X-Auth-Token` 或 URL 参数 `token` 与配置文件中的一致
- Token 区分大小写

### Q3: iOS 快捷指令无法同步
**A:** 检查:
- 手机和电脑在同一局域网
- Wi-Fi 名称配置正确
- Windows 服务正常运行
- 查看 Windows 端日志文件

### Q4: 开机自启不生效
**A:** 
- 在托盘菜单中勾选"开机自启"
- 检查注册表: `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
- 确保程序路径正确

## 🔧 技术栈

- **语言**: Go 1.20
- **跨平台支持**: 条件编译 (build tags)
- **GUI**: getlantern/systray (系统托盘)
- **剪贴板**: atotto/clipboard (跨平台)
- **Windows API**: CGo + C 语言 (clipboard_windows.c)
- **HTTP**: 标准库 net/http
- **日志**: 标准库 log

## 📦 依赖说明

**所有平台:**
```bash
go get github.com/getlantern/systray
go get github.com/atotto/clipboard
```

**Linux 额外依赖:**
```bash
# Ubuntu/Debian
sudo apt-get install xclip  # 或 xsel (X11)
sudo apt-get install wl-clipboard  # (Wayland)

# Fedora/RHEL
sudo dnf install xclip
sudo dnf install wl-clipboard

# Arch Linux
sudo pacman -S xclip
sudo pacman -S wl-clipboard
```

**macOS:**
无需额外依赖,系统自带剪贴板支持

## 📄 许可证

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📧 联系方式

如有问题或建议,请通过 Issue 联系。

---

**享受跨设备剪贴板同步的便利！** 🎉

