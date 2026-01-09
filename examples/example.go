package main

import (
	"fmt"

	"github.com/BroadcastFacilityController/nmos-go-client"
	"github.com/BroadcastFacilityController/nmos-go-client/common"
)

func main() {

	endpointHRefs := []string{
		"http://10.10.60.10:8080/x-nmos",
		//"http://10.10.71.11:80/x-nmos",
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
						fmt.Println("Nodes:")
						fmt.Println(endpoint.IS04().V1_0().QueryGetNodes())
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
						fmt.Println("Nodes:")
						fmt.Println(endpoint.IS04().V1_1().QueryGetNodes())
						//fmt.Println("Devices:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetDevices())
						//fmt.Println("Flows:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetFlows())
						//fmt.Println("Sources:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetSources())
						//fmt.Println("Senders:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetSenders())
						//fmt.Println("Receivers:")
						//fmt.Println(endpoint.IS04().V1_1().QueryGetReceivers())
					}
					if vers.Equals(common.NewAPIVersion(1, 2)) {
						fmt.Println("Nodes:")
						fmt.Println(endpoint.IS04().V1_2().QueryGetNodes())
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
			}
		}
	}
}
