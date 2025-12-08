# 📱 iOS Shortcuts Guide

Complete guide for using Clipboard Bridge with iOS Shortcuts.

> **Note:** If you use iPhone + Mac, use Apple's Universal Clipboard instead.

## 📥 Basic Setup

### Pull from PC

1. Open "Shortcuts" app → Tap "+"
2. Add action: "Get Contents of URL"
   - URL: `http://192.168.1.100:5678/pull` (replace with your PC IP)
   - Method: GET
   - If token set: Add header `X-Auth-Token` = `your-token`
3. Add action: "Set Clipboard"
4. Add action: "Show Notification" → "✓ Synced from PC"
5. Save shortcut

### Push to PC

1. Open "Shortcuts" app → Tap "+"
2. Add action: "Get Clipboard"
3. Add action: "Get Contents of URL"
   - URL: `http://192.168.1.100:5678/push`
   - Method: POST
   - Request Body: Form
   - Add field: `text` = [Clipboard]
   - If token set: Add header `X-Auth-Token` = `your-token`
4. Add action: "Show Notification" → "✓ Sent to PC"
5. Save shortcut

## 🎯 Smart Sync (Recommended)

Auto-detect sync direction:

1. Get PC clipboard via `/meta`
2. Get iPhone clipboard
3. Compare and decide:
   - Same → Show "Already synced"
   - iPhone empty → Pull from PC
   - PC empty → Push to PC
   - Different → Show menu to choose

See Chinese docs for detailed step-by-step.

## ⚡ Quick Access

### Home Screen Widget

1. Long-press home screen
2. Tap "+"
3. Search "Shortcuts"
4. Add widget
5. Select your clipboard sync shortcut

### Siri Voice Control

Say: "Hey Siri, run clipboard sync"

### Automation

Create automation in Shortcuts app:

**Trigger:** Connect to WiFi
- Select your home/office WiFi
- Action: Run clipboard sync shortcut
- Disable "Ask Before Running"

**Trigger:** Open specific app
- E.g., Notes, Safari
- Action: Pull from PC
- Auto-runs when you open the app

## 🐛 Troubleshooting

See [FAQ](./faq.md#ios-specific) for common iOS issues.

---

**Back to:** [Documentation](../README.md)

