package hooks

import "github.com/sirupsen/logrus"

// IHook interface includes all methods
type IHook interface {
	SetLevel(int)

	Levels() []logrus.Level
	Fire(entry *logrus.Entry) error
}

// Hook struct implements interface methods
type Hook struct {
	logLevels []logrus.Level
}

// SetLevel : set level for the hook
func (hook *Hook) SetLevel(level int) {
	if level < 0 || level > 6 {
		level = 4 // Default is Info level
	}
	for i := 0; i < level+1; i++ {
		hook.logLevels = append(hook.logLevels, logrus.AllLevels[i])
	}
}

// Levels returns all log levels
func (hook *Hook) Levels() []logrus.Level {
	return hook.logLevels
}

// Fire will identify how to 'print' the logging message.
// This method will be overridden when defining std hook or db hook
func (hook *Hook) Fire(entry *logrus.Entry) error {
	return nil
}
