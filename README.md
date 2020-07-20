Gloom client
------------

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

`$(GOPATH)/bin/qtdeploy build desktop .`

### Run

Execute the built binary.

`./deploy/linux/gloom-client`
