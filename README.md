# nevern@divultion tg/ds
### In Short
Nevern is a multi-client, multi-threaded reverse shell written in Go. We are aiming for backcompitability with all existing NetCat reverse shells.

Disclaimer: This reverse shell should only be used in the lawful, remote administration of authorized systems. Accessing a computer network without authorization or permission is illegal. (Keep yuping ofc y need that to develop your botnet, if so consider giving me some `donations` PLS PLS PLS PLS PLS 🥲)

## How Does It Work?
`nevern@divultion` server consists of 2 main components:
- Service & Reverse shell
Separate Service which handles all connections
- Cli
Tool for you to easily view and access those connections

## Run
REMEBER TO SET CONSTS IN `server.go` AND `client.py` FILES IF YOURE USING THEM
### Server
Use this command to run reverse shell service (Keep it running)
`nevern --host=192.168.0.106 --p=5030`
Use this command to run cli & get everything you need from shell
`nevern-cli help`
### Client
Litterally use any reverse shell client on the internet, this project is aimed for backcompitability with all existing NetCat reverse shells.
But for demonstation purposes we've added client.py so you can use that by running it
`python client.py`

## Manually building from source
- Go tools 1.19+ Required - (https://golang.org/doc/install)
- Clone this repo (`git clone https://github.com/divultion/nevern`)
### CLI
`cd cli`
`go build .`
### Service
`cd service`
`go build .`

### Donations
BTC: `bc1qhfznv6y2anu88e2p2ysfvzryuz5y9vcw22u75u`
ETH: `0xeAf0CafBB9eCde4Dc856f604F75195c8869Ae98A`
LTC: `ltc1q8n3udq3r9rfx63x48gxat3wsyf843q7w6yewqr`
DOGE: `DF2qchigxP7Lm8FCsrHKnT8odwAHf3S61j`
SOL: `EZ5J7Jy5mGXQxsYFSMdFzGXASRGm49qeFWvWfHCjYV9D`
XRP: `rh1ydmzXb8WKk9Y3hu7jVkWSGG8yTYVZ1F`
