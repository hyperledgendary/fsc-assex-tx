# Ensure all properties are exported as shell env-vars
set export

# set the current directory, and the location of the test dats
CWDIR := justfile_directory()

_default:
  @just -f {{justfile()}} --list

# Build the code and pull down the idemix tool
build:
    go mod tidy
    go build -o bin/dev dev.go
    go get github.com/IBM/idemix/tools/idemixgen@v0.0.0-20220113150823-80dd4cb2d74e

# Generate the configuration based on the topology.go
gen:
    #!/bin/bash
    set -e -x -o pipefail
    if [ -d _cfg ]; then
      rm -rf _cfg
    else
      mkdir -p _cfg
    fi
    
    rm -rf cmd

    export FAB_BINS=/home/matthew/github.com/hyperledger/fabric/build/bin
    ./bin/dev network generate -p _cfg

# Start the topology as defined in topology.go
go:
    #!/bin/bash
    set -e -x -o pipefail
    export FAB_BINS=/home/matthew/github.com/hyperledger/fabric/build/bin
    ./bin/dev network start -p _cfg