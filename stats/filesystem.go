//go:build darwin || freebsd || windows

package stats

import "errors"

func Filesystem() ([]FilesystemStats, error) {
	return nil, errors.New("platform not supported")
}
