# ⚙️ Configuration Reference

Complete configuration options for Clipboard Bridge.

## 📁 Config File Location

File: `config.json` (same directory as executable)

Created automatically on first run with defaults.

## 📋 Configuration Options

### Complete Example

```json
{
  "port": 5678,
  "token": "my-secret-token-123",
  "auto_start": true,
  "auto_firewall": true,
  "log_level": "info"
}
```

### port

**Type:** Integer  
**Default:** `5678`  
**Range:** `1024-65535`

HTTP service port.

**Example:**
```json
{
  "port": 8888
}
```

**Note:** Restart service after changing port.

### token

**Type:** String  
**Default:** `""` (empty, no authentication)

API access token for security.

**Example:**
```json
{
  "token": "my-secret-token-123"
}
```

**Security Tips:**
- Use strong, random tokens
- Keep token secret
- Change periodically

### auto_start

**Type:** Boolean  
**Default:** `true`

Enable auto-start on system boot.

**Platforms:**
- Windows: Auto-configured via registry
- Linux/macOS: Manual setup required

**Example:**
```json
{
  "auto_start": false
}
```

### auto_firewall

**Type:** Boolean  
**Default:** `true`

Automatically configure firewall rules.

**Platforms:**
- Windows only (uses `netsh`)
- Linux/macOS: Ignored

**Example:**
```json
{
  "auto_firewall": false
}
```

### log_level

**Type:** String  
**Default:** `"info"`  
**Options:** `"debug"`, `"info"`, `"error"`

Logging verbosity level.

**Levels:**
- `error`: Errors only
- `info`: Important operations + errors (recommended)
- `debug`: All details (for troubleshooting)

**Example:**
```json
{
  "log_level": "debug"
}
```

## 🔧 Platform-Specific Notes

### Windows

- `auto_start`: Registry key at `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
- `auto_firewall`: Uses `netsh advfirewall` command

### Linux

- `auto_start`: Add to `~/.config/autostart/` manually
- Firewall: Configure with `ufw` or `iptables`

### macOS

- `auto_start`: System Settings → Users & Groups → Login Items
- Firewall: System Settings → Security & Privacy → Firewall

## 📝 Applying Changes

After editing `config.json`:

1. Save file
2. Restart service:
   - Right-click tray icon → Quit
   - Run executable again

Changes take effect immediately on restart.

## 🔗 Related

- [Quick Start](./quick-start.md) - Initial setup
- [FAQ](./faq.md) - Common questions
- [API Reference](./api-reference.md) - API docs

---

**Back to:** [Documentation](../README.md)

