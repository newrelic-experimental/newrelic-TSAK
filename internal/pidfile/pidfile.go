// Package pidfile manages pid files.
package pidfile

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/facebookgo/atomicfile"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)


var	errNotConfigured = errors.New("pidfile not configured")
var	Pidfile string


// IsNotConfigured returns true if the error indicates the pidfile location has
// not been configured.
func IsNotConfigured(err error) bool {
	if err == errNotConfigured {
		return true
	}
	return false
}

// GetPidfilePath returns the configured pidfile path.
func GetPidfilePath() string {
	return Pidfile
}

// SetPidfilePath sets the pidfile path.
func SetPidfilePath(p string) {
	Pidfile = p
}

// Write the pidfile based on the flag. It is an error if the pidfile hasn't
// been configured.
func Write() error {
	if Pidfile == "" {
		return errNotConfigured
	}

	if err := os.MkdirAll(filepath.Dir(Pidfile), os.FileMode(0755)); err != nil {
		return err
	}

	file, err := atomicfile.New(Pidfile, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("error opening pidfile %s: %s", Pidfile, err)
	}
	defer file.Close() // in case we fail before the explicit close

  log.Trace(fmt.Sprintf("Application PID() = %v", os.Getpid()))
	_, err = fmt.Fprintf(file, "%d", os.Getpid())
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

// Read the pid from the configured file. It is an error if the pidfile hasn't
// been configured.
func Read() (int, error) {
	if Pidfile == "" {
		return 0, errNotConfigured
	}

	d, err := ioutil.ReadFile(Pidfile)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(bytes.TrimSpace(d)))
	if err != nil {
		return 0, fmt.Errorf("error parsing pid from %s: %s", Pidfile, err)
	}

	return pid, nil
}
