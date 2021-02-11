package clusterManagement

type OnElectionCallback interface {
	OnElectedToBeLeader()
	OnWorker()
}
