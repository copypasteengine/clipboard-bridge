# 🔨 Building from Source

How to compile Clipboard Bridge from source code.

## 🖥️ Desktop Service

### Requirements

- Go 1.20 or later
- GCC (MinGW-w64 for Windows)

### Build Steps

```bash
# Clone repository
git clone https://github.com/copypasteengine/clipboard-bridge.git
cd clipboard-bridge

# Install dependencies
go mod download

# Build for your platform

# Windows (no console window)
go build -ldflags="-H windowsgui" -o clipboard-bridge.exe

# Linux
go build -o clipboard-bridge

# macOS
go build -o clipboard-bridge
```

### Linux Dependencies

```bash
# Ubuntu/Debian
sudo apt-get install xclip libgtk-3-dev

# Fedora
sudo dnf install xclip gtk3-devel

# Arch
sudo pacman -S xclip gtk3
```

## 📱 Android App

### Requirements

- Android Studio or JDK 17+
- Android SDK 34

### With Android Studio

1. Open `android-app` directory in Android Studio
2. Wait for Gradle sync
3. Build → Make Project
4. Run on device or emulator

### Command Line

```bash
cd android-app

# Debug APK
./gradlew assembleDebug

# APK location
# app/build/outputs/apk/debug/app-debug.apk

# Install to device
adb install app/build/outputs/apk/debug/app-debug.apk
```

See [Chinese Building Guide](../zh-CN/BUILDING.md) for more details.

---

**Back to:** [Documentation](../README.md)

