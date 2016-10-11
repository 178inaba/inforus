package inforus

import (
	"path"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// Skip the following 3 frames.
// runtime.Callers
// github.com/178inaba/inforus.Hook.Fire
// github.com/178inaba/inforus.(*Hook).Fire
const skipFrameCnt = 3

// AddHookDefault is ...
func AddHookDefault() {
	log.AddHook(Hook{file: true, function: false, line: true})
}

// AddHook is ...
func AddHook(file, function, line bool) {
	log.AddHook(Hook{file: file, function: function, line: line})
}

// Hook is ...
type Hook struct {
	file     bool
	function bool
	line     bool
}

// Levels is ...
func (h Hook) Levels() []log.Level {
	return log.AllLevels
}

// Fire is ...
func (h Hook) Fire(entry *log.Entry) error {
	pc := make([]uintptr, 64)
	cnt := runtime.Callers(skipFrameCnt, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i])
		name := fu.Name()
		if !strings.Contains(name, "github.com/Sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			if h.file {
				entry.Data["file"] = path.Base(file)
			}

			if h.function {
				entry.Data["func"] = path.Base(name)
			}

			if h.line {
				entry.Data["line"] = line
			}

			break
		}
	}

	return nil
}
