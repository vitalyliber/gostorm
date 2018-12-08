# gostorm

Minimalistic performance and load testing tool written in Go.

## How to use gostorm?

- Download and install to MacOS

```
curl -L https://github.com/vitalyliber/gostorm/releases/download/v.1.0.1/gostorm_macos -o /usr/local/bin/gostorm
sudo chmod a+x /usr/local/bin/gostorm
```

- Download and install to Linux


```
curl -L https://github.com/vitalyliber/gostorm/releases/download/v.1.0.1/gostorm_linux -o /usr/local/bin/gostorm
sudo chmod a+x /usr/local/bin/gostorm
```

- Run and paste url:
```bash
Enter site url: https://example.com
```
- Enter number of streams:

```bash
Enter number of streams: 500
```

- Check the site accessible

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