package main

import (
	"fmt"
	"hyperledgendary/smart-asset-tx/v2/topology"

	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/cmd"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/cmd/network"
	view "github.com/hyperledger-labs/fabric-smart-client/platform/view/services/client/view/cmd"
)

// development and test only starter to bootstrap a FSC and Fabric network
// for secure asset transfer

func main() {

	fmt.Println("Smart Asset Transfer Development Starter")

	m := cmd.NewMain("smart-atx", "0.1")
	mainCmd := m.Cmd()
	mainCmd.AddCommand(network.NewCmd(topology.Topology()...))
	mainCmd.AddCommand(view.NewCmd())
	m.Execute()
}
