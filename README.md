# StopPropaganda

Russia has invaded Ukraine and commited various war crimes. Russian media says that Russia has not commited any war crimes, has no casualties and they doing this "special operation" to protect Ukrainians from...Ukraine.

I believe that Russian propaganda websites should be down for their propaganda, therefore I created a simple DOS application that can be deployed almost anywhere.

## Usage

### Docker

### Binaries

### Building from source

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
