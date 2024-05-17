//go:build darwin || freebsd || windows

package stats

import (
	"errors"
	"time"
)

func CPU(d time.Duration) ([]CpuStats, error) {
	return nil, errors.New("platform not supported")
}
