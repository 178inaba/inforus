package filerus

import (
	"path"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// AddHook is ...
func AddHook() {
	log.AddHook(Hook{})
}

// Hook is ...
type Hook struct{}

// Levels is ...
func (h Hook) Levels() []log.Level {
	return log.AllLevels
}

// Fire is ...
func (h Hook) Fire(entry *log.Entry) error {
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/Sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(name)
			entry.Data["line"] = line
			break
		}
	}

	return nil
}
