//go:build darwin || freebsd || windows

package stats

import (
	"errors"
	"time"
)

func Network(d time.Duration) ([]NetworkStats, error) {
	return nil, errors.New("platform not supported")
}
