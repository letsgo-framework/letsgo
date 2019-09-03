package log

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

var logger = log.New(colorable.NewColorableStderr(), "", 0)

type logFunc func(string, ...interface{})
type logFuncInterface func(...interface{})

var (
	Debug logFunc
	Debugln logFuncInterface
	Info  logFunc
	Infoln logFuncInterface
	Warn  logFunc
	Warnln logFuncInterface
	Error logFunc
	Errorln logFuncInterface
)

var colors = map[string]string{
	"reset":          "0",
	"red":            "31",
	"green":          "32",
	"yellow":         "33",
	"blue":           "34",
	"magenta":        "35",
	"cyan":           "36",
	"bold_red":       "31;1",
	"bold_green":     "32;1",
	"bold_yellow":    "33;1",
	"bold_blue":      "34;1",
	"bold_magenta":   "35;1",
	"bold_cyan":      "36;1",
	"bright_red":     "31;2",
	"bright_green":   "32;2",
	"bright_yellow":  "33;2",
	"bright_blue":    "34;2",
	"bright_magenta": "35;2",
	"bright_cyan":    "36;2",
}

var settings = map[string]string{
	"log_color_debug": "cyan",
	"log_color_info":  "green",
	"log_color_warn":  "yellow",
	"log_color_error": "red",
}

var levels = []string{"error", "warn", "info", "debug"}

func getenv() string {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "debug"
	}

	return logLevel
}


func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func newLogFunc(prefix string, omitLog bool) func(string, ...interface{}) {
	color, clear := "", ""
	color = fmt.Sprintf("\033[%sm", logColor(prefix))
	clear = fmt.Sprintf("\033[%sm", colors["reset"])
	prefix = fmt.Sprintf("%-11s", prefix)

	if prefix != "error" {
		return func(format string, v ...interface{}) {
			now := time.Now()
			timeString := fmt.Sprintf("%d:%d:%02d", now.Hour(), now.Minute(), now.Second())
			format = fmt.Sprintf("%s%s %s |%s %s", color, timeString, prefix, clear, format)
			if omitLog == false {
				logger.Printf(format, v...)
			}
		}
	} else {
		return func(format string, v ...interface{}) {
			now := time.Now()
			timeString := fmt.Sprintf("%d:%d:%02d", now.Hour(), now.Minute(), now.Second())
			format = fmt.Sprintf("%s%s %s |%s %s", color, timeString, prefix, clear, format)
			if omitLog == false {
				logger.Fatalf(format, v...)
			}
		}
	}
}

func newLogFuncInterface(prefix string, omitLog bool) func(...interface{}) {
	color, clear := "", ""
	color = fmt.Sprintf("\033[%sm", logColor(prefix))
	clear = fmt.Sprintf("\033[%sm", colors["reset"])
	prefix = fmt.Sprintf("%-11s", prefix)

	if prefix != "error" {
		return func(v ...interface{}) {
			now := time.Now()
			timeString := fmt.Sprintf("%d:%d:%02d", now.Hour(), now.Minute(), now.Second())
			format := fmt.Sprintf("%s%s %s |%s", color, timeString, prefix, clear)
			if omitLog == false {
				logger.Println(format)
				logger.Println(v...)
			}
		}
	} else {
		return func(v ...interface{}) {
			now := time.Now()
			timeString := fmt.Sprintf("%d:%d:%02d", now.Hour(), now.Minute(), now.Second())
			format := fmt.Sprintf("%s%s %s |%s", color, timeString, prefix, clear)
			if omitLog == false {
				logger.Println(format)
				logger.Println(v...)
			}
		}
	}
}

func logColor(logName string) string {
	settingsKey := fmt.Sprintf("log_color_%s", logName)
	colorName := settings[settingsKey]

	return colors[colorName]
}

func Fatal(err error) {
	logger.Fatal(err)
}

func InitLogFuncs() {
	// Configure Logging
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "./log/letsgo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	targetLevel := getenv()

	targetLevelIndex := find(levels, targetLevel)

	if targetLevelIndex == 3 {
		Debug = newLogFunc("debug", false)
		Debugln = newLogFuncInterface("debug", false)
		Info = newLogFunc("info", false)
		Infoln = newLogFuncInterface("info", false)
		Warn = newLogFunc("warn", false)
		Warnln = newLogFuncInterface("warn", false)
		Error = newLogFunc("error", false)
		Errorln = newLogFuncInterface("error", false)
	}
	if targetLevelIndex == 2 {
		Debug = newLogFunc("debug", true)
		Debugln = newLogFuncInterface("debug", true)
		Info = newLogFunc("info", false)
		Infoln = newLogFuncInterface("info", false)
		Warn = newLogFunc("warn", false)
		Warnln = newLogFuncInterface("warn", false)
		Error = newLogFunc("error", false)
		Errorln = newLogFuncInterface("error", false)
	}
	if targetLevelIndex == 1 {
		Debug = newLogFunc("debug", true)
		Info = newLogFunc("info", true)
		Infoln = newLogFuncInterface("info", true)
		Warn = newLogFunc("warn", false)
		Warnln = newLogFuncInterface("warn", false)
		Error = newLogFunc("error", false)
		Errorln = newLogFuncInterface("error", false)
	}
	if targetLevelIndex == 0 {
		Debug = newLogFunc("debug", true)
		Info = newLogFunc("info", true)
		Infoln = newLogFuncInterface("info", true)
		Warn = newLogFunc("warn", true)
		Warnln = newLogFuncInterface("warn", true)
		Error = newLogFunc("error", false)
		Errorln = newLogFuncInterface("error", false)
	}

	Debug("Log level set to %s", targetLevel)
}