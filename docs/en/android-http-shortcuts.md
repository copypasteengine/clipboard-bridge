# 📲 Android HTTP Shortcuts App Guide

Alternative Android solution using HTTP Shortcuts app (no native app needed).

## 📥 Installation

Install [HTTP Shortcuts](https://play.google.com/store/apps/details?id=ch.rmy.android.http_shortcuts) from Google Play.

## ⚙️ Setup

### Create "Pull from PC" Shortcut

1. Open HTTP Shortcuts
2. Tap "+" button
3. Select "Regular Shortcut"
4. Configure:
   - Name: Pull from PC
   - URL: `http://192.168.1.100:5678/pull`
   - Method: GET
   - Add Header: `X-Auth-Token` = `your-token` (if needed)
   - Response Handling → Success:
     - Add "Copy to Clipboard"
     - Add "Show Toast": ✓ Synced from PC

### Create "Push to PC" Shortcut

1. Tap "+" button
2. Select "Regular Shortcut"
3. Configure:
   - Name: Push to PC
   - URL: `http://192.168.1.100:5678/push`
   - Method: POST
   - Request Body: Form
     - Add parameter: `text` = `{clipboard}`
   - Add Header: `X-Auth-Token` = `your-token` (if needed)
   - Response Handling → Success:
     - Add "Show Toast": ✓ Sent to PC

## 🏠 Add to Home Screen

1. Long-press shortcut
2. Select "Place on Home Screen"
3. Adjust position

## 🆚 vs Native App

| Feature | Native App | HTTP Shortcuts |
|---------|-----------|----------------|
| Setup | ⭐⭐⭐ Easy | ⭐⭐ Moderate |
| UI | ⭐⭐⭐⭐⭐ Beautiful | ⭐⭐ Basic |
| Quick Access | Widgets, QS, Shortcuts | Home screen only |
| Auto-sync | ✅ Yes | ❌ No |
| Free | ✅ Yes | ✅ Yes |

**Recommendation:** Use native app for better experience.

---

**Back to:** [Android Guide](./android-guide.md)

