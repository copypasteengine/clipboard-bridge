# 📱 Android User Guide

Complete guide for using Clipboard Bridge on Android.

## 📥 Installation

1. Download `clipboard-bridge-android-*.apk` from [Releases](https://github.com/copypasteengine/clipboard-bridge/releases)
2. Enable "Install from unknown sources" in Settings
3. Install the APK
4. Open app

## ⚙️ Initial Setup

1. Tap ⚙️ (Settings) icon in top-right
2. Enter PC address: `http://192.168.1.100:5678`
3. (Optional) Enter token if configured
4. Tap "Save"
5. Connection status should show "Connected" ✓

## 🎯 Basic Usage

### Manual Sync

**Pull from PC:**
1. Tap "Pull from PC" button
2. PC clipboard content copied to phone
3. Paste anywhere ✓

**Push to PC:**
1. Copy text on phone
2. Tap "Push to PC" button  
3. Content sent to PC ✓

**Smart Sync:**
1. Tap "Smart Sync" button
2. App auto-detects which direction to sync
3. Done ✓

### Auto-Sync (Optional)

Enable auto-sync to automatically send phone clipboard to PC:

1. In app, toggle "Auto Sync" switch to ON
2. Grant notification permission
3. Now when you copy on phone, it auto-syncs to PC!

**Note:** Shows notification "Auto-sync enabled"

## ⚡ Quick Access (Recommended!)

### Quick Settings Tile (Fastest)

**Setup:**
1. Swipe down notification bar twice
2. Tap pencil icon (Edit)
3. Find "Smart Sync" tile
4. Drag to quick settings area
5. Done!

**Usage:**
```
Copy → Swipe down → Tap "Smart Sync" → Done!
```

### Home Screen Widget

**Setup:**
1. Long-press home screen
2. Tap "Widgets"
3. Find "Clipboard Bridge"
4. Drag to home screen

**Widget has 3 buttons:**
- ⬇️ Pull from PC
- ⬆️ Push to PC
- 🔄 Smart Sync

### App Shortcuts

**Usage:**
1. Long-press app icon
2. Select action from menu
3. Done!

**Create standalone shortcut:**
- Long-press icon → Long-press menu item → Drag to home screen

See [Android Quick Access](./android-quick-access.md) for detailed guide.

## 📋 App Interface

```
┌──────────────────────────┐
│ Clipboard Bridge     ⚙️  │
├──────────────────────────┤
│ ● Connected          🔄  │
│ http://192.168.1.100:5678│
├──────────────────────────┤
│ 📱 Phone Clipboard   🔄  │
│ Hello from phone         │
├──────────────────────────┤
│ 💻 PC Clipboard      🔄  │
│ Hello from PC            │
│ Updated: 2 mins ago      │
├──────────────────────────┤
│  [ ⬇️  Pull from PC ]    │
│  [ ⬆️  Push to PC ]      │
│  [ 🔄  Smart Sync ]      │
├──────────────────────────┤
│ 🔄 Auto Sync       [ON]  │
│ Auto-sync to PC enabled  │
└──────────────────────────┘
```

## 🐛 Troubleshooting

**Cannot connect:**
- Ensure phone and PC on same WiFi
- Check PC firewall allows port 5678
- Verify IP address is correct
- Check desktop service is running

**Token error:**
- Ensure token matches in app and config.json
- Token is case-sensitive

**Auto-sync not working:**
- Check notification permission granted
- Disable battery optimization for app
- Ensure service is running (notification visible)

See [FAQ](./faq.md) for more issues.

## 🔗 Related Docs

- [Quick Start](./quick-start.md) - Initial setup
- [Quick Access](./android-quick-access.md) - Widgets & shortcuts
- [API Reference](./api-reference.md) - HTTP API
- [FAQ](./faq.md) - Common questions

---

**Back to:** [Documentation Index](../README.md)

