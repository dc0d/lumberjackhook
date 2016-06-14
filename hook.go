//Package lumberjackhook provides a logrus hook for sending logs to lumberjack
package lumberjackhook

import (
	"fmt"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/Sirupsen/logrus"
)

//InitLumberjackHook creates & initializes a LumberjackHook
func InitLumberjackHook(levels []logrus.Level, lumberjackLogger *lumberjack.Logger) *LumberjackHook {
	res := new(LumberjackHook)

	_l := levels
	if len(_l) == 0 {
		_l = defaultLevels()
	}
	res._levels = _l
	res._lumberjack = lumberjackLogger

	return res
}

//LumberjackHook is a logrus hook sends logs to InfluxDB
type LumberjackHook struct {
	_levels     []logrus.Level
	_lumberjack *lumberjack.Logger
}

//Levels implementation of interface logrus.Hook
func (h *LumberjackHook) Levels() []logrus.Level {
	return h._levels
}

//Fire implementation of interface logrus.Hook
func (h *LumberjackHook) Fire(entry *logrus.Entry) error {
	s, err := entry.String()
	if err != nil {
		return err
	}
	_, err = h._lumberjack.Write([]byte(s))
	return err
}

func getVal(fields logrus.Fields, key string) (string, bool) {
	value, ok := fields[key]
	if ok {
		return fmt.Sprintf("%v", value), ok
	}
	return "", ok
}

func defaultLevels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
