# 📱 Android 应用构建指南

## 🚀 快速开始

### 方式 1: 使用 Android Studio（推荐）

1. **安装 Android Studio**
   - 下载：https://developer.android.com/studio
   - 安装 Android SDK 和构建工具

2. **打开项目**
   ```bash
   # 克隆仓库
   git clone https://github.com/copypasteengine/clipboard-bridge.git
   cd clipboard-bridge/android-app
   ```
   
   - 启动 Android Studio
   - File → Open → 选择 `android-app` 目录
   - 等待 Gradle 同步完成

3. **运行应用**
   - 连接 Android 设备（开启 USB 调试）或启动模拟器
   - 点击 Run 按钮（绿色三角形）或按 `Shift+F10`
   - 应用会自动安装并启动

### 方式 2: 命令行构建

**前提条件：**
- 安装 JDK 11 或更高版本
- 设置 ANDROID_HOME 环境变量

**构建步骤：**

```bash
cd android-app

# Windows
gradlew.bat assembleDebug

# Linux/macOS
chmod +x gradlew
./gradlew assembleDebug
```

**输出位置：**
```
app/build/outputs/apk/debug/app-debug.apk
```

**安装到设备：**
```bash
adb install app/build/outputs/apk/debug/app-debug.apk
```

## 📦 构建 Release 版本

### 1. 生成签名密钥

```bash
keytool -genkey -v -keystore my-release-key.jks \
  -alias clipboard-bridge \
  -keyalg RSA \
  -keysize 2048 \
  -validity 10000

# 会提示输入密码和证书信息
```

### 2. 配置签名

创建 `android-app/keystore.properties`：

```properties
storePassword=你的密钥库密码
keyPassword=你的密钥密码
keyAlias=clipboard-bridge
storeFile=../my-release-key.jks
```

在 `app/build.gradle.kts` 中添加：

```kotlin
// 在 android 块之前
val keystorePropertiesFile = rootProject.file("keystore.properties")
val keystoreProperties = Properties()
if (keystorePropertiesFile.exists()) {
    keystoreProperties.load(FileInputStream(keystorePropertiesFile))
}

android {
    signingConfigs {
        create("release") {
            keyAlias = keystoreProperties["keyAlias"] as String
            keyPassword = keystoreProperties["keyPassword"] as String
            storeFile = file(keystoreProperties["storeFile"] as String)
            storePassword = keystoreProperties["storePassword"] as String
        }
    }

    buildTypes {
        release {
            signingConfig = signingConfigs.getByName("release")
            // ... 其他配置
        }
    }
}
```

### 3. 构建 Release APK

```bash
./gradlew assembleRelease
```

输出：`app/build/outputs/apk/release/app-release.apk`

### 4. 优化和混淆

Release 版本已启用：
- ✅ 代码混淆（ProGuard）
- ✅ 资源压缩
- ✅ APK 体积优化

## 🔍 测试

### 单元测试

```bash
./gradlew test
```

### UI 测试

```bash
./gradlew connectedAndroidTest
```

## 📊 项目信息

**当前版本：** 1.0.0  
**最低 SDK：** 24 (Android 7.0)  
**目标 SDK：** 34 (Android 14)  
**包名：** com.copypasteengine.clipboardbridge

## 🐛 常见问题

### Gradle 同步失败

1. 检查网络连接
2. 清理缓存：`./gradlew clean`
3. 重新同步：File → Sync Project with Gradle Files

### 构建失败

```bash
# 清理并重新构建
./gradlew clean assembleDebug
```

### 找不到 SDK

设置 ANDROID_HOME 环境变量：
```bash
# Windows
setx ANDROID_HOME "C:\Users\你的用户名\AppData\Local\Android\Sdk"

# Linux/macOS
export ANDROID_HOME=$HOME/Android/Sdk
```

## 📝 开发指南

### 修改服务器默认配置

编辑 `ClipboardViewModel.kt`：

```kotlin
private fun loadSettings() {
    // 修改默认服务器地址
    val defaultUrl = "http://192.168.1.100:5678"
    // ...
}
```

### 添加新功能

应用使用 MVVM 架构：
- **UI**: `ui/ClipboardBridgeScreen.kt`
- **逻辑**: `ui/ClipboardViewModel.kt`
- **数据**: `data/ClipboardRepository.kt`
- **网络**: `network/ClipboardApiService.kt`

### 自定义 UI 主题

编辑 `ui/theme/Color.kt` 修改配色方案。

## 📄 许可证

MIT License - 同主项目

---

**需要帮助？** 查看主 [README](../README.md) 或提交 [Issue](https://github.com/copypasteengine/clipboard-bridge/issues)

