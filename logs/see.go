package logs

import "github.com/cihub/seelog"

/**
seelog配置
*/
func init() {
	logger, err := seelog.LoggerFromConfigAsFile("./res/seelog.xml")

	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	seelog.ReplaceLogger(logger)
}

func Error(err ...interface{}) {
	defer seelog.Flush()
	seelog.Error(err)
}

func Info(infos ...interface{}) {
	defer seelog.Flush()
	seelog.Info(infos)
}

func Debug(msg ...interface{}) {
	defer seelog.Flush()
	seelog.Debug(msg)
}
