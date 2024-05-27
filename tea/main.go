package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path"
	"termisphere-agent/stats"
	"time"
)

func main() {
	app := &cli.App{
		Name:  "tea",
		Usage: "termisphere agent",
	}

	app.Commands = append(app.Commands, &cli.Command{
		Name:      "copy",
		Usage:     "copy data to clipboard by osc 52",
		ArgsUsage: "[data]",
		Action: func(c *cli.Context) error {
			data := c.Args().First()
			if data == "" {
				return nil
			}
			if data == "-" {
				data = ""
				for {
					var b [1024]byte
					n, err := os.Stdin.Read(b[:])
					if err != nil {
						break
					}
					data += string(b[:n])
				}
			}

			data = base64.StdEncoding.EncodeToString([]byte(data))

			fmt.Print("\x1B]52;c;" + data + "\x07")
			fmt.Println("copy data to clipboard")

			return nil
		},
	})

	app.Commands = append(app.Commands, &cli.Command{
		Name:      "open",
		Usage:     "open a editor or file manager in termisphere app",
		ArgsUsage: "[path]",
		Action: func(c *cli.Context) error {
			p := c.Args().First()
			if p == "" {
				p = "."
			}

			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			p = path.Join(wd, p)

			stat, err := os.Stat(p)
			if err != nil {
				return err
			}

			if !stat.IsDir() {
				fmt.Print("\x1B]1332;OpenFile=" + p + "\x07")
				fmt.Println("open editor for", p)
			} else {
				fmt.Print("\x1B]1332;OpenDirectory=" + p + "\x07")
				fmt.Println("open file manager for", p)
			}

			return nil
		},
	})

	app.Commands = append(app.Commands, &cli.Command{
		Name:      "stats",
		Usage:     "stats report system stats",
		ArgsUsage: "[duration]",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "arch", Usage: "report arch info"},
			&cli.BoolFlag{Name: "platform", Usage: "report platform info"},
			&cli.BoolFlag{Name: "distro", Usage: "report distro info"},
			&cli.BoolFlag{Name: "hostname", Usage: "report hostname info"},
			&cli.BoolFlag{Name: "uptime", Usage: "report uptime stats"},
			&cli.BoolFlag{Name: "cpu", Usage: "report cpu stats"},
			&cli.BoolFlag{Name: "memory", Usage: "report memory stats"},
			&cli.BoolFlag{Name: "block-device", Usage: "report block device stats"},
			&cli.BoolFlag{Name: "filesystem", Usage: "report filesystem stats"},
			&cli.BoolFlag{Name: "network", Usage: "report network stats"},
			&cli.BoolFlag{Name: "all", Usage: "report all stats"},
			&cli.BoolFlag{Name: "json", Usage: "output json format"},
		},
		Action: func(c *cli.Context) error {
			arch := c.Bool("arch")
			platform := c.Bool("platform")
			distro := c.Bool("distro")
			hostname := c.Bool("hostname")
			uptime := c.Bool("uptime")
			cpu := c.Bool("cpu")
			memory := c.Bool("memory")
			blockDevice := c.Bool("block-device")
			filesystem := c.Bool("filesystem")
			network := c.Bool("network")
			all := c.Bool("all")

			a := c.Args().First()
			if a == "" {
				a = "1s"
			}
			d, err := time.ParseDuration(a)
			if err != nil {
				return err
			}

			if res, err := stats.Fetch(d, stats.Request{
				Arch:        all || arch,
				Platform:    all || platform,
				Distro:      all || distro,
				Hostname:    all || hostname,
				Uptime:      all || uptime,
				Cpu:         all || cpu,
				Memory:      all || memory,
				BlockDevice: all || blockDevice,
				Filesystem:  all || filesystem,
				Network:     all || network,
			}); err == nil {
				b, err := json.MarshalIndent(res, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(b))
			} else {
				b, err := json.MarshalIndent(map[string]any{"error": err.Error()}, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(b))
			}

			return nil
		},
	})

	app.Commands = append(app.Commands, &cli.Command{
		Name:  "ping",
		Usage: "ping print pong",
		Action: func(c *cli.Context) error {
			fmt.Println("pong")

			return nil
		},
	})

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
