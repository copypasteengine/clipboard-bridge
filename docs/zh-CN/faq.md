# ❓ 常见问题

## 📱 基本问题

### 可以同步什么内容？

**✅ 支持：**
- 纯文本
- 网址链接
- 验证码
- 聊天消息和笔记
- 代码片段
- 任何文字内容

**❌ 不支持：**
- 图片或照片
- 文件或文档
- 富文本格式
- 二进制数据

如需传输图片/文件，请使用 [LocalSend](https://localsend.org/)

### 安全吗？

- 仅使用本地网络（无云端）
- 可选 Token 认证
- 开源透明
- 建议仅在可信网络使用

### 需要注册账号吗？

不需要！无需注册、无需账号、无需云服务。

## 🔌 连接问题

### 手机无法连接电脑

**检查清单：**
1. 手机和电脑在同一 WiFi
2. 电脑防火墙允许 5678 端口
3. 使用电脑的局域网 IP（如 192.168.1.100，不是 127.0.0.1）
4. 桌面服务正在运行（查看托盘图标）

**测试连接：**
在手机浏览器打开 `http://电脑IP:5678/ping`  
应该显示：`PONG`

### 如何找到电脑 IP？

**方法1：** 鼠标悬停在托盘图标上（提示显示 IP）

**方法2：** 右键托盘图标（第一行菜单显示地址）

**方法3：** 命令行
- Windows: `ipconfig`（查看 IPv4 地址）
- Linux/Mac: `ip addr` 或 `ifconfig`

### Token 认证失败

- 确保 Token 与 `config.json` 中完全一致
- Token 区分大小写
- 检查 HTTP 头是 `X-Auth-Token`（不是 `Authorization`）

## 📱 Android 专属

### App 总是停止

- 关闭应用的电池优化
- 授予通知权限
- 检查网络连接

### 小部件不工作

- 确保在 App 中已配置服务器
- 先在 App 内测试同步
- 重新添加小部件

### 快速设置磁贴不显示

- 需要 Android 7.0+
- 如需要可重新安装 App
- 编辑时在"所有磁贴"中查找

### 自动同步耗电吗？

自动同步耗电极少（<1%）。如果担心：
- 使用快速设置磁贴代替（需要时点击）
- 关闭自动同步，改用手动同步

## 📱 iOS 专属

### 快捷指令显示"未授权"

- 添加 HTTP 头 `X-Auth-Token` 及你的 Token
- 或添加 URL 参数 `?token=your-token`

### 快捷指令不工作

- 检查 URL 包含 `http://` 前缀
- 验证 IP 地址正确
- 确保桌面服务运行中
- 先在浏览器测试：`http://IP:5678/ping`

### 可以自动化吗？

可以！在快捷指令 App 创建自动化：
- 触发器：连接到 WiFi
- 操作：运行剪贴板同步快捷指令
- 关闭"运行前询问"

## 🖥️ 桌面服务

### 服务无法启动

- 检查 5678 端口是否已被占用
- 尝试在 `config.json` 中更改端口
- 查看日志：`clipboard_bridge.log`

### 防火墙阻止连接

**Windows:** 应自动配置。如果不行：
```cmd
netsh advfirewall firewall add rule name="ClipboardBridge" dir=in action=allow protocol=TCP localport=5678
```

**Linux:**
```bash
sudo ufw allow 5678/tcp
```

**macOS:**
系统设置 → 安全性与隐私 → 防火墙 → 防火墙选项 → 添加应用

### 如何更改端口？

编辑 `config.json`：
```json
{
  "port": 8888
}
```

重启服务。

### 开机自启不工作（Windows）

- 在托盘菜单启用：右键 → 勾选"开机自启"
- 检查注册表：`HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`

## 🔧 高级问题

### 可以使用 HTTPS 吗？

桌面服务使用 HTTP。如需 HTTPS：
- 使用反向代理（Nginx/Caddy）
- 配置 SSL 证书
- 更新客户端 URL 为 `https://`

### 可以从互联网访问吗？

**不建议！** 设计仅用于局域网。

如果必须：
- 设置强 Token
- 使用 VPN 或 SSH 隧道
- 考虑安全风险

### 如何在两个手机间同步？

不直接支持。变通方法：
1. 手机 A → 电脑（push）
2. 手机 B → 电脑（pull）

或使用专门的工具如 [Join](https://joaoapps.com/join/)

## 📊 性能问题

### 速度快吗？

- 桌面：<50ms 响应时间
- Android（局域网）：<500ms
- iOS：与 Android 类似

### 占用资源多吗？

- 桌面：CPU <0.1%，内存 ~15MB
- Android：内存 ~30MB
- 电池影响：极少（<1%）

## 🔗 仍需帮助？

- 查看 [架构设计](./architecture.md) 了解工作原理
- 查看 [API 参考](../en/api-reference.md) 了解集成方式
- [提交 Issue](https://github.com/copypasteengine/clipboard-bridge/issues)

---

**返回：** [文档索引](../README.md)

