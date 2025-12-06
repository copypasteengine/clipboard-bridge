package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
)

// ------------------ 全局变量 ------------------

var (
	serverRunning = false
	server        *http.Server

	lastUpdated   int64  = time.Now().Unix()
	lastClipboard string = ""
	clipboardMu   sync.RWMutex // 保护剪贴板相关变量

	logFile *os.File
	cfgMu   sync.RWMutex // 保护配置变量
)

type Config struct {
	Port         int    `json:"port"`
	Token        string `json:"token"`
	AutoStart    bool   `json:"auto_start"`
	AutoFirewall bool   `json:"auto_firewall"`
	LogLevel     string `json:"log_level"` // info, debug, error
}

var cfg Config

// 日志级别常量
const (
	LogLevelError = "error"
	LogLevelInfo  = "info"
	LogLevelDebug = "debug"
)

// 其他常量
const (
	MaxRequestBodySize = 10 * 1024 * 1024 // 10MB
	MinPort            = 1024
	MaxPort            = 65535
)

// ------------------ 工具函数 ------------------

func exeDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(exe)
}

// 获取本机局域网IP地址
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logError("获取网络接口失败", err)
		return "127.0.0.1"
	}

	var ips []string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip := ipNet.IP.String()
				// 过滤掉虚拟网卡的IP
				if !isVirtualIP(ip) {
					ips = append(ips, ip)
				}
			}
		}
	}

	logDebug(fmt.Sprintf("检测到的IP地址: %v", ips))

	if len(ips) == 0 {
		logInfo("未检测到有效的网络IP，使用 127.0.0.1")
		return "127.0.0.1"
	}

	// 优先返回常见私有网段的IP
	// 192.168.x.x (家庭网络)
	for _, ip := range ips {
		if len(ip) >= 7 && ip[:7] == "192.168" {
			logDebug(fmt.Sprintf("选择IP: %s (192.168 网段)", ip))
			return ip
		}
	}

	// 172.16.x.x - 172.31.x.x (企业网络, RFC 1918)
	for _, ip := range ips {
		if strings.HasPrefix(ip, "172.") {
			parts := strings.Split(ip, ".")
			if len(parts) >= 2 {
				if secondOctet, err := strconv.Atoi(parts[1]); err == nil {
					if secondOctet >= 16 && secondOctet <= 31 {
						logDebug(fmt.Sprintf("选择IP: %s (172.16-31 网段)", ip))
						return ip
					}
				}
			}
		}
	}

	// 10.x.x.x (大型企业网络)
	for _, ip := range ips {
		if len(ip) >= 3 && ip[:3] == "10." {
			logDebug(fmt.Sprintf("选择IP: %s (10.x 网段)", ip))
			return ip
		}
	}

	// 如果都没有，返回第一个
	logDebug(fmt.Sprintf("选择IP: %s (第一个可用IP)", ips[0]))
	return ips[0]
}

// 判断是否是虚拟网卡IP
func isVirtualIP(ip string) bool {
	// 只过滤掉明显的问题IP
	// 169.254.x.x 是 APIPA 自动分配的地址，通常表示网络配置有问题
	if len(ip) >= 7 && ip[:7] == "169.254" {
		return true
	}

	return false
}

// ------------------ 日志系统 ------------------

func initLogger() {
	logPath := filepath.Join(exeDir(), "clipboard_bridge.log")
	var err error
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 如果无法创建日志文件，尝试写到临时目录
		logPath = filepath.Join(os.TempDir(), "clipboard_bridge.log")
		logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return // 放弃日志记录
		}
	}

	// 同时输出到标准输出和日志文件
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	logInfo(t("log_separator"))
	logInfo(t("program_start", logPath))
	logInfo(t("log_separator"))
}

func shouldLog(level string) bool {
	configLevel := cfg.LogLevel
	if configLevel == "" {
		configLevel = LogLevelInfo // 默认 info 级别
	}

	// error 级别：只记录错误
	// info 级别：记录错误和信息（默认）
	// debug 级别：记录所有内容

	switch configLevel {
	case LogLevelError:
		return level == LogLevelError
	case LogLevelInfo:
		return level == LogLevelError || level == LogLevelInfo
	case LogLevelDebug:
		return true
	default:
		return level == LogLevelError || level == LogLevelInfo
	}
}

func logInfo(message string) {
	if !shouldLog(LogLevelInfo) {
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] [INFO] %s", timestamp, message)
	if logFile != nil {
		logFile.WriteString(logMsg + "\n")
		logFile.Sync()
	}
	fmt.Println(logMsg)
}

func logDebug(message string) {
	if !shouldLog(LogLevelDebug) {
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] [DEBUG] %s", timestamp, message)
	if logFile != nil {
		logFile.WriteString(logMsg + "\n")
		logFile.Sync()
	}
	fmt.Println(logMsg)
}

func logError(message string, err error) {
	if !shouldLog(LogLevelError) {
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] [ERROR] %s: %v", timestamp, message, err)
	if logFile != nil {
		logFile.WriteString(logMsg + "\n")
		logFile.Sync()
	}
	fmt.Println(logMsg)
}

// ------------------ 配置管理 ------------------

func loadConfig() {
	path := filepath.Join(exeDir(), "config.json")
	data, err := os.ReadFile(path)

	cfgMu.Lock()
	defer cfgMu.Unlock()

	if err != nil {
		// 不存在就用默认配置
		cfg = Config{
			Port:         5678,
			Token:        "",
			AutoStart:    false,
			AutoFirewall: true,
			LogLevel:     "info",
		}
		return
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Println("读取 config.json 失败，使用默认配置:", err)
		cfg = Config{
			Port:         5678,
			Token:        "",
			AutoStart:    false,
			AutoFirewall: true,
			LogLevel:     "info",
		}
	}

	// 验证和修正配置
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}
	if cfg.LogLevel != LogLevelError && cfg.LogLevel != LogLevelInfo && cfg.LogLevel != LogLevelDebug {
		fmt.Printf("无效的日志级别: %s，使用默认值 info\n", cfg.LogLevel)
		cfg.LogLevel = "info"
	}
	if cfg.Port < MinPort || cfg.Port > MaxPort {
		fmt.Printf("无效的端口号: %d，使用默认值 5678\n", cfg.Port)
		cfg.Port = 5678
	}
}

func saveConfig() error {
	cfgMu.RLock()
	configCopy := cfg
	cfgMu.RUnlock()

	path := filepath.Join(exeDir(), "config.json")
	data, err := json.MarshalIndent(configCopy, "", "  ")
	if err != nil {
		logError("序列化配置失败", err)
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		logError("保存配置文件失败", err)
		return err
	}

	logDebug("配置已保存到文件: " + path)
	return nil
}

// ------------------ HTTP 服务 ------------------

func handlePush(w http.ResponseWriter, r *http.Request) {
	if !checkToken(r) {
		logInfo(fmt.Sprintf("Push 请求: Token 验证失败 (来自 %s)", r.RemoteAddr))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	// 限制请求体大小
	r.Body = http.MaxBytesReader(w, r.Body, MaxRequestBodySize)

	r.ParseForm()
	text := r.FormValue("text")

	if text == "" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logError("读取请求体失败", err)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to read request body"))
			return
		}
		text = string(body)
	}

	logInfo(fmt.Sprintf("收到 Push 请求 (来自 %s)，内容长度: %d 字节", r.RemoteAddr, len(text)))
	logDebug(fmt.Sprintf("Push 内容详情: %s", getPreview(text)))

	// 写入剪贴板，并检查错误
	err := clipboard.WriteAll(text)
	if err != nil {
		logError("写入剪贴板失败", err)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to write clipboard: " + err.Error()))
		return
	}

	clipboardMu.Lock()
	lastClipboard = text
	lastUpdated = time.Now().Unix()
	clipboardMu.Unlock()

	logInfo(fmt.Sprintf("✓ 成功写入剪贴板，内容长度: %d 字节", len(text)))

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("OK"))
}

func getPreview(text string) string {
	if len(text) > 50 {
		return text[:50] + "..."
	}
	return text
}

func handlePull(w http.ResponseWriter, r *http.Request) {
	if !checkToken(r) {
		logInfo(fmt.Sprintf("Pull 请求: Token 验证失败 (来自 %s)", r.RemoteAddr))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	logInfo(fmt.Sprintf("收到 Pull 请求 (来自 %s)", r.RemoteAddr))
	text, err := clipboard.ReadAll()
	if err != nil {
		logError("读取剪贴板失败", err)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to read clipboard: " + err.Error()))
		return
	}
	logInfo(fmt.Sprintf("✓ 成功读取剪贴板，内容长度: %d 字节", len(text)))
	logDebug(fmt.Sprintf("Pull 内容详情: %s", getPreview(text)))

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(text))
}

func handleMeta(w http.ResponseWriter, r *http.Request) {
	if !checkToken(r) {
		logInfo(fmt.Sprintf("Meta 请求: Token 验证失败 (来自 %s)", r.RemoteAddr))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Unauthorized",
		})
		return
	}

	logInfo(fmt.Sprintf("收到 Meta 请求 (来自 %s)", r.RemoteAddr))
	text, err := clipboard.ReadAll()
	if err != nil {
		logError("读取剪贴板失败", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	clipboardMu.RLock()
	updated := lastUpdated
	clipboardMu.RUnlock()

	meta := map[string]interface{}{
		"text":    text,
		"updated": updated,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(meta)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	if !checkToken(r) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	logDebug(fmt.Sprintf("收到 Ping 请求 (来自 %s)", r.RemoteAddr))
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("PONG"))
}

func startServer() {
	ensureFirewallRule()

	cfgMu.RLock()
	port := cfg.Port
	cfgMu.RUnlock()

	mux := http.NewServeMux()
	mux.HandleFunc("/push", handlePush)
	mux.HandleFunc("/pull", handlePull)
	mux.HandleFunc("/meta", handleMeta)
	mux.HandleFunc("/ping", handlePing)

	server = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	localIP := getLocalIP()
	go func() {
		logInfo(t("service_started"))
		logInfo(t("external_access", localIP, port))
		logInfo(t("local_access", port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logError(t("error_server"), err)
		}
	}()
}

func stopServer() {
	if server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logError(t("error_server_close"), err)
			server.Close() // 强制关闭
		}
		logInfo(t("service_stopped"))
	}
}

func checkToken(r *http.Request) bool {
	cfgMu.RLock()
	token := cfg.Token
	cfgMu.RUnlock()

	if token == "" {
		return true // 未设置 token 则不校验
	}
	t := r.Header.Get("X-Auth-Token")
	if t == "" {
		t = r.URL.Query().Get("token")
	}
	return t == token
}

// ------------------ 托盘 UI ------------------

func onReady() {
	systray.SetTitle(t("app_title"))

	cfgMu.RLock()
	port := cfg.Port
	autoStart := cfg.AutoStart
	cfgMu.RUnlock()

	localIP := getLocalIP()
	systray.SetTooltip(fmt.Sprintf("%s - %s:%d", t("app_title"), localIP, port))

	icon, _ := loadIcon()
	systray.SetIcon(icon)

	mInfo := systray.AddMenuItem(t("service_address", fmt.Sprintf("http://%s:%d", localIP, port)), t("ext_access"))
	mInfo.Disable()
	mLocal := systray.AddMenuItem(t("local_address", fmt.Sprintf("http://localhost:%d", port)), t("local_test"))
	mLocal.Disable()
	systray.AddSeparator()

	mAuto := systray.AddMenuItemCheckbox(t("auto_start"), "", autoStart)
	mStart := systray.AddMenuItem(t("start_service"), "")
	mStop := systray.AddMenuItem(t("stop_service"), "")
	systray.AddSeparator()
	mLog := systray.AddMenuItem(t("open_log"), "")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem(t("quit"), "")

	// 自动启动服务
	startServer()
	serverRunning = true
	mStart.Hide()

	go func() {
		for {
			select {
			case <-mAuto.ClickedCh:
				cfgMu.Lock()
				cfg.AutoStart = !cfg.AutoStart
				newAutoStart := cfg.AutoStart
				cfgMu.Unlock()

				if newAutoStart {
					mAuto.Check()
				} else {
					mAuto.Uncheck()
				}
				updateAutostart()
			case <-mStart.ClickedCh:
				if !serverRunning {
					startServer()
					serverRunning = true
					mStart.Hide()
					mStop.Show()
				}
			case <-mStop.ClickedCh:
				if serverRunning {
					stopServer()
					serverRunning = false
					mStop.Hide()
					mStart.Show()
				}
			case <-mLog.ClickedCh:
				openLogFile()
			case <-mQuit.ClickedCh:
				stopServer()
				if logFile != nil {
					logFile.Close()
				}
				systray.Quit()
				return
			}
		}
	}()
}

func loadIcon() ([]byte, error) {
	return iconData, nil
}

func onExit() {
	if logFile != nil {
		logFile.Close()
	}
}

func maskToken(token string) string {
	if token == "" {
		return "(未设置)"
	}
	if len(token) <= 4 {
		return "****"
	}
	return token[:2] + "****" + token[len(token)-2:]
}

// ------------------ main ------------------

func main() {
	initLanguage()  // 初始化语言
	initLogger()
	loadConfig()

	cfgMu.RLock()
	logInfo(t("config_loaded", cfg.Port, maskToken(cfg.Token), cfg.AutoStart, cfg.LogLevel))
	cfgMu.RUnlock()

	go initClipboardListener()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		logInfo("收到退出信号，程序即将退出")
		if server != nil {
			stopServer()
		}
		if logFile != nil {
			logFile.Close()
		}
		os.Exit(0)
	}()

	systray.Run(onReady, onExit)
}

