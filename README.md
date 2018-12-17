# gostorm

Minimalistic performance and load testing tool written in Go.

## How to use gostorm?

- Download and install to MacOS

```
curl -L https://github.com/vitalyliber/gostorm/releases/download/v1.0.2/gostorm_macos -o /usr/local/bin/gostorm
sudo chmod a+x /usr/local/bin/gostorm
```

- Download and install to Linux


```bash
curl -L https://github.com/vitalyliber/gostorm/releases/download/v1.0.2/gostorm_linux -o /usr/local/bin/gostorm
sudo chmod a+x /usr/local/bin/gostorm
```

- Run with flags:
```bash
gostorm -url=https://some-url.ru -streams=500 -timeout=1
```

- Check the site accessible

## Flags

```bash
gostorm -h
```

Output:
```bash
Usage of ./main:
  -streams int
        a number of streams (default 42)
  -timeout int
        timeout in minutes for working of program (default 1)
  -url string
        a site url (default "https://some-url.com")

```

## Build for your platform

Linux

```bash
GOOS=linux GOARCH=amd64 go build -o gostorm_lunux
```

Mac OS

```bash
GOOS=darwin GOARCH=amd64 go build -o gostorm_macos
```

Windows

```bash
GOOS=windows GOARCH=amd64 go build -o gostorm_windows
```
