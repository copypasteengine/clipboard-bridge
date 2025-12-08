# ⚙️ 配置参考

Clipboard Bridge 完整配置选项说明。

## 📁 配置文件位置

文件：`config.json`（与可执行文件同目录）

首次运行自动创建，使用默认值。

## 📋 配置选项

### 完整示例

```json
{
  "port": 5678,
  "token": "my-secret-token-123",
  "auto_start": true,
  "auto_firewall": true,
  "log_level": "info"
}
```

### port（端口）

**类型：** 整数  
**默认值：** `5678`  
**范围：** `1024-65535`

HTTP 服务监听端口。

**示例：**
```json
{
  "port": 8888
}
```

**注意：** 修改端口后需重启服务。

### token（令牌）

**类型：** 字符串  
**默认值：** `""`（空，不启用认证）

API 访问令牌，用于安全保护。

**示例：**
```json
{
  "token": "my-secret-token-123"
}
```

**安全建议：**
- 使用强随机 Token
- 妥善保管 Token
- 定期更换

### auto_start（开机自启）

**类型：** 布尔值  
**默认值：** `true`

系统启动时自动运行服务。

**平台支持：**
- Windows：通过注册表自动配置
- Linux/macOS：需手动配置

**示例：**
```json
{
  "auto_start": false
}
```

### auto_firewall（自动防火墙）

**类型：** 布尔值  
**默认值：** `true`

自动配置防火墙规则。

**平台支持：**
- 仅 Windows（使用 `netsh`）
- Linux/macOS：忽略此选项

**示例：**
```json
{
  "auto_firewall": false
}
```

### log_level（日志级别）

**类型：** 字符串  
**默认值：** `"info"`  
**可选值：** `"debug"`, `"info"`, `"error"`

日志详细程度。

**级别说明：**
- `error`：仅错误
- `info`：重要操作 + 错误（推荐）
- `debug`：所有详细信息（故障排查）

**示例：**
```json
{
  "log_level": "debug"
}
```

## 🔧 平台特定说明

### Windows

- `auto_start`：注册表位置 `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
- `auto_firewall`：使用 `netsh advfirewall` 命令

### Linux

- `auto_start`：需手动添加到 `~/.config/autostart/`
- 防火墙：使用 `ufw` 或 `iptables` 配置

### macOS

- `auto_start`：系统设置 → 用户与群组 → 登录项
- 防火墙：系统设置 → 安全性与隐私 → 防火墙

## 📝 应用更改

编辑 `config.json` 后：

1. 保存文件
2. 重启服务：
   - 右键托盘图标 → 退出
   - 重新运行程序

更改会在重启后立即生效。

## 🔗 相关文档

- [快速开始](./quick-start.md) - 初次安装
- [常见问题](./faq.md) - 问题解答
- [API 参考](../en/api-reference.md) - API 文档

---

**返回：** [文档索引](../README.md)

