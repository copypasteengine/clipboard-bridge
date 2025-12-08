# 🤝 Contributing to Clipboard Bridge

Thank you for considering contributing! / 感谢你考虑为项目做贡献！

## 🌍 Ways to Contribute / 贡献方式

### 1. Report Bugs / 报告问题

Found a bug? [Open an issue](https://github.com/copypasteengine/clipboard-bridge/issues/new)

**Please include / 请包含：**
- Operating system and version
- Steps to reproduce
- Expected vs actual behavior
- Log files (if applicable)

### 2. Suggest Features / 功能建议

Have an idea? [Open an issue](https://github.com/copypasteengine/clipboard-bridge/issues/new)

**Note / 注意：** This project focuses on **text clipboard sync only**. Image/file transfer suggestions will be redirected to appropriate projects.

### 3. Improve Documentation / 改进文档

- Fix typos / 修正错别字
- Improve clarity / 提高清晰度
- Add examples / 添加示例
- Translate to new languages / 翻译新语言

### 4. Submit Code / 提交代码

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes
4. Test thoroughly
5. Commit: `git commit -m "feat: add new feature"`
6. Push: `git push origin feature/my-feature`
7. Open a Pull Request

## 📝 Code Guidelines / 代码规范

### Go Code

- Follow standard Go formatting: `gofmt`
- Add comments for exported functions
- Handle errors properly
- Use meaningful variable names

### Kotlin Code

- Follow Kotlin coding conventions
- Use Jetpack Compose best practices
- Add KDoc comments for public functions
- Test on different Android versions

## 🌐 Translation Guidelines / 翻译指南

### Adding New Language / 添加新语言

**For Android App:**
1. Copy `android-app/app/src/main/res/values/strings.xml`
2. Create `values-XX/strings.xml` (XX = language code)
3. Translate all strings
4. Test on device

**For Desktop Service:**
1. Edit `i18n.go`
2. Add new language to `translations` map
3. Update `detectSystemLanguage()` function
4. Test with `LANG=xx_XX.UTF-8`

**For Documentation:**
1. Create `docs/XX/` directory
2. Translate markdown files
3. Update `docs/README.md` index

## ✅ Pull Request Checklist / PR 检查清单

- [ ] Code follows project style
- [ ] All tests pass
- [ ] Documentation updated
- [ ] Commits have clear messages
- [ ] No sensitive information in commits

## 🙏 Thank You! / 谢谢！

Every contribution helps make this project better!

---

**Questions?** Open an issue or discussion!

