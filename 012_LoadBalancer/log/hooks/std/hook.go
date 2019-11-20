package std

import (
	"io"
	"os"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/012_LoadBalancer/log/hooks"
	"github.com/sirupsen/logrus"
)

// Hook is a hook that writes logs of specified LogLevels to specified Writer
type Hook struct {
	hooks.Hook
	Writer io.Writer
}

// Init initializes terminal hook by the given logging level
func (hook *Hook) Init(level int) {
	hook.Writer = os.Stdout
	hook.SetLevel(level)
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *Hook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}
