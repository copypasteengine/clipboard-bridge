# 🔌 API Reference

Complete HTTP API documentation for Clipboard Bridge.

## Base URL

```
http://YOUR_PC_IP:5678
```

Default port: `5678`

## Authentication

### Token Header (Recommended)

```http
X-Auth-Token: your-token-here
```

### URL Parameter

```http
?token=your-token-here
```

If no token is configured, authentication is disabled.

## Endpoints

### GET /pull

Get clipboard content from PC.

**Request:**
```http
GET /pull HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token
```

**Response:**
```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

Hello World
```

**Example (curl):**
```bash
curl -H "X-Auth-Token: your-token" \
  http://192.168.1.100:5678/pull
```

### POST /push

Set clipboard content on PC.

**Method 1: Form Data**
```http
POST /push HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token
Content-Type: application/x-www-form-urlencoded

text=Hello+World
```

**Method 2: Raw Body**
```http
POST /push HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token

Hello World
```

**Response:**
```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

OK
```

**Example (curl):**
```bash
# Form data
curl -X POST \
  -H "X-Auth-Token: your-token" \
  -d "text=Hello World" \
  http://192.168.1.100:5678/push

# Raw body
curl -X POST \
  -H "X-Auth-Token: your-token" \
  --data-binary "Hello World" \
  http://192.168.1.100:5678/push
```

### GET /meta

Get clipboard metadata (content + timestamp).

**Request:**
```http
GET /meta HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token
```

**Response:**
```json
{
  "text": "Hello World",
  "updated": 1733400000
}
```

- `text`: Clipboard content
- `updated`: Unix timestamp (seconds)

**Example (curl):**
```bash
curl -H "X-Auth-Token: your-token" \
  http://192.168.1.100:5678/meta
```

### GET /ping

Health check endpoint.

**Request:**
```http
GET /ping HTTP/1.1
Host: 192.168.1.100:5678
X-Auth-Token: your-token
```

**Response:**
```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

PONG
```

**Example (curl):**
```bash
curl -H "X-Auth-Token: your-token" \
  http://192.168.1.100:5678/ping
```

## Error Responses

### 401 Unauthorized

Token verification failed.

```
HTTP/1.1 401 Unauthorized
Content-Type: text/plain; charset=utf-8

Unauthorized
```

### 400 Bad Request

Invalid request body.

```
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8

Failed to read request body
```

### 500 Internal Server Error

Server error (e.g., clipboard access failed).

```
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain; charset=utf-8

Failed to write clipboard: [error details]
```

## Limits

- **Request Body Size**: 10 MB maximum
- **Timeouts**: 
  - Read: 30 seconds
  - Write: 30 seconds
  - Idle: 120 seconds

## Client Examples

### Python

```python
import requests

BASE_URL = "http://192.168.1.100:5678"
TOKEN = "your-token"
headers = {"X-Auth-Token": TOKEN}

# Pull from PC
text = requests.get(f"{BASE_URL}/pull", headers=headers).text

# Push to PC
requests.post(f"{BASE_URL}/push", data={"text": "Hello"}, headers=headers)

# Get metadata
meta = requests.get(f"{BASE_URL}/meta", headers=headers).json()
print(f"Text: {meta['text']}, Updated: {meta['updated']}")
```

### JavaScript (Node.js)

```javascript
const axios = require('axios');

const BASE_URL = 'http://192.168.1.100:5678';
const headers = { 'X-Auth-Token': 'your-token' };

// Pull from PC
const text = await axios.get(`${BASE_URL}/pull`, { headers });

// Push to PC
await axios.post(`${BASE_URL}/push`, 
  new URLSearchParams({ text: 'Hello' }), 
  { headers }
);

// Get metadata
const meta = await axios.get(`${BASE_URL}/meta`, { headers });
console.log(meta.data);
```

### PowerShell

```powershell
$baseUrl = "http://192.168.1.100:5678"
$headers = @{ "X-Auth-Token" = "your-token" }

# Pull from PC
$text = Invoke-RestMethod -Uri "$baseUrl/pull" -Headers $headers

# Push to PC
$body = @{ text = "Hello World" }
Invoke-RestMethod -Uri "$baseUrl/push" -Method Post -Body $body -Headers $headers

# Get metadata
$meta = Invoke-RestMethod -Uri "$baseUrl/meta" -Headers $headers
```

---

**Back to:** [Documentation Index](../README.md)

