# 🚀 Quick Start Guide

Get clipboard sync working between your devices in 5 minutes!

## Step 1: Install Desktop Service (3 minutes)

### Windows

1. Download `clipboard-bridge-windows-amd64.zip` from [Releases](https://github.com/copypasteengine/clipboard-bridge/releases)
2. Extract to any folder
3. Double-click `clipboard-bridge.exe`
4. Check system tray icon ✓

Service starts automatically on port 5678.

### Linux

```bash
wget https://github.com/copypasteengine/clipboard-bridge/releases/latest/download/clipboard-bridge-linux-amd64.tar.gz
tar -xzf clipboard-bridge-linux-amd64.tar.gz
sudo apt-get install xclip  # Install clipboard tool
chmod +x clipboard-bridge
./clipboard-bridge
```

### macOS

```bash
curl -LO https://github.com/copypasteengine/clipboard-bridge/releases/latest/download/clipboard-bridge-macos-arm64.tar.gz
tar -xzf clipboard-bridge-macos-arm64.tar.gz
chmod +x clipboard-bridge
./clipboard-bridge
```

## Step 2: Get Your PC's IP Address (1 minute)

**Method 1:** Hover over tray icon
```
Clipboard Bridge - 192.168.1.100:5678
```

**Method 2:** Right-click tray icon, first menu item shows:
```
📡 Service Address: http://192.168.1.100:5678
```

**Remember your IP!** e.g., `192.168.1.100`

## Step 3: Install Mobile Client (2 minutes)

### Android (Recommended)

1. Download `clipboard-bridge-android-*-debug.apk`
2. Install on device (enable "Install from unknown sources")
3. Open app → Tap ⚙️ → Enter `http://192.168.1.100:5678`
4. Save → Done!

**Quick Access Setup:**
- Swipe down notification bar → Edit → Add "Smart Sync" tile
- Long-press home screen → Widgets → Add "Clipboard Bridge"

### iOS

1. Open "Shortcuts" app → Tap "+"
2. Add "Get Contents of URL"
   - URL: `http://192.168.1.100:5678/pull`
   - Method: GET
3. Add "Set Clipboard"
4. Add "Show Notification": ✓ Synced

See [iOS Guide](./ios-guide.md) for details.

## ✅ Test Sync

**PC to Phone:**
1. Copy text on PC
2. Tap "Smart Sync" on phone (swipe down or widget)
3. Paste on phone
4. ✓ Success!

**Phone to PC:**
1. Copy text on phone
2. (Auto-syncs if enabled) or tap Quick Settings
3. Paste on PC
4. ✓ Success!

---

**Next:** [Android Quick Access](./android-quick-access.md) | [Full Documentation](../README.md)

