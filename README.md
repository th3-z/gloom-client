Gloom client
------------


### Development

You'll need Golang 1.13+ installed before following these steps.

#### Setup

* `git clone https://github.com/th3-z/gloom-client.git`
* `cd gloom-client`
* Install deps (Example is Ubuntu 18.04)
    - `sudo apt-get -y install build-essential libglu1-mesa-dev libpulse-dev libglib2.0-dev`
* Install Qt binding
    - `export GO111MODULE=off; go get -v github.com/therecipe/qt/cmd/... && $(go env GOPATH)/bin/qtsetup test && $(go env GOPATH)/bin/qtsetup -test=false`
*

#### Building

* `$(GOPATH)/qtdeploy build desktop .`
* `./deploy/linux/gloom-client`
