<p align="center">
    <img width=125 height=125 src="https://files.th3-z.xyz/standing/3378fac86d7e477dc1f3e3fb1659c6b3.png"/>
</p>

<h1 align="center">Gloom</h1>


[![Go Report Card](https://goreportcard.com/badge/github.com/th3-z/gloom-client)](https://goreportcard.com/report/github.com/th3-z/gloom-client) [![GitHub license](https://img.shields.io/github/license/th3-z/gloom-client)](https://github.com/th3-z/gloom-client/blob/master/LICENSE)

## About

Client for [Gloom](https://github.com/th3-z/gloom-server) file server.

## Development

You'll need Golang 1.13+ installed before following these steps. Insutructions are for Ubuntu 18.04, may differ for other operating systems.

### Installation

* `git clone https://github.com/th3-z/gloom-client.git`
* `cd gloom-client`
* Install dependencies _(Ubuntu 18.04)_
    - `sudo apt-get -y install build-essential libglu1-mesa-dev libpulse-dev libglib2.0-dev`
* Install Qt binding
    - `export GO111MODULE=off; go get -v github.com/therecipe/qt/cmd/... && $(go env GOPATH)/bin/qtsetup test && $(go env GOPATH)/bin/qtsetup -test=false`

### Build

Optionally, add `$(GOPATH)/bin/` to your `PATH` environment variable.

* `$(GOPATH)/bin/qtdeploy build desktop .`

### Run

Execute the built binary.

* `./deploy/linux/gloom-client`
