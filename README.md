# udp-ping

Simple UDP client and echo server for doing a ICMP style ping but with UDP instead

## build client for windows

GOOS=windows GOARCH=amd64 go build -o udp-ping.exe client/main.go


## build for (old) linux server

CGO_ENABLED=0 go build -o udp-echo server/main.go




