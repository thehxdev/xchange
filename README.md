# xChange

[xChange](https://github.com/thehxdev/xchange) is a command-line tool to manage [Xray-Core's](https://github.com/xtls/xray-core) config files from local configs.

## Build form source
The `go` compiler and `git` program must be already installed on your system.
**This tools is only for linux users!**

1. Clone the repository
```bash
git clone --depth 1 --branch main https://github.com/thehxdev/xchange.git
cd xchange
```
2. Build source code
```bash
go build .
```

## Arguments
`-dir`: Your Config files directory (default `~/.local/xconfs`)
`-xconf`: Path to Xray's Config file (default `/etc/xray/config.json`)
`-myip`: Print you public IP address to `stdout`
`-tun2socks`: install `tun2socks` executable to `~/.local/bin` (amd64)

The default behavior is changing config file. `-myip` and `-tun2socks` are single flags so use one of them at the time.
`-dir` and `-xconf` flags are used to tell the app where local config files are located and where is Xray's main config file. So they are NOT operational flags and you can use them at the same time to specify the correct directory for local configs and correct path for xray config file.

### Examples
```bash
xchange -myip #shows your public IP address

xchange -tun2socks #install/update tun2socks (amd64)

# your config pool -> ~/xray-configs
# xray main config -> /etc/xray/config.json
xchange -dir ~/xray-configs -xconf /etc/xray/config.json
```

