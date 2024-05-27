//go:build darwin || freebsd || windows

package stats

import (
	"os"
	"runtime"
	"strings"
)

func Arch() (string, error) {
	return strings.ToLower(runtime.GOARCH), nil
}

func Platform() (string, error) {
	return "", nil
}

func Distro() (string, error) {
	return "", nil
}

func Hostname() (string, error) {
	return os.Hostname()
}

func Uptime() (uint64, error) {
	return 0, nil
}
