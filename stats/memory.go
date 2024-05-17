//go:build darwin || freebsd || windows

package stats

import "errors"

func Memory() (*MemoryStats, error) {
	return nil, errors.New("platform not supported")
}
