package main

import (
	"dspetrov/distributed-search/clusterManagement"
	"dspetrov/distributed-search/networking"
	"dspetrov/distributed-search/search"
	"fmt"
	"net"
	"strings"
)

type OnElectionAction struct {
	serviceRegistry *clusterManagement.ServiceRegistry
	port            int
	webServer       networking.WebServer
}

func newOnElectionAction(serviceRegistry *clusterManagement.ServiceRegistry, port int) OnElectionAction {
	ea := OnElectionAction{
		serviceRegistry: serviceRegistry,
		port:            port,
	}

	return ea
}

func (ea *OnElectionAction) OnElectedToBeLeader() {
	ea.serviceRegistry.UnregisterFromCluster()
	ea.serviceRegistry.RegisterForUpdates()
}

func (ea *OnElectionAction) OnWorker() {
	searchWorker := search.SearchWorker{}
	ea.webServer = networking.NewWebServer(ea.port, searchWorker)
	ea.webServer.StartServer()

	ipAddress := getLocalIpAddress()
	currentServerAddress := fmt.Sprintf("http://%v:%v%v", ipAddress, ea.port, searchWorker.GetEndpoint())
	ea.serviceRegistry.RegisterToCluster(currentServerAddress)
}

func getLocalIpAddress() string {
	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {
		if interf.Name == "Wi-Fi" {
			if addrs, err := interf.Addrs(); err == nil {
				for _, addr := range addrs {
					addrStr := addr.String()
					if len(addrStr) <= 20 {
						return addrStr[:strings.IndexByte(addrStr, '/')]
					}
				}
			}
		}
	}

	return ""
}
