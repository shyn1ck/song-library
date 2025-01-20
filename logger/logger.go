package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"song-library/configs"
)

var (
	Info    *log.Logger
	Error   *log.Logger
	Warning *log.Logger
	Debug   *log.Logger
)

func Init() error {
	logParams := configs.AppSettings.LogParams

	if _, err := os.Stat(logParams.LogDirectory); os.IsNotExist(err) {
		err = os.Mkdir(logParams.LogDirectory, 0755)
		if err != nil {
			return err
		}
	}

	lumberLogInfo := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogInfo),
		MaxSize:    logParams.MaxSizeMegabytes, // megabytes
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,   // Days.
		Compress:   logParams.Compress, // Disabled by default.
		LocalTime:  logParams.LocalTime,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogError),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	lumberLogWarning := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogWarn),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogDebug),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	gin.DefaultWriter = io.MultiWriter(lumberLogInfo)

	Info = log.New(gin.DefaultWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(lumberLogError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(lumberLogWarning, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(lumberLogDebug, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}
