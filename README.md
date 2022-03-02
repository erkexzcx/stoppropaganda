# StopPropaganda

[![Github All Releases](https://img.shields.io/github/downloads/erkexzcx/stoppropaganda/total.svg)](https://github.com/erkexzcx/stoppropaganda/releases)
[![Docker Hub Pulls](https://img.shields.io/docker/pulls/erikmnkl/stoppropaganda)](https://hub.docker.com/r/erikmnkl/stoppropaganda)
[![Go Report Card](https://goreportcard.com/badge/github.com/erkexzcx/stoppropaganda)](https://goreportcard.com/report/github.com/erkexzcx/stoppropaganda)
[![Tests](https://img.shields.io/github/workflow/status/erkexzcx/stoppropaganda/tests?maxAge=30&label=tests&logo=github)](https://github.com/erkexzcx/stoppropaganda/actions)
[![Release](https://img.shields.io/github/workflow/status/erkexzcx/stoppropaganda/release?maxAge=30&label=release&logo=github)](https://github.com/erkexzcx/stoppropaganda/actions)

Russia has invaded Ukraine and committed various war crimes. Russian media denies any of its attacks on civilian as well as denies any Russian troops casualties. According to them, they are doing this "special operation" to protect Ukrainians from...Ukraine.

Mykhailo Federov (Vice Prime Minister and Minister of Digital Transformation of Ukraine) has shared [this twitter post](https://twitter.com/FedorovMykhailo/status/1497642156076511233) encouraging cyber attack on certain targets via Telegram group. This will be primary source of the target websites for this application.

Some foreign media and even countries (e.g. Belarus) publicly support Russian aggression towards Ukraine, therefore I created this simple DOS application that targets certain websites/endpoints in order to take them down or significantly distrupt their services.

**DISCLAIMER**: (D)DOS'ing is **illegal**! Usage of this tool is intended for educational purposes only.

# Usage

## Docker

Easiest way is to use Docker:
```bash
# Download latest docker image
docker pull erikmnkl/stoppropaganda

# If exists, remove running container
docker rm -f stoppropaganda

# Create new container
docker run --name stoppropaganda -d --ulimit nofile=128000:128000 -p "8049:8049/tcp" erikmnkl/stoppropaganda
```

Use environment variables to change settings (for example `--env SP_WORKERS=50 SP_DNSWORKERS=500`) to change configuration. Available environment variables (and their defaults):
```
SP_WORKERS="20"
SP_TIMEOUT="10s"
SP_DNSWORKERS="100"
SP_DNSTIMEOUT="125ms"
SP_USERAGENT="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36"
```

**NOTE**: `SP_WORKERS` means workers per website, not in total. Same with `SP_DNSWORKERS`. For example, 5 websites * 20 workers = 100 workers in total.

Then you can see status in this URL: `http://<ip>:8049/status`  
or without browser (Linux only): `curl http://<ip>:8049/status | less`

## docker-compose

`docker-compose.yaml` and other examples are available.

Usage:
```bash
# Pull latest image
docker-compose pull

# Create/recreate container
docker-compose up -d
```

Also see [Docker](#docker) for additional information.

## Kubernetes

You can also use `kubectl`:
```bash
kubectl create ns stoppropaganda
kubectl apply -f stoppropaganda.yaml
```
**NOTE**: edit `stoppropaganda.yaml` with required number of replicas.

Also see [Docker](#docker) for additional information.

## Binaries

Download binary from [releases](https://github.com/erkexzcx/stoppropaganda/releases/).

Use from terminal:

```bash
# (Linux only) make the binary executable
chmod +x stoppropaganda_v0.0.1_linux_x86_64

# (Linux only) Increase open files limit for current terminal session
ulimit -n 128000

# Show help
$ ./stoppropaganda_v0.0.1_linux_x86_64 --help

# Use with defaults
./stoppropaganda_v0.0.1_linux_x86_64

# Use with increased workers count (you might experience "too many open files" error on some systems)
./stoppropaganda_v0.0.1_linux_x86_64 --workers 50 --dnsworkers 500
```

Then open in your browser to see the status: http://127.0.0.1:8049/status

You might want to create SystemD script (Linux only) to autostart this on boot. Create `/etc/systemd/system/stoppropaganda.service` with below contents:
```
[Unit]
Description=Stoppropaganda service
After=network-online.target

[Service]
LimitAS=infinity
LimitRSS=infinity
LimitCORE=infinity
LimitNOFILE=128000
ExecStart=/path/to/binary --workers 50 --dnsworkers 500
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

Then `systemctl daemon-reload && systemctl enable --now stoppropaganda.service`. To stop, use `systemctl stop stoppropaganda.service`.

# Building from source

Uninstall any existing Golang installations if you installed from official Linux repos. They are usually outdated and might not work at all.

Download and install Golang using [these instructions](https://go.dev/doc/install). Validate if Golang binary works:
```bash
$ go version
go version go1.17.7 linux/amd64
```

Download this repo:
```bash
git clone https://github.com/erkexzcx/stoppropaganda.git
cd stoppropaganda
```

Now you have 2 options to run this application:
```bash
# (Linux only) Increase open files limit for current terminal session
ulimit -n 128000

# Run without compiling to binary
go run cmd/stoppropaganda/main.go --help

# Build binary and run it
go build -o stoppropaganda cmd/stoppropaganda/main.go
./stoppropaganda --help
```

You can also build for other architectures/platforms as well, see `build.sh` file.

# Recommendations

* Increase `workers`/`dnsworkers` count from 20/100 (default) to e.g. 100/1000 for greater effect, but check the logs if you are not getting `too many open files`. If so, see [this](https://stackoverflow.com/questions/880557/socket-accept-too-many-open-files).
* Adjust dnstimeout based on your location.  Eastern America ~125-150ms and in Europe this is likely much lower.  To properly adjust this value, check the /status page and if all queries are "successful", lower this value ~20ms and try again until "success" queries are low and thus "timeout errors" increase.
* Change `useragent` to yours (used for websites only). See [this](https://www.whatismybrowser.com/detect/what-is-my-user-agent/).
* General recommendation is to use VPN, but this is not necessary. Remember - DOS/DDOS is **illegal**.

# Inspiration

This application was inspired by the following projects:
* https://www.reddit.com/r/hacking/comments/t1a8is/simple_html_dos_script_for_russian_sites/
* https://norussian.xyz/
* https://stop-russian-desinformation.near.page/
* https://russianwarshipgofuckyourself.club/
