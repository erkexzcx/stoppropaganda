# StopPropaganda

Russia has invaded Ukraine and commited various war crimes. Russian media says that Russia has not commited any war crimes, has no casualties and they doing this "special operation" to protect Ukrainians from...Ukraine.

I believe that Russian propaganda websites should be down for their propaganda, therefore I created a simple DOS application that can be deployed almost anywhere.

## Usage

### Docker

Easiest way is to use Docker:
```bash
docker run -d erikmnkl/stoppropaganda
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
      SP_WORKERS: "100"
      SP_TIMEOUT: "10s"
      SP_USERAGENT: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36"
```

### Binaries

Download binary from [releases](https://github.com/erkexzcx/stoppropaganda/releases/).

Use from terminal:

```bash
# Show help
$ ./stoppropaganda_v0.0.1_linux_x86_64 --help

# Use with defaults
./stoppropaganda_v0.0.1_linux_x86_64

# Use with increased workers count (you might experience "too many open files" error on some systems)
./stoppropaganda_v0.0.1_linux_x86_64 --workers 1000
```

Then open in your browser to see the status: http://127.0.0.1:8049/status

You might want to create SystemD script (Linux only) to autostart this on boot. Create `/etc/systemd/system/stoppropaganda.service` with below contents:
```
[Unit]
Description=Stoppropaganda service
After=network-online.target

[Service]
ExecStart=/path/to/binary --workers 1000
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

Then `systemctl daemon-reload && systemctl enable --now stoppropaganda.service`. To stop, use `systemctl stop stoppropaganda.service`.

## Building from source

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
# Run without compiling to binary
go run stoppropaganda.go --help

# Build binary and run it
go build -o stoppropaganda stoppropaganda.go
./stoppropaganda --help
```

You can also build for other architectures/platforms as well, see `build.sh` file.

## Inspiration

This application was inspired by the following projects:
* https://www.reddit.com/r/hacking/comments/t1a8is/simple_html_dos_script_for_russian_sites/
* https://norussian.tk/
* https://stop-russian-desinformation.near.page/
