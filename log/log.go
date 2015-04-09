/*
Package log provides support for logging to stdout and stderr.

Log entries will be logged in the following format:

    timestamp hostname tag[pid]: SEVERITY Message
*/
package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

type ConfdFormatter struct {
}

func (c *ConfdFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	return []byte(fmt.Sprintf("%s %s %s[%d]: %s %s\n", timestamp, hostname, tag, os.Getpid(), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

// tag represents the application name generating the log message. The tag
// string will appear in all log entires.
var tag string

func init() {
	tag = os.Args[0]
	log.SetFormatter(&ConfdFormatter{})
}

// SetTag sets the tag.
func SetTag(t string) {
	tag = t
}

// SetQuiet sets quiet mode.
//
// TODO (bacongobbler): remove entirely once v0.9.0 lands
func SetQuiet() {
	Fatal("--quiet has been deprecated in favour of --log-level")
}

// SetDebug sets debug mode.
//
// TODO (bacongobbler): remove entirely once v0.9.0 lands
func SetDebug() {
	Fatal("--debug has been deprecated in favour of --log-level")
}

// SetVerbose sets verbose mode.
//
// TODO (bacongobbler): remove entirely once v0.9.0 lands
func SetVerbose() {
	Fatal("--verbose has been deprecated in favour of --log-level")
}

func SetLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		Fatal(fmt.Sprintf(`not a valid level: "%s"`, level))
	}
	log.SetLevel(lvl)
}

// Debug logs a message with severity DEBUG.
func Debug(msg string) {
	log.Debug(msg)
}

// Error logs a message with severity ERROR.
func Error(msg string) {
	log.Error(msg)
}

// Fatal logs a message with severity ERROR followed by a call to os.Exit().
func Fatal(msg string) {
	log.Fatal(msg)
}

// Info logs a message with severity INFO.
func Info(msg string) {
	log.Info(msg)
}

// Warning logs a message with severity WARNING.
func Warning(msg string) {
	log.Warning(msg)
}
