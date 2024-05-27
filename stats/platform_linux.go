package stats

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func Arch() (string, error) {
	return strings.ToLower(runtime.GOARCH), nil
}

func Platform() (string, error) {
	return "linux", nil
}

func Distro() (string, error) {
	release, err := godotenv.Read("/etc/os-release")
	if err != nil {
		return "", err
	}

	distro, _ := release["ID"]
	if distro == "" {
		distro = "unknown"
	}

	return distro, nil
}

func Hostname() (string, error) {
	return os.Hostname()
}

func Uptime() (uint64, error) {
	contents, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(string(contents))
	if len(fields) < 1 {
		return 0, errors.New("unexpected content in /proc/uptime")
	}

	uptime, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, err
	}

	return uint64(uptime), nil
}
