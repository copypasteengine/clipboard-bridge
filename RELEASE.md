# 📦 发布指南

本文档说明如何发布新版本到 GitHub Release。

## 🚀 快速发布流程

### 1. 准备发布

确保所有改动已经提交：

```bash
git status
git add .
git commit -m "描述你的改动"
git push origin main
```

### 2. 创建版本标签

```bash
# 创建标签（例如 v1.0.0）
git tag v1.0.0

# 推送标签到 GitHub
git push origin v1.0.0
```

### 3. 自动构建

推送标签后，GitHub Actions 会自动：
- ✅ 编译 Windows 可执行文件
- ✅ 创建配置文件示例
- ✅ 打包成 ZIP 文件
- ✅ 创建 GitHub Release
- ✅ 上传构建产物

### 4. 完成发布

访问你的仓库 Release 页面，编辑自动创建的 Release：
- 添加更新日志
- 标记重要改动
- 如果是重大更新，取消 `prerelease` 标记

## 📝 版本号规范

遵循语义化版本 (Semantic Versioning)：

- `v1.0.0` - 主版本号.次版本号.修订号
- `v1.0.0-beta.1` - Beta 测试版
- `v1.0.0-rc.1` - Release Candidate 候选版本

**示例：**
- `v1.0.0` → `v1.0.1` - 修复 Bug
- `v1.0.0` → `v1.1.0` - 新增功能
- `v1.0.0` → `v2.0.0` - 破坏性更新

## 🔧 本地构建测试

### 单平台构建

**Windows:**
```powershell
# 清理旧文件
del ClipboardBridge.exe

# 构建
go build -ldflags="-H windowsgui" -o ClipboardBridge.exe

# 测试运行
./ClipboardBridge.exe
```

**Linux/macOS:**
```bash
# 清理旧文件
rm -f clipboard-bridge

# 构建
go build -o clipboard-bridge

# 测试运行
./clipboard-bridge
```

### 多平台构建

**使用构建脚本:**

```bash
# Linux/macOS
chmod +x build.sh
./build.sh v1.0.0

# Windows
build.bat v1.0.0
```

构建脚本会在 `dist/` 目录生成所有平台的可执行文件：
- `ClipboardBridge-windows-amd64.exe`
- `clipboard-bridge-linux-amd64`
- `clipboard-bridge-linux-arm64`
- `clipboard-bridge-macos-amd64`
- `clipboard-bridge-macos-arm64`

**手动交叉编译:**

```bash
# Windows (在 Linux/macOS 上)
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
  go build -ldflags="-H windowsgui" -o ClipboardBridge.exe

# Linux ARM64 (在 x64 上)
GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc \
  go build -o clipboard-bridge

# macOS ARM64 (在 Intel Mac 上)
GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 \
  go build -o clipboard-bridge
```

> **注意**: 交叉编译需要对应平台的交叉编译工具链

## 📋 发布检查清单

在推送标签前，确认：

- [ ] 代码已在本地测试通过
- [ ] README.md 已更新
- [ ] 版本号符合规范
- [ ] 提交信息清晰明确
- [ ] 没有敏感信息（Token、密码等）

## 🐛 常见问题

### Q1: GitHub Actions 构建失败
**A:** 检查：
- Go 版本是否兼容
- 依赖是否正确
- 查看 Actions 日志获取详细错误

### Q2: 如何删除错误的标签
**A:** 
```bash
# 删除本地标签
git tag -d v1.0.0

# 删除远程标签
git push origin :refs/tags/v1.0.0
```

### Q3: 如何修改已发布的 Release
**A:** 
- 访问 GitHub Release 页面
- 点击 Release 右侧的编辑按钮
- 修改内容后保存

## 📊 项目文件说明

**提交到 Git 的文件:**
```
clipboard-bridge/
├── .github/
│   └── workflows/
│       └── build-release.yml    # 自动构建配置
├── .gitignore                   # Git 忽略规则
├── README.md                    # 项目文档
├── RELEASE.md                   # 发布指南（本文档）
├── config.example.json          # 配置示例
├── main_common.go               # 跨平台通用代码
├── clipboard_windows.go         # Windows 特定实现
├── clipboard_unix.go            # Unix/Linux/macOS 实现
├── icon.go                      # 图标资源
├── clipboard_windows.c          # C 剪贴板监听（仅 Windows）
├── build.sh                     # Linux/macOS 构建脚本
├── build.bat                    # Windows 构建脚本
├── go.mod                       # Go 依赖
└── go.sum                       # Go 依赖校验
```

**不提交的文件（已在 .gitignore）:**
```
- *.exe, clipboard-bridge       # 二进制文件（通过 Release 发布）
- dist/, build/                 # 构建输出目录
- config.json                   # 用户配置（可能含敏感信息）
- clipboard_bridge.log          # 运行日志
- *.log, *.tmp, *.bak          # 临时文件
```

## 🎯 推荐工作流

**日常开发:**
```bash
# 1. 开发新功能
git checkout -b feature/new-feature
# ... 编写代码 ...

# 2. 提交改动
git add .
git commit -m "feat: 添加新功能"

# 3. 合并到主分支
git checkout main
git merge feature/new-feature
git push origin main
```

**发布版本:**
```bash
# 1. 确保在主分支
git checkout main
git pull origin main

# 2. 创建并推送标签
git tag v1.1.0
git push origin v1.1.0

# 3. 等待 GitHub Actions 自动构建
# 4. 在 GitHub 上编辑 Release 说明
```

## 🔗 相关链接

- GitHub Actions 文档: https://docs.github.com/actions
- 语义化版本规范: https://semver.org/lang/zh-CN/
- Git 标签管理: https://git-scm.com/book/zh/v2/Git-基础-打标签

---

**祝发布顺利！** 🎉

