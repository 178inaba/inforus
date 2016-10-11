package filerus

import (
	"path"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// Skip the following.
// runtime.Callers
// github.com/178inaba/filerus.Hook.Fire
// github.com/178inaba/filerus.(*Hook).Fire
const skipFrameCnt = 3

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
	pc := make([]uintptr, 64)
	cnt := runtime.Callers(skipFrameCnt, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i])
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
