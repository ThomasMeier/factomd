# Installation

You need to set up Go environment with golang 1.12 or higher (but not above the verion listed [here](https://github.com/FactomProject/factomd/blob/master/engine/factomParams.go#L140) ). You also need git.  See the [Install from source](https://github.com/FactomProject/FactomDocs/blob/master/installFromSourceDirections.md) directions for more details and wallet installation instructions.

### Install the dependency management program

First check if golang 1.12 or higher is installed.  some operating systems install older versions.

`go version` should return something like
`go version go1.12 linux/amd64`

Next install Glide, which gets the dependencies for factomd and places them in the `$GOPATH/src/github/FactomProject/factomd/vendor` directory.

`go get -u github.com/Masterminds/glide`

### Install factomd Full Node

```
cd $GOPATH/src/github.com/FactomProject/
git clone https://github.com/FactomProject/factomd $GOPATH/src/github.com/FactomProject/factomd

# clear the glide cache since it has been known to cause errors
# deleting the $GOPATH/src/github.com/FactomProject/factomd/vendor  directory may be useful too
glide cc
cd $GOPATH/src/github.com/FactomProject/factomd
# this command download the dependencies and sets them to the right version
glide install
# install factomd with either the install.sh script or:
go install -ldflags "-X github.com/FactomProject/factomd/engine.Build=`git rev-parse HEAD` -X github.com/FactomProject/factomd/engine.FactomdVersion=`cat VERSION`" -v
# you can optionally use a config file to run in a non-standard mode
# mkdir -p ~/.factom/m2/
# cp $GOPATH/src/github.com/FactomProject/factomd/factomd.conf ~/.factom/m2/

```





