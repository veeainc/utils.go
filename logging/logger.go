package logging

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"path/filepath"
	"runtime"
	"strings"
)

var logger = logrus.New()

type LF = logrus.Fields // alias

func init() {
	logger.SetFormatter(&prefixed.TextFormatter{
		ForceColors:true,
		FullTimestamp:true,
		ForceFormatting: true,
		SpacePadding: 0,
	})
	logger.SetLevel(logrus.TraceLevel)
}

func SetLogLevel(level logrus.Level) {
	logger.SetLevel(level)
}

func GetNamedLogger() * logrus.Entry {
	_, file, _, ok := runtime.Caller(1)
	if ok {
		filename := filepath.Base(file)
		name := strings.TrimSuffix(filename, filepath.Ext(filename))
		return logger.WithFields(logrus.Fields{"prefix": name})
	}
	panic("cannot get named logger")
}

