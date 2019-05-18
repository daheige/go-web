package Logger

import (
	"go-web/app/utils"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/daheige/thinkgo/common"
)

/**
 * {
    "level": 200,
    "level_name": "info",
    "local_time": "2019-05-07 23:31:26.439",
    "msg": "exec end",
    "line_no": 38,
    "file_path": "/web/hg-mux/app/extension/Logger/log.go",
    "context":
    {
        "current_file": "RequestWare.go",
        "current_line": 36,
        "host": "[::1]:52846",
        "log_id": "f0421e4bf186d512ced7c1228ac8f9ee",
        "method": "GET",
        "options":
        {
            "exec_time": 0.000320906
        },
        "plat": "web",
        "request_uri": "/",
        "tag": "_",
        "ua": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"
    }
}
*/
func writeLog(req *http.Request, levelName string, message interface{}, context map[string]interface{}) {
	tag := strings.Replace(req.RequestURI, "/", "_", -1)
	ua := req.Header.Get("User-Agent")

	//日志信息
	logId := req.Context().Value("log_id")
	if logId == nil {
		logId = common.RndUuidMd5()
	}

	//函数调用信息
	_, file, line, _ := runtime.Caller(2)

	logInfo := map[string]interface{}{
		"tag":          tag,
		"request_uri":  req.RequestURI,
		"log_id":       logId,
		"options":      context,
		"host":         req.RemoteAddr,
		"current_line": line,
		"current_file": filepath.Base(file),
		"ua":           ua,
		"plat":         utils.GetDeviceByUa(ua), //当前设备匹配
		"method":       req.Method,
	}

	switch levelName {
	case "info":
		common.InfoLog(message, logInfo)
	case "debug":
		common.DebugLog(message, logInfo)
	case "notice":
		common.NoticeLog(message, logInfo)
	case "warn":
		common.WarnLog(message, logInfo)
	case "error":
		common.ErrorLog(message, logInfo)
	case "emergency":
		common.EmergLog(message, logInfo)
	}
}

func Info(req *http.Request, message interface{}, context map[string]interface{}) {
	writeLog(req, "info", message, context)
}

func Debug(req *http.Request, message interface{}, context map[string]interface{}) {
	writeLog(req, "debug", message, context)
}

func Notice(req *http.Request, message interface{}, context map[string]interface{}) {
	writeLog(req, "notice", message, context)
}

func Warn(req *http.Request, message interface{}, context map[string]interface{}) {
	writeLog(req, "warn", message, context)
}

func Error(req *http.Request, message interface{}, context map[string]interface{}) {
	writeLog(req, "error", message, context)
}

func Emergency(req *http.Request, message interface{}, context map[string]interface{}) {
	writeLog(req, "emergency", message, context)
}
