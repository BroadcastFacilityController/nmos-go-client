package main

import (
	"fmt"

	"github.com/BroadcastFacilityController/nmos-go-client"
	"github.com/BroadcastFacilityController/nmos-go-client/common"
)

func main() {
	endpoint, err := nmos.NewNMOSEndpoint("http://10.10.71.11:80/x-nmos")
	if err != nil {
		panic(err)
	}
	fmt.Println(endpoint.GetSupportedAPIs())
	apis, _ := endpoint.GetSupportedAPIs()
	fmt.Println(endpoint.IS04().GetAPIs())
	fmt.Println(endpoint.IS04().GetAPIVersions(common.NODE))
	for _, api := range apis {
		fmt.Println(api)
		switch api {
		case common.NODE:
			// IS-04 Node
			fmt.Print("Versions: ")
			fmt.Println(endpoint.IS04().GetAPIVersions(api))
			fmt.Print("V1.0: Node: ")
			fmt.Println(endpoint.IS04().V1_0().NodeGetSelf())
			fmt.Print("V1.0: Flows: ")
			fmt.Println(endpoint.IS04().V1_0().NodeGetFlows())
			fmt.Print("V1.0: Devices: ")
			fmt.Println(endpoint.IS04().V1_0().NodeGetDevices())
			fmt.Print("V1.0: Senders: ")
			fmt.Println(endpoint.IS04().V1_0().NodeGetSenders())
			fmt.Print("V1.1: Node: ")
			fmt.Println(endpoint.IS04().V1_1().NodeGetSelf())
			fmt.Print("V1.1: Flows: ")
			fmt.Println(endpoint.IS04().V1_1().NodeGetFlows())
			fmt.Print("V1.1: Devices: ")
			fmt.Println(endpoint.IS04().V1_1().NodeGetDevices())
			fmt.Print("V1.1: Senders: ")
			fmt.Println(endpoint.IS04().V1_1().NodeGetSenders())
			fmt.Print("V1.1: Receivers: ")
			fmt.Println(endpoint.IS04().V1_1().NodeGetReceivers())
			//fmt.Print("Receivers: ")
			//fmt.Println(endpoint.GetNodeAPI().GetV1_0().GetReceivers())
			// Check unpatch
			//fmt.Println(endpoint.GetNodeAPI().GetV1_0().PutReceiverSubscription("25a670fe-7d37-56e9-9e29-478654a8f275", nil))
			// Check patch
			//sender := node_v1_0.Sender{
			//	ID:           "9edc30a2-133f-591e-840c-cee8a289259b",
			//	FlowID:       "8a714a98-f688-5193-8db5-84c772cee526",
			//	DeviceID:     "ed400049-36ed-525e-933c-9c223eaebd81",
			//	Description:  "Telestream SPG9000/sender/Video1",
			//	Label:        "SPG-PRI SPG-PRI Video 1 sender",
			//	ManifestHRef: "http://10.10.251.2:80/x-manifest/senders/9edc30a2-133f-591e-840c-cee8a289259b/manifest",
			//	Tags: map[string][]string{
			//		"urn:x-nmos:tag:grouphint/v1.0": {"SPG9000:sender Video1"},
			//	},
			//	Transport: node_v1_0.RTP,
			//	Version:   "1761135285:780319076",
			//}
			//fmt.Println(endpoint.GetNodeAPI().GetV1_0().PutReceiverSubscription("25a670fe-7d37-56e9-9e29-478654a8f275", &sender))
		}
	}
}
