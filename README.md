# Pingo

Pingo is a simple ping utility written in Go. It allows you to send ICMP echo requests to a specified domain name or ip address and receive echo replies.

### Flags

To see the usage information, use the `-h` flag:
```sh
Usage of pingo:
  -c string
        ip of the dns client (default "1.1.1.1")
  -d string
        name of the domain (default "google.com")
  -i int
        request interval in seconds (default 1)
  -t int
        request timeout in milliseconds (default 1000)
```

### Example Usage
```sh
pingo -d google.com
[216.239.38.120 -> 192.168.1.109] [RTT: 26ms, TTL: 116] [Seq: 0, Code: Echo Reply]
[216.239.38.120 -> 192.168.1.109] [RTT: 30ms, TTL: 116] [Seq: 1, Code: Echo Reply]
[216.239.38.120 -> 192.168.1.109] [RTT: 30ms, TTL: 116] [Seq: 2, Code: Echo Reply]
[216.239.38.120 -> 192.168.1.109] [RTT: 28ms, TTL: 116] [Seq: 3, Code: Echo Reply]
...
```

### Installation

You can install the CLI tool using `go install`:

```sh
go install github.com/code-brew-lab/pingo/cmd/pingo@v1.0.0
```

or you can compile it manually by using the Makefile. See Makefile for more details
```sh
make build
```

### Feedback
Your feedback is important to us. If you encounter any issues or have suggestions for improvement, please open an issue on our [GitHub repository](https://github.com/code-brew-lab/pingo/issues).

Thank you for using `pingo`!