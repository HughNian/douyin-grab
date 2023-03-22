package logger

import (
	"douyin-grab/pkg/file"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type Level int

var (
	F  *os.File
	F2 *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	debug      *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

	logMod = ""
)

var colors = map[string]func(a ...interface{}) string{
	"Warning": color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
	"Panic":   color.New(color.BgRed).Add(color.Bold).SprintFunc(),
	"Error":   color.New(color.FgRed).Add(color.Bold).SprintFunc(),
	"Info":    color.New(color.FgCyan).Add(color.Bold).SprintFunc(),
	"Debug":   color.New(color.FgWhite).Add(color.Bold).SprintFunc(),
}

var spaces = map[string]string{
	"Warning": "",
	"Panic":   "  ",
	"Error":   "  ",
	"Info":    "   ",
	"Debug":   "  ",
}

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup initialize the log instance
func Init(tag string) {
	var err error
	filePath := getLogFilePath()
	fileName := tag + getLogFileName()
	logMod = os.Getenv("LOG_MODE")
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
	F2, err = file.MustOpen("debug", filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	debug = log.New(F2, DefaultPrefix, log.LstdFlags)
	// 这里启动一个携程去定时检测当前时间 是否是同一天
	go func() {
		for {
			last := time.Now().Format(os.Getenv("TIME_FORMAT"))
			time.Sleep(1 * time.Second)
			now := time.Now().Format(os.Getenv("TIME_FORMAT"))

			if last != now {
				filePath := getLogFilePath()
				fileName := tag + getLogFileName()
				F, err = file.MustOpen(fileName, filePath)
				if err != nil {
					log.Fatalf("logging.Setup err: %v", err)
				}
				logger = log.New(F, DefaultPrefix, log.LstdFlags)
			}
		}
	}()
}

//func checkDate() {
//	curDate := time.Now().Format(os.Getenv("TIME_FORMAT"))
//	if lastDate != curDate {
//		Init()
//	}
//}

func File(fileName, msg string) {
	filePath := getLogFilePath()
	logMod = os.Getenv("LOG_MODE")
	f, err := file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger := log.New(f, DefaultPrefix, log.LstdFlags)

	logger.Println(msg)
}

// Debug output logs at debug level
func Debug(format string, v ...interface{}) {
	setPrefix(DEBUG)
	msg := fmt.Sprintf(format, v...)
	if logMod == "debug" {
		Println("Debug", msg)

	}
	debug.Println(msg)
}

// Info output logs at info level
func Info(format string, v ...interface{}) {
	setPrefix(INFO)
	msg := fmt.Sprintf(format, v...)
	Println("Info", msg)
	logger.Println(msg)
}

// Warn output logs at warn level
func Warn(format string, v ...interface{}) {
	setPrefix(WARNING)
	msg := fmt.Sprintf(format, v...)
	if logMod == "debug" {
		Println("Warning", msg)
	}
	logger.Println(msg)
}

// Error output logs at error level
func Error(format string, v ...interface{}) {
	setPrefix(ERROR)
	msg := fmt.Sprintf(format, v...)
	if logMod == "debug" {
		Println("Error", msg)
	}
	logger.Println(msg)
}

// Fatal output logs at fatal level
func Fatal(format string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func Println(prefix string, msg string) {
	// TODO Release时去掉

	c := color.New()

	_, _ = c.Printf(
		"%s%s %s %s\n",
		colors[prefix]("["+prefix+"]"),
		spaces[prefix],
		time.Now().Format("2006-01-02 15:04:05"),
		msg,
	)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)

	//checkDate()
}
