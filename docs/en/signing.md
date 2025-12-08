# 🔐 Android APK Signing

Understanding Android APK signing types.

## 📝 Signing Types

| Feature | Debug Signing | Release Signing |
|---------|--------------|-----------------|
| **Can Install** | ✅ Yes | ✅ Yes |
| **User Experience** | ⚠️ "Test app" warning | ✅ Normal install |
| **Google Play** | ❌ No | ✅ Yes |
| **Configuration** | ✅ Automatic | ❌ Manual setup |
| **Key Security** | ⚠️ Public | ✅ Private |
| **APK Size** | Larger (~5MB) | Smaller (~3MB) |
| **Code Obfuscation** | ❌ No | ✅ Yes |
| **GitHub Actions** | ✅ Auto-build | 📝 Needs config |
| **Use Case** | GitHub/OSS distribution | Google Play Store |

## 🎯 Current Release

**We provide Debug-signed APK:**
- ✅ Ready to install
- ✅ Fully functional
- ⚠️ System shows "test app" warning
- ✅ Perfect for open source distribution

**User installation:**
1. Download APK
2. Enable "Install from unknown sources"
3. Install (may show "test app" warning)
4. Tap "Install anyway"
5. ✓ Works perfectly!

## 🔧 Adding Release Signing (Optional)

If you want to publish to Google Play or remove the "test app" warning:

See [Chinese Signing Guide](../zh-CN/SIGNING.md) for detailed configuration steps.

---

**Back to:** [Building Guide](./building.md)

