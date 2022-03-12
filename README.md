# StopPropaganda

[![Github All Releases](https://img.shields.io/github/downloads/erkexzcx/stoppropaganda/total.svg)](https://github.com/erkexzcx/stoppropaganda/releases)
[![Docker Hub Pulls](https://img.shields.io/docker/pulls/erikmnkl/stoppropaganda)](https://hub.docker.com/r/erikmnkl/stoppropaganda)
[![Go Report Card](https://goreportcard.com/badge/github.com/erkexzcx/stoppropaganda)](https://goreportcard.com/report/github.com/erkexzcx/stoppropaganda)

Russia has invaded Ukraine and committed various war crimes. Russian media denies any of its attacks on civilian as well as any Russian troops casualties. According to them, they are doing this "special operation" to protect Ukrainians from...Ukraine. On top of that, some foreign media and even countries (e.g. Belarus) publicly support Russian aggression, therefore I created this simple DOS application that targets those websites/endpoints in order to take them down or significantly disrupt their services.

Mykhailo Federov (Vice Prime Minister and Minister of Digital Transformation of Ukraine) has shared [this twitter post](https://twitter.com/FedorovMykhailo/status/1497642156076511233) encouraging cyber attack on certain targets via Telegram group. This will be primary source of the target websites for this application.

*Русский военный корабль, иди нахуй!*

**DISCLAIMER**: (D)DOS'ing is **illegal**! Usage of this tool is intended for educational purposes only.

- [Usage](#usage)
  * [Docker](#docker)
  * [docker-compose](#docker-compose)
  * [Kubernetes](#kubernetes)
  * [Android](#android)
  * [Binaries](#binaries)
- [Configuration](#configuration)
  * [bind](#bind)
  * [workers](#workers)
  * [timeout](#timeout)
  * [useragent](#useragent)
  * [dnsworkers](#dnsworkers)
  * [dnstimeout](#dnstimeout)
  * [dialspersecond](#dialspersecond)
  * [dialconcurrency](#dialconcurrency)
  * [proxy](#proxy)
  * [proxybypass](#proxybypass)
  * [algorithm](#algorithm)
  * [maxprocs](#maxprocs)
- [Web UI](#web-ui)
- [Building from source](#building-from-source)
- [Troubleshooting](#troubleshooting)
  * [Too many open files](#too-many-open-files)
  * [Crashing](#crashing)
  * [Detected as virus](#detected-as-virus)
- [Recommendations](#recommendations)
- [Inspiration](#inspiration)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>

# Usage

## Docker

Usage:
```bash
docker pull erikmnkl/stoppropaganda # Download latest docker image
docker rm -f stoppropaganda         # Remove existing container (if any)

# Run container
docker run --name stoppropaganda -d --ulimit nofile=128000:128000 -p "8049:8049/tcp" erikmnkl/stoppropaganda
```

Also see [Configuration](#configuration) and [Web UI](#web-ui). For `docker run`, pass environment variables using `-e` argument, for example `-e SP_WORKERS=50 -e SP_DNSWORKERS=500`.

## docker-compose

[docker-compose.yaml](https://github.com/erkexzcx/stoppropaganda/raw/main/docker-compose.yaml) and other docker-compose YAML examples are available.

Usage:
```bash
docker-compose pull  # Pull latest image
docker-compose up -d # Create/recreate container
```

See [Docker](#docker) for additional information. Also see [Configuration](#configuration) and [Web UI](#web-ui).

## Kubernetes

See [stoppropaganda.yaml](https://github.com/erkexzcx/stoppropaganda/raw/main/stoppropaganda.yaml).

You can also use `kubectl`:
```bash
kubectl create ns stoppropaganda
kubectl apply -f stoppropaganda.yaml
```
**NOTE**: edit `stoppropaganda.yaml` with required number of replicas.

See [Docker](#docker) for additional information. Also see [Configuration](#configuration) and [Web UI](#web-ui).

## Android

In order to use on Android:
- Install [Termux from Google Play](https://play.google.com/store/apps/details?id=com.termux).
- Install [Automate from Google Play](https://play.google.com/store/apps/details?id=com.llamalab.automate).
- Download [StopPropaganda_launcher.flo](https://github.com/erkexzcx/stoppropaganda/raw/main/StopPropaganda_launcher.flo) and import (simply opening it with Automate will import it).
- Launch imported Flow into Automate app and it will guide you step by step.

**NOTE**: On startup it will pull the latest version automatically, so in order to update - stop the app in Termux and re-run this Flow.

More advanced users might want to edit Automate flow themselves to further customize configuration. See [Configuration](#configuration) and [Usage](#usage).

## Binaries

Download binary from [releases](https://github.com/erkexzcx/stoppropaganda/releases/).

Additional steps needed for Linux prior usage:
```bash
# Make downloaded binary executable
chmod +x stoppropaganda_v0.0.1_linux_x86_64

# Increase open files limit for current terminal session to a maximum allowed by a kernel
ulimit -n unlimited
```

Usage:
```bash
# Execute binary
./stoppropaganda_v0.0.1_linux_x86_64 --help

# Example
./stoppropaganda_v0.0.1_linux_x86_64 --workers 10000 --dnsworkers 50000
```

Linux users might want to autostart this on boot, see [stoppropaganda.service](https://github.com/erkexzcx/stoppropaganda/raw/main/stoppropaganda.service). Upload that file to `/etc/systemd/system/stoppropaganda.service` with updated `ExecStart` value and then usage;
```bash
# Reload SystemD daaemon (after editing service file)
systemctl daemon-reload

# Then usage
systemctl enable stoppropaganda.service
systemctl disable stoppropaganda.service
systemctl start stoppropaganda.service
systemctl stop stoppropaganda.service
systemctl status stoppropaganda.service
systemctl kill stoppropaganda.service
journalctl -f -u stoppropaganda.service
```

# Configuration

Configuration can only be done in 2 ways:
* Command line arguments
* Environment variables

## bind

Configuration via command line argument `-bind ":8049"` or via environment variable `SP_BIND=":8049"`.

Default value of `:8049` is the same as `0.0.0.0:8049` which means web interface is accessible externally on port `8049`. If you want to limit web interface to be accessible only from the same host, use `127.0.0.1:8049`.

## workers

Configuration via command line argument `-workers 1000` or via environment variable `SP_WORKERS=1000`.

Default value of `1000` means that there will be a pool of 1000 workers that will DOS all the defined websites.

## timeout

Configuration via command line argument `-timeout 10s` or via environment variable `SP_TIMEOUT=10s`.

Default value of `10s` means that worker will wait for a website response for `10s` until it gives up.

## useragent

Configuration via command line argument `-useragent "..."` or via environment variable `SP_USERAGENT="..."`.

Default value is `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36`. User agent is sent with HTTP requests to defined websites.

## dnsworkers

Configuration via command line argument `-dnsworkers 100` or via environment variable `SP_DNSWORKERS=100`.

Default value of `100` means that there will be a pool of 100 workers that will DOS all the defined DNS servers.

## dnstimeout

Configuration via command line argument `-dnstimeout 1s` or via environment variable `SP_DNSTIMEOUT=1s`.

Default value of `1s` means that worker will wait for a DNS server response for `1s` until it gives up.

## dialspersecond

Configuration via command line argument `-dialspersecond 2500` or via environment variable `SP_DIALSPERSECOND=2500`.

Default value of `2500` means that there will be maximum of 2500 TCP SYN packets sent per second from fasthttp.

## dialconcurrency

Configuration via command line argument `-dialconcurrency 10000` or via environment variable `SP_DIALCONCURRENCY=2000`.

Default value of `10000` means that there will be maximum of 10000 concurrent dials from fasthttp.

## proxy

Configuration via command line argument `-proxy ""` or via environment variable `SP_PROXY=""`.

Proxy supports SOCKS4, SOCKS5 and HTTP proxies (or comma separated proxy chains). For example `-proxy "socks5://tor:9050"`.

Usage of this parameter can be combined with `proxybypass` parameter.

## proxybypass

Configuration via command line argument `-proxybypass ""` or via environment variable `SP_PROXYBYPASS=""`.

For example `-proxybypass "localhost"`.

This parameter is only applicable when used with [proxy](#proxy) parameter.

## algorithm

Configuration via command line argument `-algorithm fair` or via environment variable `SP_ALGORITHM="fair"`.

The algorithm defines in what manner you want websites to be DOS'ed. It directly impacts resource usage and performance of this application.

Available algorithms:
- `fair` (Default)
  - Known as "workers per website" (specified amount of workers will be divided equally for each website).
  - Specifying less workers than websites will result in some websites without workers.
  - Uses less CPU and RAM.
  - By nature it wastes more traffic and generally has bigger impact.
- `rr`
  - Known as "pool of workers" (each worker will take the next pending website from the queue).
  - Can be used with as low as 1 worker.
  - Uses more CPU and RAM.
  - By nature it prioritizes slower websites.


## maxprocs

Configuration via command line argument `-maxprocs 1` or via environment variable `SP_MAXPROCS=1`.

Defines amount of system threads (`runtime.GOMAXPROCS`) used by the program.

Default value of 1 provides some optimization, because Golang doesn't have to use mutexes, atomics 
and inter-process synchronization mechanisms.

# Web UI

As of now there is no fancy web interface, only a JSON pre-formatted output.

The following endpoints are available:
```
http://127.0.0.1:8049/status
http://127.0.0.1:8049/dnscache
http://127.0.0.1:8049/downloaded
```

Example usage on Linux:
```bash
# Simple:
curl http://127.0.0.1:8049/status

# Using JQ to format output
curl http://127.0.0.1:8049/status | jq .
```

Also see [bind](#bind) for host and port configuration.

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
# Increase open files limit for current terminal session to a maximum allowed by a kernel
ulimit -n unlimited

# Run without compiling to binary
go run cmd/stoppropaganda/main.go --help

# Build binary and run it
go build -o stoppropaganda cmd/stoppropaganda/main.go
./stoppropaganda --help
```

You can also build for other architectures/platforms as well, see `build.sh` file.

# Troubleshooting

## Too many open files

Most Linux distributions have limits on how many files (connections) can be opened to prevents things like [fork bomb](https://en.wikipedia.org/wiki/Fork_bomb).

More information on how to increase them [here](https://stackoverflow.com/questions/880557/socket-accept-too-many-open-files).

## Crashing

Reason 1: Make sure you have enough RAM. It is also wise to monitor RAM and CPU usage once application is started. Since March 3, with the latest release it was switched from workers per website/dns to pool of workers. Once you find a sweet spot resource-wise, there should be no need to change it with updates.

Reason 2: Work in progress. Always check if (a) Russians are still invading Ukraine and (b) there is a new release available.

## Detected as virus

Some anti-virus applications detects this application as virus. This application is not a virus and it has no malicious code in it. In fact, this application is open source and by the nature - everyone is free to inspect the code, improve it and build their own binaries by themselves (if you don't trust my binaries).

There are several reasons why this application might be tagged as virus:
* Putin wants this application to be treated as virus. [Haters do exist](https://github.com/erkexzcx/stoppropaganda/issues?q=label%3AAsshole).
* This application has hardcoded Yandex DNS servers (helps with some russian websites "geoblocking") that ignores your network settings.
* At the moment this application has ~400 hardcoded targets, which is not common in regular applications.
* Malware that is used to DDOS targets is usually working in pretty much the same manner as this application, so this application might be treated as such.

# Recommendations

* Increase `workers`/`dnsworkers` count for a greater effect. For example, `-workers 1000000` worked great with 1GBPS internet.
* Adjust `dnstimeout` based on your location. Change to something like `200ms` and see how it behaves. If "success" queries are low and thus "timeout errors" increase - increase timeout.
* Change `useragent` to yours (used for websites only). See [this](https://www.whatismybrowser.com/detect/what-is-my-user-agent/).
* General recommendation is to use VPN, but this is not necessary. Remember - DOS/DDOS is **illegal**.

# Inspiration

This application was inspired by the following projects:
* https://www.reddit.com/r/hacking/comments/t1a8is/simple_html_dos_script_for_russian_sites/
* https://norussian.xyz/
* https://stop-russian-desinformation.near.page/
* https://russianwarshipgofuckyourself.club/
