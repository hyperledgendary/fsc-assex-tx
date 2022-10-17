package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledgendary/smart-asset-tx/states"
	"github.com/hyperledgendary/smart-asset-tx/views"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/core/id"

	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/client"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/client/web"
)

func main() {
	webClientConfig, err := client.NewWebClientConfigFromFSC("/home/matthew/_workspace/src/github.com/hyperledgendary/smart-asset-tx/_cfg/fsc/nodes/issuer")
	must(err)

	webClient, err := web.NewClient(webClientConfig)
	must(err)

	aliceIdCert, _ := id.LoadIdentity("/home/matthew/_workspace/src/github.com/hyperledgendary/smart-asset-tx/_cfg/fsc/crypto/peerOrganizations/fsc.example.com/peers/alice.fsc.example.com/msp/signcerts/alice.fsc.example.com-cert.pem")
	// issuerId, _ := ioutil.ReadFile("/home/matthew/_workspace/src/github.com/hyperledgendary/smart-asset-tx/_cfg/fsc/crypto/peerOrganizations/fsc.example.com/peers/issuer.fsc.example.com/msp/signcerts/issuer.fsc.example.com-cert.pem")
	approverIdCert, _ := id.LoadIdentity("/home/matthew/_workspace/src/github.com/hyperledgendary/smart-asset-tx/_cfg/fsc/crypto/peerOrganizations/fsc.example.com/peers/approver.fsc.example.com/msp/signcerts/approver.fsc.example.com-cert.pem")

	asset := states.Asset{
		ObjectType:        "minifig",
		ID:                "aa001",
		Owner:             aliceIdCert,
		PublicDescription: "Batman Minifig",
		PrivateProperties: []byte{},
	}

	payload, err := json.Marshal(&views.Issue{
		Asset:     &asset,
		Recipient: aliceIdCert,
		Approver:  approverIdCert,
	})
	fmt.Println(string(payload))
	must(err)

	res, err := webClient.CallView("issue", payload)
	must(err)

	fmt.Printf("%v\n", res)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
