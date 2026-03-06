package main

import (
	"fmt"

	"github.com/BroadcastFacilityController/nmos-go-client"
	"github.com/BroadcastFacilityController/nmos-go-client/common"
)

func main() {
	doAPIs := make(map[common.APIType]bool)
	doAPIs[common.NODE] = false
	doAPIs[common.QUERY] = true
	doAPIs[common.REGISTRATION] = false
	doAPIs[common.CONNECTION] = false

	endpointHRefs := []string{
		"http://10.10.60.10:8080/x-nmos", // Cerebrum Registry
		//"http://10.10.71.11:80/x-nmos", // VB-440
		//"http://10.10.251.2:80/x-nmos", // Telestream SPG9000
	}
	endpoints := make([]nmos.NMOSEndpoint, len(endpointHRefs))
	for i, href := range endpointHRefs {
		endpoint, err := nmos.NewNMOSEndpoint(href)
		if err != nil {
			panic(err)
		}
		endpoints[i] = *endpoint
	}

	for i, endpoint := range endpoints {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("Endpoint HRef: ", endpointHRefs[i])
		fmt.Print("Endpoint APIs: ")
		apis, err := endpoint.GetSupportedAPIs()
		if err != nil {
			panic(err)
		}
		fmt.Println(apis)
		for _, api := range apis {
			doAPI, doAPI_ok := doAPIs[api]
			if !doAPI_ok || !doAPI {
				continue
			}
			fmt.Println("------------------------------")
			fmt.Println("API: ", api)
			switch api {
			case common.NODE:
				versions, err := endpoint.IS04().GetAPIVersions(api)
				if err != nil {
					panic(err)
				}
				fmt.Println("Versions: ", versions)
				for _, vers := range versions {
					fmt.Println("----------")
					fmt.Println("Version: ", vers)
					if vers.Equals(common.NewAPIVersion(1, 0)) {
						fmt.Println("Node:")
						fmt.Println(endpoint.IS04().V1_0().NodeGetSelf())
						fmt.Println("Devices:")
						fmt.Println(endpoint.IS04().V1_0().NodeGetDevices())
						fmt.Println("Flows:")
						fmt.Println(endpoint.IS04().V1_0().NodeGetFlows())
						fmt.Println("Sources:")
						fmt.Println(endpoint.IS04().V1_0().NodeGetSources())
						fmt.Println("Senders:")
						fmt.Println(endpoint.IS04().V1_0().NodeGetSenders())
						fmt.Println("Receivers:")
						fmt.Println(endpoint.IS04().V1_0().NodeGetReceivers())
					}
					if vers.Equals(common.NewAPIVersion(1, 1)) {
						fmt.Println("Nodes:")
						fmt.Println(endpoint.IS04().V1_1().NodeGetSelf())
						fmt.Println("Devices:")
						fmt.Println(endpoint.IS04().V1_1().NodeGetDevices())
						fmt.Println("Flows:")
						fmt.Println(endpoint.IS04().V1_1().NodeGetFlows())
						fmt.Println("Sources:")
						fmt.Println(endpoint.IS04().V1_1().NodeGetSources())
						fmt.Println("Senders:")
						fmt.Println(endpoint.IS04().V1_1().NodeGetSenders())
						fmt.Println("Receivers:")
						fmt.Println(endpoint.IS04().V1_1().NodeGetReceivers())
					}
					if vers.Equals(common.NewAPIVersion(1, 2)) {
						fmt.Println("Nodes:")
						fmt.Println(endpoint.IS04().V1_2().NodeGetSelf())
						fmt.Println("Devices:")
						fmt.Println(endpoint.IS04().V1_2().NodeGetDevices())
						fmt.Println("Flows:")
						fmt.Println(endpoint.IS04().V1_2().NodeGetFlows())
						fmt.Println("Sources:")
						fmt.Println(endpoint.IS04().V1_2().NodeGetSources())
						fmt.Println("Senders:")
						fmt.Println(endpoint.IS04().V1_2().NodeGetSenders())
						fmt.Println("Receivers:")
						fmt.Println(endpoint.IS04().V1_2().NodeGetReceivers())
					}
					if vers.Equals(common.NewAPIVersion(1, 3)) {
						fmt.Println("Nodes:")
						fmt.Println(endpoint.IS04().V1_3().NodeGetSelf())
						fmt.Println("Devices:")
						fmt.Println(endpoint.IS04().V1_3().NodeGetDevices())
						fmt.Println("Flows:")
						fmt.Println(endpoint.IS04().V1_3().NodeGetFlows())
						fmt.Println("Sources:")
						fmt.Println(endpoint.IS04().V1_3().NodeGetSources())
						fmt.Println("Senders:")
						fmt.Println(endpoint.IS04().V1_3().NodeGetSenders())
						fmt.Println("Receivers:")
						fmt.Println(endpoint.IS04().V1_3().NodeGetReceivers())
					}
				}
			case common.QUERY:
				versions, err := endpoint.IS04().GetAPIVersions(api)
				if err != nil {
					panic(err)
				}
				fmt.Println("Versions: ", versions)
				for _, vers := range versions {
					fmt.Println("----------")
					fmt.Println("Version: ", vers)
					if vers.Equals(common.NewAPIVersion(1, 0)) {
						//fmt.Println("Nodes:")
						//fmt.Println(endpoint.IS04().V1_0().QueryGetNodes())
						//fmt.Println("Devices:")
						//fmt.Println(endpoint.IS04().V1_0().QueryGetDevices())
						//fmt.Println("Flows:")
						//fmt.Println(endpoint.IS04().V1_0().QueryGetFlows())
						//fmt.Println("Sources:")
						//fmt.Println(endpoint.IS04().V1_0().QueryGetSources())
						//fmt.Println("Senders:")
						//fmt.Println(endpoint.IS04().V1_0().QueryGetSenders())
						//fmt.Println("Receivers:")
						//fmt.Println(endpoint.IS04().V1_0().QueryGetReceivers())
					}
					if vers.Equals(common.NewAPIVersion(1, 1)) {
						//fmt.Println("Nodes:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetNodes())
						//fmt.Println("Devices:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetDevices())
						//fmt.Println("Flows:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetFlows())
						fmt.Println("Sources:")
						fmt.Println(endpoint.IS04().V1_1().QueryGetSources())
						//fmt.Println("Senders:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetSenders())
						//fmt.Println("Receivers:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetReceivers())
					}
					if vers.Equals(common.NewAPIVersion(1, 2)) {
						//fmt.Println("Nodes:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetNodes())
						//fmt.Println("Devices:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetDevices())
						//fmt.Println("Flows:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetFlows())
						//fmt.Println("Sources:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetSources())
						//fmt.Println("Senders:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetSenders())
						//fmt.Println("Receivers:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetReceivers())
					}
					if vers.Equals(common.NewAPIVersion(1, 3)) {
						fmt.Println("Nodes:")
						fmt.Println(endpoint.IS04().V1_3().QueryGetNodes())
						//fmt.Println("Devices:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetDevices())
						//fmt.Println("Flows:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetFlows())
						//fmt.Println("Sources:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetSources())
						//fmt.Println("Senders:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetSenders())
						//fmt.Println("Receivers:")
						//fmt.Println(endpoint.IS04().V1_2().QueryGetReceivers())
					}
				}
			case common.CONNECTION:
				versions, err := endpoint.IS05().GetAPIVersions(common.CONNECTION)
				if err != nil {
					panic(err)
				}
				fmt.Println("Versions: ", versions)
				for _, vers := range versions {
					fmt.Println("----------")
					fmt.Println("Version: ", vers)
					if vers.Equals(common.NewAPIVersion(1, 0)) {
						// IS-05 V1.0
						fmt.Println("Senders:")
						senders, err := endpoint.IS05().V1_0().ConnectionGetSenders()
						fmt.Println(senders, err)
						if len(senders) > 0 {
							fmt.Println("Sender 0 Constraints:")
							fmt.Println(endpoint.IS05().V1_0().ConnectionGetSenderConstraints(senders[0]))
							fmt.Println("Sender 0 Active:")
							fmt.Println(endpoint.IS05().V1_0().ConnectionGetSenderActive(senders[0]))
							fmt.Println("Sender 0 Staged:")
							fmt.Println(endpoint.IS05().V1_0().ConnectionGetSenderStaged(senders[0]))
							fmt.Println("Sender 0 Transport File:")
							fmt.Println(endpoint.IS05().V1_0().ConnectionGetSenderTransportFile(senders[0]))
						}
						fmt.Println("Receivers:")
						receivers, err := endpoint.IS05().V1_0().ConnectionGetReceivers()
						fmt.Println(receivers, err)
						if len(receivers) > 0 {
							fmt.Println("Receiver 0 Constraints:")
							fmt.Println(endpoint.IS05().V1_0().ConnectionGetReceiverConstraints(receivers[0]))
							fmt.Println("Receiver 0 Active:")
							fmt.Println(endpoint.IS05().V1_0().ConnectionGetReceiverActive(receivers[0]))
							fmt.Println("Receiver 0 Staged:")
							fmt.Println(endpoint.IS05().V1_0().ConnectionGetReceiverStaged(receivers[0]))
						}
					}
					if vers.Equals(common.NewAPIVersion(1, 1)) {
						// IS-05 V1.0
						fmt.Println("Senders:")
						senders, err := endpoint.IS05().V1_1().ConnectionGetSenders()
						fmt.Println(senders, err)
						if len(senders) > 0 {
							fmt.Println("Sender 0 Constraints:")
							fmt.Println(endpoint.IS05().V1_1().ConnectionGetSenderConstraints(senders[0]))
							fmt.Println("Sender 0 Active:")
							fmt.Println(endpoint.IS05().V1_1().ConnectionGetSenderActive(senders[0]))
							fmt.Println("Sender 0 Staged:")
							fmt.Println(endpoint.IS05().V1_1().ConnectionGetSenderStaged(senders[0]))
							fmt.Println("Sender 0 Transport File:")
							fmt.Println(endpoint.IS05().V1_1().ConnectionGetSenderTransportFile(senders[0]))
							fmt.Println("Sender 0 Transport Type:")
							fmt.Println(endpoint.IS05().V1_1().ConnectionGetSenderTransportType(senders[0]))
						}
						fmt.Println("Receivers:")
						receivers, err := endpoint.IS05().V1_1().ConnectionGetReceivers()
						fmt.Println(receivers, err)
						if len(receivers) > 0 {
							fmt.Println("Receiver 0 Constraints:")
							fmt.Println(endpoint.IS05().V1_1().ConnectionGetReceiverConstraints(receivers[0]))
							fmt.Println("Receiver 0 Active:")
							fmt.Println(endpoint.IS05().V1_1().ConnectionGetReceiverActive(receivers[0]))
							fmt.Println("Receiver 0 Staged:")
							fmt.Println(endpoint.IS05().V1_1().ConnectionGetReceiverStaged(receivers[0]))
							fmt.Println("Receiver 0 Transport Type:")
							fmt.Println(endpoint.IS05().V1_1().ConnectionGetReceiverTransportType(receivers[0]))
						}
					}
				}
			}
		}
	}
}
