# 📋 Clipboard Bridge

[中文](./README.zh-CN.md) | **English**

**Lightweight text clipboard sync across devices.** Share text between your computer and phone via simple HTTP API.

> **📝 Plain Text Only** - For images/files, use [LocalSend](https://localsend.org/) or [KDE Connect](https://kdeconnect.kde.org/)

## 🎯 Who Should Use This?

**✅ Perfect For:**
- Android + Windows/Linux/macOS
- iPhone + Windows/Linux

**❌ Not Needed:**
- iPhone + Mac (use Apple's Universal Clipboard)

## ✨ Features

- 🌐 HTTP REST API
- 📱 Native Android App (Material Design 3)
- 🌍 Multi-language (EN/ZH/JA)
- 🔄 Smart Sync (auto-detect direction)
- ⚡ Quick Access (widgets, shortcuts, quick settings)
- 🔒 Optional token authentication
- 📊 System tray UI
- ⚡ Lightweight (CPU <0.1%, RAM ~15MB)

## 🖥️ Platforms

| Platform | Type | Features |
|----------|------|----------|
| Windows / Linux / macOS | Desktop Service | HTTP API, System tray, Auto-start |
| Android | Native App | Auto-sync, Widgets, Quick Settings |
| iOS | Shortcuts | Siri control, Automation |

## 🚀 Quick Start

### 1. Download

**Desktop:** [Releases Page](https://github.com/copypasteengine/clipboard-bridge/releases)
- Windows: `clipboard-bridge-windows-amd64.zip`
- Linux: `clipboard-bridge-linux-amd64.tar.gz`
- macOS: `clipboard-bridge-macos-arm64.tar.gz`

**Android:** `clipboard-bridge-android-*.apk`

### 2. Install & Run

**Desktop:** Extract and run (Windows: double-click .exe)  
**Android:** Install APK → Open → Configure server address

### 3. Start Syncing!

Copy text on one device → Sync → Paste on another device ✓

**📖 Detailed Guide:** [Quick Start](./docs/en/quick-start.md)

## 📚 Documentation

**User Guides:**
- 🚀 [Quick Start](./docs/en/quick-start.md) - 5-minute setup
- 📱 [Android Guide](./docs/en/android-guide.md) - App & quick access
- 📱 [iOS Guide](./docs/en/ios-guide.md) - Shortcuts setup
- ❓ [FAQ](./docs/en/faq.md) - Common questions

**Reference:**
- 🔌 [API Reference](./docs/en/api-reference.md) - HTTP API
- ⚙️ [Configuration](./docs/en/configuration.md) - Config options

**Developer:**
- 🏗️ [Architecture](./docs/zh-CN/architecture.md) - System design
- 🔨 [Building](./docs/zh-CN/BUILDING.md) - Compile from source

**📖 [All Documentation](./DOCS.md)** - Complete index

## 🌐 API Example

```bash
# Get clipboard from PC
curl http://192.168.1.100:5678/pull

# Send to PC
curl -X POST http://192.168.1.100:5678/push -d "text=Hello"
```

See [API Reference](./docs/en/api-reference.md) for details.

## 🔒 Security

- Use Token authentication for production
- LAN-only recommended
- No cloud/internet connection needed

## 📄 License

MIT License

## 🤝 Contributing

Issues and PRs welcome! See [CONTRIBUTING.md](./CONTRIBUTING.md)

---

**🌟 Star this project if you find it useful!**
