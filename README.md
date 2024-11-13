# udp-ping

Simple UDP client and echo server for doing an ICMP style ping but with UDP instead

# TODO

add ability for the server to include the arrival time in the response so the RTT can be split

## build client for windows

    GOOS=windows GOARCH=amd64 go build -o udp-ping.exe client/main.go

## build for (old) linux server

    CGO_ENABLED=0 go build -o udp-echo server/main.go




