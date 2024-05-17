//go:build darwin || freebsd || windows

package stats

import (
	"errors"
	"time"
)

func BlockDevice(d time.Duration) ([]BlockDeviceStats, error) {
	return nil, errors.New("platform not supported")
}
