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
	workersServiceRegistry      *clusterManagement.ServiceRegistry
	coordinatorsServiceRegistry *clusterManagement.ServiceRegistry
	port                        int
	webServer                   *networking.WebServer
}

func newOnElectionAction(workersServiceRegistry *clusterManagement.ServiceRegistry, coordinatorsServiceRegistry *clusterManagement.ServiceRegistry, port int) *OnElectionAction {
	ea := OnElectionAction{
		workersServiceRegistry:      workersServiceRegistry,
		coordinatorsServiceRegistry: coordinatorsServiceRegistry,
		port:                        port,
	}

	return &ea
}

func (ea *OnElectionAction) OnElectedToBeLeader() {
	ea.workersServiceRegistry.UnregisterFromCluster()
	ea.workersServiceRegistry.RegisterForUpdates()

	if ea.webServer != nil {
		ea.webServer.Stop()
	}

	searchCoordinator := search.NewSearchCoordinator(ea.workersServiceRegistry, networking.NewWebClient())
	ea.webServer = networking.NewWebServer(ea.port, searchCoordinator)
	go ea.webServer.StartServer()

	ipAddress := getLocalIpAddress()
	currentServerAddress := fmt.Sprintf("http://%v:%v%v", ipAddress, ea.port, searchCoordinator.GetEndpoint())
	ea.coordinatorsServiceRegistry.RegisterToCluster(currentServerAddress)
}

func (ea *OnElectionAction) OnWorker() {
	searchWorker := search.NewSearchWorker()
	ea.webServer = networking.NewWebServer(ea.port, searchWorker)
	go ea.webServer.StartServer()

	ipAddress := getLocalIpAddress()
	currentServerAddress := fmt.Sprintf("http://%v:%v%v", ipAddress, ea.port, searchWorker.GetEndpoint())
	ea.workersServiceRegistry.RegisterToCluster(currentServerAddress)
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
