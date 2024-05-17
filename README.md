## Termisphere Agent

Termisphere Agent 是一个用 Go 语言编写的实用程序，主要目的是为了辅助 Termisphere 软件从 SSH 服务器上获取数据。

Termisphere Agent is a utility program written in the Go language designed primarily to assist the Termisphere software in fetching data from SSH servers.

### 功能 / Features

- OSC52 剪切板共享 / OSC52 Clipboard Sharing
- 在 Termisphere 中打开文件或目录 / Open file or directory in Termisphere
- 监看系统运行数据 / Monitor system running stats

更多命令帮助请查看 `tea help`。

For more command help, please see `tea help`.

### 安装 / Installation

Linux / amd64:

```bash
sudo rm -f /usr/local/bin/tea && sudo wget https://mirror.ghproxy.com/https://github.com/codemutex/termisphere-agent/releases/latest/download/tea_linux_amd64 -O /usr/local/bin/tea && sudo chmod +x /usr/local/bin/tea
```

Linux / arm64:

```bash
sudo rm -f /usr/local/bin/tea && sudo wget https://mirror.ghproxy.com/https://github.com/codemutex/termisphere-agent/releases/latest/download/tea_linux_arm64 -O /usr/local/bin/tea && sudo chmod +x /usr/local/bin/tea
```