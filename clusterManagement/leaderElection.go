package clusterManagement

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-zookeeper/zk"
)

const ELECTION_NAMESPACE = "/election"

type LeaderElection struct {
	conn               *zk.Conn
	currentZnodeName   string
	onElectionCallback OnElectionCallback
}

func NewLeaderElection(conn *zk.Conn, onElectionCallback OnElectionCallback) LeaderElection {
	le := LeaderElection{
		conn:               conn,
		onElectionCallback: onElectionCallback,
	}

	return le
}

func (le *LeaderElection) VolunteerForLeadership() {
	znodePrefix := ELECTION_NAMESPACE + "/c_"
	znodeFullPath, err := le.conn.Create(znodePrefix, []byte{}, zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}

	fmt.Println("znode name", znodeFullPath)
	le.currentZnodeName = strings.ReplaceAll(znodeFullPath, ELECTION_NAMESPACE+"/", "")
}

func (le *LeaderElection) ReelectLeader() {
	var predecessorZnodeName string
	var ch <-chan zk.Event
	exists := false
	for !exists {
		children, _, err := le.conn.Children(ELECTION_NAMESPACE)
		if err != nil {
			panic(err)
		}

		sort.Strings(children)
		smallestChild := children[0]

		if smallestChild == le.currentZnodeName {
			fmt.Println("I am the leader")
			le.onElectionCallback.OnElectedToBeLeader()
			return
		} else {
			fmt.Println("I am not the leader")

			predecessorIndex := sort.SearchStrings(children, le.currentZnodeName) - 1
			predecessorZnodeName = children[predecessorIndex]

			exists, _, ch, err = le.conn.ExistsW(ELECTION_NAMESPACE + "/" + predecessorZnodeName)
			if err != nil {
				panic(err)
			}
		}
	}

	go le.processExistsEvent(ch)

	le.onElectionCallback.OnWorker()

	fmt.Println("Watching znode", predecessorZnodeName)
	fmt.Println()
}

func (le *LeaderElection) processExistsEvent(ch <-chan zk.Event) {
	event := <-ch
	if event.Type == zk.EventNodeDeleted {
		le.ReelectLeader()
	}
}
