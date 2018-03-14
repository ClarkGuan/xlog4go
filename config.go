package xlog4go

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type ConfFileWriter struct {
	On              bool   `json:"On"`
	LogPath         string `json:"LogPath"`
	RotateLogPath   string `json:"RotateLogPath"`
	LogRetain       int    `json:"LogRetain"`
	WfLogPath       string `json:"WfLogPath"`
	RotateWfLogPath string `json:"RotateWfLogPath"`
	WfLogRetain     int    `json:"WfLogRetain"`
}

type ConfConsoleWriter struct {
	On    bool `json:"On"`
	Color bool `json:"Color"`
}

type LogConfig struct {
	Level string            `json:"LogLevel"`
	FW    ConfFileWriter    `json:"FileWriter"`
	CW    ConfConsoleWriter `json:"ConsoleWriter"`
}

func SetupLogWithPtr(lc *LogConfig) (err error) {
	if lc.FW.On {
		if len(lc.FW.LogPath) > 0 {
			w := NewFileWriter()
			w.SetFileName(lc.FW.LogPath)
			w.SetPathPattern(lc.FW.RotateLogPath)
			w.SetLogLevelFloor(TRACE)
			if len(lc.FW.WfLogPath) > 0 {
				w.SetLogLevelCeil(INFO)
			} else {
				w.SetLogLevelCeil(ERROR)
			}
			w.SetRetainHour(lc.FW.LogRetain)
			Register(w)
		}

		if len(lc.FW.WfLogPath) > 0 {
			wfw := NewFileWriter()
			wfw.SetFileName(lc.FW.WfLogPath)
			wfw.SetPathPattern(lc.FW.RotateWfLogPath)
			wfw.SetLogLevelFloor(WARNING)
			wfw.SetLogLevelCeil(ERROR)
			wfw.SetRetainHour(lc.FW.WfLogRetain)
			Register(wfw)
		}
	}

	if lc.CW.On {
		w := NewConsoleWriter()
		w.SetColor(lc.CW.Color)
		Register(w)
	}

	switch lc.Level {
	case "trace":
		SetLevel(TRACE)

	case "debug":
		SetLevel(DEBUG)

	case "info":
		SetLevel(INFO)

	case "warning":
		SetLevel(WARNING)

	case "error":
		SetLevel(ERROR)

	case "fatal":
		SetLevel(FATAL)

	default:
		err = errors.New("Invalid log level")
	}
	return
}

func SetupLogWithJson(cnt []byte) (err error) {
	var lc LogConfig

	if err = json.Unmarshal(cnt, &lc); err != nil {
		return
	}

	return SetupLogWithPtr(&lc)
}

func SetupLogWithFile(file string) (err error) {
	if cnt, err := ioutil.ReadFile(file); err != nil {
		return err
	} else {
		return SetupLogWithJson(cnt)
	}
}
