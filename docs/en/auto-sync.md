# 🔄 Auto-Sync Feature

Automatic clipboard synchronization - similar to Apple's Universal Clipboard.

## 🎯 What's Possible?

### ✅ Android → PC (Fully Automatic)

**Implementation:** Foreground Service with clipboard listener

**How it works:**
1. Enable "Auto Sync" in app
2. Service monitors clipboard changes
3. When you copy on Android, auto-sends to PC
4. Shows notification "✓ Auto-synced"

**User experience:**
```
Copy on Android → (auto-sync) → Paste on PC
```

**Similarity to Apple:** 90%

### ⚡ PC → Android (1-tap)

**Implementation:** Quick Settings Tile

**How it works:**
1. Copy on PC
2. Swipe down → Tap "Smart Sync"
3. Done!

**User experience:**
```
Copy on PC → Swipe → Tap → Paste on Android
```

**Only 1 step!**

## 💡 iOS Limitations

iOS system restrictions:
- ❌ No background clipboard monitoring
- ❌ Apps suspended when not in foreground
- ❌ No persistent background services

**Best option for iOS:**
- Siri voice: "Hey Siri, sync clipboard"
- Home screen shortcuts
- Automation triggers (WiFi, app open, etc.)

## 🆚 Comparison

| Feature | Apple Universal | This Project (Android) |
|---------|----------------|----------------------|
| Android → PC | - | ✅ Fully automatic |
| PC → Android | - | ⚡ 1-tap |
| iOS → Mac | ✅ Automatic | - |
| Mac → iOS | ✅ Automatic | - |
| Cross-platform | ❌ Apple only | ✅ Any OS combo |
| Requires account | ✅ iCloud | ❌ No account |

## 📱 Enabling Auto-Sync

**In Android App:**
1. Open app
2. Toggle "Auto Sync" switch to ON
3. Grant notification permission
4. Done!

**When enabled:**
- Service runs in background
- Notification shows status
- Auto-syncs when clipboard changes
- Can pause/resume anytime

## 🔋 Battery Impact

Minimal (<1% per day)
- Efficient clipboard listener
- No polling or timers
- Only sends on actual changes

---

**Back to:** [Android Guide](./android-guide.md) | [Documentation](../README.md)

