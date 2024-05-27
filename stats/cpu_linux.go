package stats

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"termisphere-agent/utils"
	"time"
)

type linuxCpuStats struct {
	us uint64
	sy uint64
	ni uint64
	id uint64
	wa uint64
	hi uint64
	si uint64
	st uint64
}

func Cpu(d time.Duration) ([]CpuStats, error) {
	stat, err := os.ReadFile("/proc/stat")
	if err != nil {
		return nil, err
	}

	var res []linuxCpuStats
	for _, line := range strings.Split(string(stat), "\n") {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		if !strings.HasPrefix(parts[0], "cpu") {
			continue
		}
		if parts[0] == "cpu" {
			continue
		}

		cpuID, err := strconv.ParseInt(parts[0][3:], 10, 64)
		if err != nil {
			return nil, err
		}
		for len(res) <= int(cpuID) {
			res = append(res, linuxCpuStats{})
		}

		res[cpuID], err = parseCPU(parts[1:])
		if err != nil {
			return nil, err
		}
	}

	time.Sleep(d)

	stat, err = os.ReadFile("/proc/stat")
	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(stat), "\n") {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		if !strings.HasPrefix(parts[0], "cpu") {
			continue
		}
		if parts[0] == "cpu" {
			continue
		}

		cpuID, err := strconv.ParseInt(parts[0][3:], 10, 64)
		if err != nil {
			return nil, err
		}

		data, err := parseCPU(parts[1:])
		if err != nil {
			return nil, err
		}

		res[cpuID] = linuxCpuStats{
			us: data.us - res[cpuID].us,
			sy: data.sy - res[cpuID].sy,
			ni: data.ni - res[cpuID].ni,
			id: data.id - res[cpuID].id,
			wa: data.wa - res[cpuID].wa,
			hi: data.hi - res[cpuID].hi,
			si: data.si - res[cpuID].si,
			st: data.st - res[cpuID].st,
		}
	}

	return utils.SlicesMap(res, func(i linuxCpuStats) CpuStats {
		return CpuStats{
			User:   i.us,
			System: i.sy,
			Idle:   i.id,
			Wait:   i.wa,
			Steal:  i.st,
			Sum:    i.us + i.sy + i.ni + i.id + i.wa + i.hi + i.si + i.st,
		}
	}), nil
}

func parseCPU(fields []string) (linuxCpuStats, error) {
	if len(fields) < 8 {
		return linuxCpuStats{}, fmt.Errorf("not enough fields in cpu stat")
	}

	us, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		return linuxCpuStats{}, err
	}

	sy, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return linuxCpuStats{}, err
	}

	ni, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return linuxCpuStats{}, err
	}

	id, err := strconv.ParseUint(fields[3], 10, 64)
	if err != nil {
		return linuxCpuStats{}, err
	}

	wa, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return linuxCpuStats{}, err
	}

	hi, err := strconv.ParseUint(fields[5], 10, 64)
	if err != nil {
		return linuxCpuStats{}, err
	}

	si, err := strconv.ParseUint(fields[6], 10, 64)
	if err != nil {
		return linuxCpuStats{}, err
	}

	st, err := strconv.ParseUint(fields[7], 10, 64)
	if err != nil {
		return linuxCpuStats{}, err
	}

	return linuxCpuStats{us: us, sy: sy, ni: ni, id: id, wa: wa, hi: hi, si: si, st: st}, nil
}
