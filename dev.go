package main

import (
	"fmt"

	"github.com/hyperledgendary/smart-asset-tx/topology"

	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/cmd"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/cmd/network"
)

// development and test only starter to bootstrap a FSC and Fabric network
// for secure asset transfer

func main() {

	fmt.Println("Smart Asset Transfer Development Starter")

	m := cmd.NewMain("smart-atx", "0.1")
	mainCmd := m.Cmd()
	mainCmd.AddCommand(network.NewCmd(topology.Topology()...))
	m.Execute()
}
