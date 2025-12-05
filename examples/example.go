package main

import (
	"fmt"

	"github.com/BroadcastFacilityController/nmos-go-client"
)

func main() {
	endpoint := nmos.NewNMOSEndpoint("10.10.71.11", 80)
	fmt.Println(endpoint.GetBaseURL())
	fmt.Println(endpoint.GetSupportedAPIs())
	apis, _ := endpoint.GetSupportedAPIs()
	for _, api := range apis {
		fmt.Print(api + ": ")
		fmt.Println(endpoint.GetAPISupportedVersions(api))
	}
}
