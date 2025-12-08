# 📋 Clipboard Bridge

[English](./README.md) | **中文**

**轻量级跨设备文本剪贴板同步工具。** 通过简单的 HTTP API 在电脑和手机之间共享文本。

> **📝 仅支持纯文本** - 如需传输图片/文件，请使用 [LocalSend](https://localsend.org/) 或 [KDE Connect](https://kdeconnect.kde.org/)

## 🎯 适合谁使用？

**✅ 完美适配：**
- Android + Windows/Linux/macOS
- iPhone + Windows/Linux

**❌ 不需要：**
- iPhone + Mac（请使用 Apple 通用剪贴板）

## ✨ 核心特性

- 🌐 HTTP REST API
- 📱 Android 原生 App（Material Design 3）
- 🌍 多语言（EN/ZH/JA）
- 🔄 智能同步（自动判断方向）
- ⚡ 快捷访问（小部件、快速设置、长按菜单）
- 🔒 可选 Token 认证
- 📊 系统托盘界面
- ⚡ 轻量高效（CPU <0.1%，内存 ~15MB）

## 🖥️ 支持平台

| 平台 | 类型 | 功能 |
|------|------|------|
| Windows / Linux / macOS | 桌面服务 | HTTP API、系统托盘、开机自启 |
| Android | 原生 App | 自动同步、小部件、快速设置 |
| iOS | 快捷指令 | Siri 控制、自动化 |

## 🚀 快速开始

### 1. 下载

**桌面：** [Releases 页面](https://github.com/copypasteengine/clipboard-bridge/releases)
- Windows: `clipboard-bridge-windows-amd64.zip`
- Linux: `clipboard-bridge-linux-amd64.tar.gz`
- macOS: `clipboard-bridge-macos-arm64.tar.gz`

**Android:** `clipboard-bridge-android-*.apk`

### 2. 安装运行

**桌面：** 解压并运行（Windows 双击 .exe）  
**Android:** 安装 APK → 打开 → 配置服务器地址

### 3. 开始同步！

在一个设备复制文本 → 同步 → 在另一个设备粘贴 ✓

**📖 详细步骤：** [快速开始指南](./docs/zh-CN/quick-start.md)

## 📚 文档

**用户指南：**
- 🚀 [快速开始](./docs/zh-CN/quick-start.md) - 5 分钟上手
- 📱 [Android 指南](./docs/zh-CN/android-guide.md) - App 与快捷访问
- 📱 [iOS 指南](./docs/zh-CN/ios-guide.md) - 快捷指令配置
- ❓ [常见问题](./docs/zh-CN/faq.md) - 问题解答

**参考文档：**
- 🔌 [API 参考](./docs/en/api-reference.md) - HTTP API
- ⚙️ [配置说明](./docs/zh-CN/configuration.md) - 配置选项

**开发者：**
- 🏗️ [架构设计](./docs/zh-CN/architecture.md) - 系统架构
- 🔨 [编译指南](./docs/zh-CN/BUILDING.md) - 从源码构建

**📖 [完整文档](./DOCS.md)** - 所有文档索引

## 🌐 API 示例

```bash
# 从电脑获取剪贴板
curl http://192.168.1.100:5678/pull

# 发送到电脑
curl -X POST http://192.168.1.100:5678/push -d "text=你好"
```

详见 [API 参考](./docs/en/api-reference.md)

## 🔒 安全建议

- 生产环境使用 Token 认证
- 建议仅在局域网使用
- 无需云端/互联网连接

## 📄 许可证

MIT License

## 🤝 参与贡献

欢迎提交 Issue 和 PR！详见 [CONTRIBUTING.md](./CONTRIBUTING.md)

---

**🌟 觉得有用就给个 Star 吧！**
