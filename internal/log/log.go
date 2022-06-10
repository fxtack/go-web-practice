package log

import (
	"go-web-practice/internal/config"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func LoadLog() {

	traceDir := filepath.Dir(config.Config.Log.Trace)
	infoDir := filepath.Dir(config.Config.Log.Info)
	warningDir := filepath.Dir(config.Config.Log.Warning)
	errorDir := filepath.Dir(config.Config.Log.Error)

	err := os.MkdirAll(traceDir, 0775)
	err = os.MkdirAll(infoDir, 0775)
	err = os.MkdirAll(warningDir, 0775)
	err = os.MkdirAll(errorDir, 0775)

	err = os.Chmod(traceDir, 0775)
	err = os.Chmod(infoDir, 0775)
	err = os.Chmod(warningDir, 0775)
	err = os.Chmod(errorDir, 0775)

	if err != nil {
		log.Fatalln(err, "create log file directory failed, please check configuration")
	}

	_, err = os.OpenFile(config.Config.Log.Trace, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fileInfo, err := os.OpenFile(config.Config.Log.Info, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fileWarning, err := os.OpenFile(config.Config.Log.Warning, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fileError, err := os.OpenFile(config.Config.Log.Error, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln(err, "log file open failed, please check configuration what file path you define.")
	}

	Trace = log.New(io.Discard,
		"Trace: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(os.Stdout, fileInfo),
		"Info : ",
		log.Ldate|log.Ltime)

	Warning = log.New(fileWarning,
		"Warning: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(os.Stderr, fileError),
		"Error: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
