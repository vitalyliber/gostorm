# gostorm

**Tool for destroying sites**

## How to use gostorm?

- Build or download version for your OS

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

iOS

```bash
GOOS=darwin GOARCH=amd64 go build -o gostorm_ios
```

Windows

```bash
GOOS=windows GOARCH=amd64 go build -o gostorm_windows
```