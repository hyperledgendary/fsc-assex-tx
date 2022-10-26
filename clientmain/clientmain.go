package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledgendary/smart-asset-tx/states"
	"github.com/hyperledgendary/smart-asset-tx/views"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/core/config"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/core/id"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/client/web"
)

func main() {

	configProvider, err := config.NewProvider("/home/matthew/_workspace/src/github.com/hyperledgendary/smart-asset-tx/_cfg/fsc/nodes/issuer")
	must(err)

	config := &web.Config{}
	if configProvider.GetBool("fsc.web.tls.enabled") {
		config.URL = fmt.Sprintf("https://%s", configProvider.GetString("fsc.web.address"))
		// use the first ca in the list
		tlsRootCertFile := configProvider.GetStringSlice("fsc.web.tls.clientRootCAs.files")[0]
		if len(tlsRootCertFile) == 0 {
			panic(errors.New("web configuration must have 'fsc.web.tls.clientRootCAs.files' key defined"))
		}
		config.CACert = configProvider.TranslatePath(tlsRootCertFile)
		config.TLSCert = configProvider.GetPath("fsc.web.tls.cert.file")
		config.TLSKey = configProvider.GetPath("fsc.web.tls.key.file")
	} else {
		config.URL = fmt.Sprintf("http://%s", configProvider.GetString("fsc.web.address"))
	}

	webClient, err := web.NewClient(config)
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

	fmt.Println("Creating the JSON request")
	payload, err := json.Marshal(&views.Issue{
		Asset:     &asset,
		Recipient: aliceIdCert,
		Approver:  approverIdCert,
	})
	fmt.Println(string(payload))
	must(err)

	fmt.Println("\nCalling the issue view")
	res, err := webClient.CallView("issue", payload)
	must(err)
	fmt.Printf("Response is: \n%s\n", string(res.([]byte)))

	fmt.Println("done")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func dumpByteSlice(b []byte) {
	var a [16]byte
	n := (len(b) + 15) &^ 15
	for i := 0; i < n; i++ {
		if i%16 == 0 {
			fmt.Printf("%4d", i)
		}
		if i%8 == 0 {
			fmt.Print(" ")
		}
		if i < len(b) {
			fmt.Printf(" %02X", b[i])
		} else {
			fmt.Print("   ")
		}
		if i >= len(b) {
			a[i%16] = ' '
		} else if b[i] < 32 || b[i] > 126 {
			a[i%16] = '.'
		} else {
			a[i%16] = b[i]
		}
		if i%16 == 15 {
			fmt.Printf("  %s\n", string(a[:]))
		}
	}
}
