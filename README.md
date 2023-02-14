# GoScanner
This is a CLI subnet scanner

Use this to "ping" devices on a given subnet on a given port via the TCP protocol. Returns response if the device is up or down.
Utilizes Go's goroutines to ping all devices concurrently, which results in a runtime with 254 addresses of 1.0232 seconds with a wait timeout of 1s to await respons from host.

# Usage
go run ping.go 192.168.1 80
or
.\ping.exe 192.168.1 80
