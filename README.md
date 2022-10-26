# Secure Smart Asset Transfer

## Development

just reciepes available for all these steps

```
Available recipes:
    build # Build the code and pull down the idemix tool
    gen   # Generate the configuration based on the topology.go
    go    # Start the topology as defined in topology.go
```

Note the justfile is configured to use the `FAB_BINS` environment variable to specify exact which Fabric binaries
should be used. All the binaries must be present.

Pull dependencies and build the 'dev' launcher and the client application (`just build`)
```
    go mod tidy
    go build -o bin/dev dev.go
    go build -o bin/client clientmain/clientmain.go
```

Create the configuration material in `_cfg`  (`just gen`)

```
go get github.com/IBM/idemix/tools/idemixgen@v0.0.0-20220113150823-80dd4cb2d74e
./bin/dev network generate -p _cfg
```

Start the network
```
# if needed
export FAB_BINS=......  
./bin/dev network start -p _cfg
```

