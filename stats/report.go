package stats

import (
	"sync"
	"time"
)

func Fetch(d time.Duration, req Request) (res SystemStats, err error) {
	var wg sync.WaitGroup

	if req.Arch {
		res.Arch, err = Arch()
		if err != nil {
			return SystemStats{}, err
		}
	}

	if req.Platform {
		res.Platform, err = Platform()
		if err != nil {
			return SystemStats{}, err
		}
	}

	if req.Distro {
		res.Distro, err = Distro()
		if err != nil {
			return SystemStats{}, err
		}
	}

	if req.Hostname {
		res.Hostname, err = Hostname()
		if err != nil {
			return SystemStats{}, err
		}
	}

	if req.Uptime {
		res.Uptime, err = Uptime()
		if err != nil {
			return SystemStats{}, err
		}
	}

	if req.Cpu {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res.Cpu, err = Cpu(d)
		}()
	}

	if req.Memory {
		res.Memory, err = Memory()
		if err != nil {
			return SystemStats{}, err
		}
	}

	if req.BlockDevice {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res.BlockDevice, err = BlockDevice(d)
		}()
	}

	if req.Filesystem {
		res.Filesystem, err = Filesystem()
		if err != nil {
			return SystemStats{}, err
		}
	}

	if req.Network {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res.Network, err = Network(d)
		}()
	}

	if wg.Wait(); err != nil {
		return SystemStats{}, err
	}

	return res, nil
}

type Request struct {
	Arch        bool
	Platform    bool
	Distro      bool
	Hostname    bool
	Uptime      bool
	Cpu         bool
	Memory      bool
	BlockDevice bool
	Filesystem  bool
	Network     bool
}

type SystemStats struct {
	Arch        string             `json:"arch,omitempty"`
	Platform    string             `json:"platform,omitempty"`
	Distro      string             `json:"distro,omitempty"`
	Hostname    string             `json:"hostname,omitempty"`
	Uptime      uint64             `json:"uptime,omitempty"`
	Cpu         []CpuStats         `json:"cpu,omitempty"`
	Memory      *MemoryStats       `json:"memory,omitempty"`
	BlockDevice []BlockDeviceStats `json:"block_device,omitempty"`
	Filesystem  []FilesystemStats  `json:"filesystem,omitempty"`
	Network     []NetworkStats     `json:"network,omitempty"`
}

type CpuStats struct {
	User   uint64 `json:"user"`
	System uint64 `json:"system"`
	Idle   uint64 `json:"idle"`
	Wait   uint64 `json:"wait"`
	Steal  uint64 `json:"steal"`
	Sum    uint64 `json:"sum"`
}

type MemoryStats struct {
	Total     uint64 `json:"total"`
	Free      uint64 `json:"free"`
	Available uint64 `json:"available"`
	Buffer    uint64 `json:"buffer"`
	Cache     uint64 `json:"cache"`
	SwapTotal uint64 `json:"swap_total"`
	SwapFree  uint64 `json:"swap_free"`
	SwapCache uint64 `json:"swap_cache"`
}

type BlockDeviceStats struct {
	Name       string `json:"name"`
	Model      string `json:"model"`
	Vendor     string `json:"vendor"`
	Serial     string `json:"serial"`
	Size       uint64 `json:"size"`
	Read       uint64 `json:"read"`
	ReadSpeed  uint64 `json:"read_speed"`
	Write      uint64 `json:"write"`
	WriteSpeed uint64 `json:"write_speed"`
}

type FilesystemStats struct {
	Mount  string `json:"mount"`
	Device string `json:"device"`
	Type   string `json:"type"`
	Total  uint64 `json:"total"`
	Free   uint64 `json:"free"`
}

type NetworkStats struct {
	Name    string   `json:"name"`
	Speed   uint64   `json:"speed"`
	MAC     string   `json:"mac"`
	MTU     uint64   `json:"mtu"`
	IP      []string `json:"ip"`
	RX      uint64   `json:"rx"`
	RXSpeed uint64   `json:"rx_speed"`
	TX      uint64   `json:"tx"`
	TXSpeed uint64   `json:"tx_speed"`
}
