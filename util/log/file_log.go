package log

import (
	"github.com/ctbsea/Go-Message/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"os"
	"strings"
	"time"
)

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

const deleteFileOnExit = true

func todayFilename() string {
	today := time.Now().Format("2006-01-02")
	return today + ".txt"
}

func  newLogFile(path string) *os.File {
	filename := todayFilename()
	//打开一个输出文件，如果重新启动服务器，它将追加到今天的文件中
	f, err := os.OpenFile(path + filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func NewRequestLogger(config config.Config) (h iris.Handler, close func() error) {
	close = func() error { return nil }
	c := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}
	path := config.Path.Runtime
	logFile := newLogFile(path)
	close = func() error {
		err := logFile.Close()
		if deleteFileOnExit {
			err = os.Remove(logFile.Name())
		}
		return err
	}
	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := logger.Columnize(
			now.Format("2006/01/02 - 15:04:05"),
			latency,
			status,
			ip,
			method,
			path,
			message,
			headerMessage)
		logFile.Write([]byte(output))
	}
	//我们不想使用记录器，一些静态请求等
	c.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
	h = logger.New(c)
	return
}
