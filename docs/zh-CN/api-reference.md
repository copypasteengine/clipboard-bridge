# 🔌 API 参考

Clipboard Bridge 完整 HTTP API 文档。

## 基础 URL

```
http://你的电脑IP:5678
```

默认端口：`5678`

## 认证方式

### Token 头部（推荐）

```http
X-Auth-Token: your-token-here
```

### URL 参数

```http
?token=your-token-here
```

如果未配置 Token，则不需要认证。

## 接口列表

### GET /pull

从电脑获取剪贴板内容。

**请求：**
```http
GET /pull HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token
```

**响应：**
```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

你好世界
```

**示例（curl）：**
```bash
curl -H "X-Auth-Token: your-token" \
  http://192.168.1.100:5678/pull
```

### POST /push

设置电脑剪贴板内容。

**方式 1：表单数据**
```http
POST /push HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token
Content-Type: application/x-www-form-urlencoded

text=你好世界
```

**方式 2：原始内容**
```http
POST /push HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token

你好世界
```

**响应：**
```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

OK
```

**示例（curl）：**
```bash
# 表单数据
curl -X POST \
  -H "X-Auth-Token: your-token" \
  -d "text=你好世界" \
  http://192.168.1.100:5678/push

# 原始内容
curl -X POST \
  -H "X-Auth-Token: your-token" \
  --data-binary "你好世界" \
  http://192.168.1.100:5678/push
```

### GET /meta

获取剪贴板元数据（内容 + 时间戳）。

**请求：**
```http
GET /meta HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token
```

**响应：**
```json
{
  "text": "你好世界",
  "updated": 1733400000
}
```

- `text`：剪贴板内容
- `updated`：Unix 时间戳（秒）

**示例（curl）：**
```bash
curl -H "X-Auth-Token: your-token" \
  http://192.168.1.100:5678/meta
```

### GET /ping

健康检查接口。

**请求：**
```http
GET /ping HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token
```

**响应：**
```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

PONG
```

**示例（curl）：**
```bash
curl -H "X-Auth-Token: your-token" \
  http://192.168.1.100:5678/ping
```

## 错误响应

### 401 Unauthorized

Token 验证失败。

```
HTTP/1.1 401 Unauthorized
Content-Type: text/plain; charset=utf-8

Unauthorized
```

### 400 Bad Request

请求体无效。

```
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8

Failed to read request body
```

### 500 Internal Server Error

服务器错误（如剪贴板访问失败）。

```
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain; charset=utf-8

Failed to write clipboard: [错误详情]
```

## 限制

- **请求体大小**：最大 10 MB
- **超时设置**：
  - 读取：30 秒
  - 写入：30 秒
  - 空闲：120 秒

## 客户端示例

### Python

```python
import requests

BASE_URL = "http://192.168.1.100:5678"
TOKEN = "your-token"
headers = {"X-Auth-Token": TOKEN}

# 从电脑获取
text = requests.get(f"{BASE_URL}/pull", headers=headers).text

# 发送到电脑
requests.post(f"{BASE_URL}/push", data={"text": "你好"}, headers=headers)

# 获取元数据
meta = requests.get(f"{BASE_URL}/meta", headers=headers).json()
print(f"内容: {meta['text']}, 更新: {meta['updated']}")
```

### JavaScript (Node.js)

```javascript
const axios = require('axios');

const BASE_URL = 'http://192.168.1.100:5678';
const headers = { 'X-Auth-Token': 'your-token' };

// 从电脑获取
const text = await axios.get(`${BASE_URL}/pull`, { headers });

// 发送到电脑
await axios.post(`${BASE_URL}/push`, 
  new URLSearchParams({ text: '你好' }), 
  { headers }
);

// 获取元数据
const meta = await axios.get(`${BASE_URL}/meta`, { headers });
console.log(meta.data);
```

### PowerShell

```powershell
$baseUrl = "http://192.168.1.100:5678"
$headers = @{ "X-Auth-Token" = "your-token" }

# 从电脑获取
$text = Invoke-RestMethod -Uri "$baseUrl/pull" -Headers $headers

# 发送到电脑
$body = @{ text = "你好世界" }
Invoke-RestMethod -Uri "$baseUrl/push" -Method Post -Body $body -Headers $headers

# 获取元数据
$meta = Invoke-RestMethod -Uri "$baseUrl/meta" -Headers $headers
```

---

**返回：** [文档索引](../README.md)

