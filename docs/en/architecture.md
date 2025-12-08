# 🏗️ Architecture

System architecture and design of Clipboard Bridge.

## 📊 System Overview

```
┌─────────────────────────────────────────────┐
│         Mobile Clients                      │
├─────────────────────────────────────────────┤
│  📱 Android App    │  📱 iOS Shortcuts      │
│  - Material UI     │  - Smart Sync         │
│  - Native Kotlin   │  - Siri Integration   │
└────────┬─────────────────┬───────────────────┘
         │                 │
         │   HTTP API      │
         │                 │
┌────────┴─────────────────┴───────────────────┐
│         Desktop Service (HTTP Server)        │
├─────────────────────────────────────────────┤
│  🌐 HTTP API Server                         │
│  ├── GET  /pull  - Get clipboard            │
│  ├── POST /push  - Set clipboard            │
│  ├── GET  /meta  - Get metadata             │
│  └── GET  /ping  - Health check             │
│                                              │
│  🔒 Auth Layer (Optional)                   │
│  └── X-Auth-Token verification              │
│                                              │
│  📋 Clipboard Layer                         │
│  ├── Windows: System API (C/CGo)           │
│  └── Unix: Polling (Go)                    │
└─────────────────────────────────────────────┘
```

## 🔄 Data Flow

### Scenario: Pull from PC

```
Phone App → GET /pull → Desktop Service → System Clipboard
          ← "Hello" ←                   ← ReadAll()

Phone sets local clipboard ✓
```

### Scenario: Push to PC

```
Phone App → POST /push → Desktop Service → System Clipboard
    text="World"                        → WriteAll()
          ← OK ←
```

## 📂 Code Structure

### Desktop Service (Go)

- `main_common.go` - Cross-platform code
- `clipboard_windows.go` - Windows implementation
- `clipboard_unix.go` - Linux/macOS implementation
- `clipboard_windows.c` - Windows C code
- `i18n.go` - Internationalization

### Android App (Kotlin)

- MVVM architecture
- Jetpack Compose UI
- Repository pattern
- Coroutines for async operations

## 🔐 Security Design

- Token-based authentication
- Request size limits (10MB)
- Timeout controls
- LAN-only recommended

## ⚡ Performance

| Platform | CPU | Memory | Response |
|----------|-----|--------|----------|
| Desktop | <0.1% | ~15MB | <50ms |
| Android | - | ~30MB | <500ms |

---

**Back to:** [Documentation](../README.md)

