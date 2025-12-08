# ❓ Frequently Asked Questions

## 📱 General

### What can I sync?

**✅ Supported:**
- Plain text
- URLs and links
- Verification codes
- Messages and notes
- Code snippets
- Any text content

**❌ Not Supported:**
- Images or photos
- Files or documents
- Rich text formatting
- Binary data

For image/file transfer, use [LocalSend](https://localsend.org/)

### Is it secure?

- Uses local network only (no cloud)
- Optional token authentication
- Open source (fully transparent)
- Recommend using on trusted networks only

### Do I need an account?

No! No registration, no account, no cloud service needed.

## 🔌 Connection Issues

### Phone cannot connect to PC

**Checklist:**
1. Phone and PC on same WiFi network
2. PC firewall allows port 5678
3. Using PC's LAN IP (e.g., 192.168.1.100, NOT 127.0.0.1)
4. Desktop service is running (check tray icon)

**Test connection:**
Open `http://YOUR_PC_IP:5678/ping` in phone browser.  
Should show: `PONG`

### How do I find my PC's IP?

**Method 1:** Hover over tray icon (shows IP in tooltip)

**Method 2:** Right-click tray icon (first menu item shows address)

**Method 3:** Command line
- Windows: `ipconfig` (look for IPv4 Address)
- Linux/Mac: `ip addr` or `ifconfig`

### Token authentication failed

- Ensure token matches exactly in app and `config.json`
- Token is case-sensitive
- Check HTTP header is `X-Auth-Token` (not `Authorization`)

## 📱 Android Specific

### App keeps stopping

- Disable battery optimization for the app
- Grant notification permission
- Check network connection

### Widget not working

- Ensure server is configured in app
- Test sync in app first
- Re-add widget

### Quick Settings tile not appearing

- Android 7.0+ required
- Reinstall app if needed
- Check in "All tiles" section when editing

### Auto-sync draining battery?

Auto-sync uses minimal battery (<1%). If concerned:
- Use Quick Settings tile instead (tap when needed)
- Disable auto-sync, use manual sync

## 📱 iOS Specific

### Shortcut shows "Unauthorized"

- Add HTTP header `X-Auth-Token` with your token
- Or add URL parameter `?token=your-token`

### Shortcut not working

- Check URL includes `http://` prefix
- Verify IP address is correct
- Ensure desktop service is running
- Test in phone browser first: `http://IP:5678/ping`

### Can I automate it?

Yes! Create automation in Shortcuts app:
- Trigger: Connect to WiFi
- Action: Run clipboard sync shortcut
- Disable "Ask Before Running"

## 🖥️ Desktop Service

### Service won't start

- Check if port 5678 is already in use
- Try changing port in `config.json`
- Check logs: `clipboard_bridge.log`

### Firewall blocking connections

**Windows:** Should auto-configure. If not:
```cmd
netsh advfirewall firewall add rule name="ClipboardBridge" dir=in action=allow protocol=TCP localport=5678
```

**Linux:**
```bash
sudo ufw allow 5678/tcp
```

**macOS:**
System Settings → Security & Privacy → Firewall → Firewall Options → Add app

### How to change port?

Edit `config.json`:
```json
{
  "port": 8888
}
```

Restart service.

### Auto-start not working (Windows)

- Enable in tray menu: Right-click → Check "Auto-Start"
- Check registry: `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`

## 🔧 Advanced

### Can I use HTTPS?

Desktop service uses HTTP. For HTTPS:
- Use reverse proxy (Nginx/Caddy)
- Configure SSL certificate
- Update client URLs to `https://`

### Can I access from internet?

**Not recommended!** Designed for LAN use only.

If you must:
- Set strong token
- Use VPN or SSH tunnel
- Consider security risks

### How to sync between two phones?

Not supported directly. Workaround:
1. Phone A → PC (push)
2. Phone B → PC (pull)

Or use dedicated app like [Join](https://joaoapps.com/join/)

## 📊 Performance

### Is it fast?

- Desktop: <50ms response time
- Android (LAN): <500ms
- iOS: Similar to Android

### Does it use much resources?

- Desktop: CPU <0.1%, RAM ~15MB
- Android: RAM ~30MB
- Battery impact: Minimal (<1%)

## 🔗 Still Need Help?

- Check [Architecture](./architecture.md) for how it works
- See [API Reference](./api-reference.md) for integration
- [Open an Issue](https://github.com/copypasteengine/clipboard-bridge/issues)

---

**Back to:** [Documentation](../README.md)

