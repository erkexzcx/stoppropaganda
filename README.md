# StopPropaganda

Russia has invaded Ukraine and commited various war crimes. Russian media denies any of its attacks on civilian as well as denies any Russian troops casualties. According to them, they are doing this "special operation" to protect Ukrainians from...Ukraine.

Mykhailo Federov (Vice Prime Minister and Minister of Digital Transformation of Ukraine) has shared [this twitter post](https://twitter.com/FedorovMykhailo/status/1497642156076511233) encouraging cyber attack on certain targets via Telegram group. This will be primary source of the target websites for this application.

Some foreign media and even countries (e.g. Belarus) publicitly support Russian aggression towards Ukraine, therefore I created this simple DOS application that targets certain websites/endpoints in order to take them down or significantly distrupt their services.

**DISCLAIMER**: (D)DOS'ing is **illegal**! Usage of this tool is intended for educational purposes only.

# Usage

## Docker

Easiest way is to use Docker:
```bash
docker run -d --ulimit nofile=128000:128000 -p "8049:8049/tcp" erikmnkl/stoppropaganda
```

Use environment variables to change settings (for example `--env SP_WORKERS=50`) to change configuration. Available environment variables (and their defaults):
```
SP_BIND=":8049"
SP_WORKERS="20"
SP_TIMEOUT="10s"
SP_USERAGENT="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36"
```

You can also use `docker-compose`:
```yaml
services:
  stoppropaganda:
    image: erikmnkl/stoppropaganda
    container_name: stoppropaganda
    restart: unless-stopped
    ports:
      - "8049:8049/tcp"
    environment:
      SP_BIND: ":8049"
      SP_WORKERS: "20"
      SP_TIMEOUT: "10s"
      SP_USERAGENT: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36"
    ulimits:
      nofile:
        soft: 128000
        hard: 128000
```

**NOTE**: `SP_WORKERS` means workers per website, not in total. For example, 5 websites * 20 workers = 100 workers in total.

Then you can see status in this URL: `http://<ip>:8049/status`

## Binaries

Download binary from [releases](https://github.com/erkexzcx/stoppropaganda/releases/).

Use from terminal:

```bash
# Show help
$ ./stoppropaganda_v0.0.1_linux_x86_64 --help

# Use with defaults
./stoppropaganda_v0.0.1_linux_x86_64

# Use with increased workers count (you might experience "too many open files" error on some systems)
./stoppropaganda_v0.0.1_linux_x86_64 --workers 50
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
ExecStart=/path/to/binary --workers 50
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
# Increase open file limits for current terminal session
ulimit -n 128000

# Run without compiling to binary
go run stoppropaganda.go --help

# Build binary and run it
go build -o stoppropaganda stoppropaganda.go
./stoppropaganda --help
```

You can also build for other architectures/platforms as well, see `build.sh` file.

# Recommendations

* Increase `workers` count from 20 (default) to e.g. 100 for greater effect, but check the logs if you are not getting `too many open files`. If so, see [this](https://stackoverflow.com/questions/880557/socket-accept-too-many-open-files).
* Change `useragent` to yours. See [this](https://www.whatismybrowser.com/detect/what-is-my-user-agent/).
* General recommendation is to use VPN, but this is not necesarry. Remember - DOS/DDOS is **illegal**.

# Inspiration

This application was inspired by the following projects:
* https://www.reddit.com/r/hacking/comments/t1a8is/simple_html_dos_script_for_russian_sites/
* https://norussian.tk/
* https://stop-russian-desinformation.near.page/
* https://russianwarshipgofuckyourself.club/
