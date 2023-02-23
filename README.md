# ProxySetup-Legacy
This is a simple way to run a HTTP proxy server

## Installation

Of course, you'll need to install go : https://go.dev/dl/

Start by downloading goproxy lib and install it https://github.com/goproxy/goproxy

Then clone this repository or just grab the content of the Main.go file
```bash
    git clone https://github.com/Dampen59/ProxySetup-Legacy.git
```

Open Main.go file in any text editor and edit the networkInterface IPv4 and the port to your situation. You can find the networkInterface value with any terminal using "ipconfig" on Windows or "ifconfig"/"ip a" on Linux systems.

## Running

You can just go with :

```bash
    go run ./Main.go
```

## Building
```bash
    go build ./Main.go
```

It will output a binary package ./Main with all libraries embedded which you can run anywhere.
