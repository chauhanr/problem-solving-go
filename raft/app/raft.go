package app

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type CMState int

const DebugCM = 1

const (
	Follower CMState = iota
	Leader
	Candidate
	Dead
)

func (s CMState) String() string {
	switch s {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	case Dead:
		return "Dead"
	default:
		panic("unreachable")

	}
}

type ConsensusModule struct {
	mu                 sync.Mutex
	id                 int
	peerIds            []int
	currentTerm        int
	log                []LogEntry
	commitIndex        int
	state              CMState
	electionResetEvent time.Time
}

type LogEntry struct {
	Command interface{}
	Term    int
}

func (cm *ConsensusModule) dlog(format string, args ...interface{}) {
	if DebugCM > 0 {
		format = fmt.Sprintf("[%d] ", cm.id) + format
		log.Printf(format, args...)
	}
}

func (cm *ConsensusModule) runElectionTimer() {
	timeDuration := cm.electionTimeout()
	cm.mu.Lock()
	termStarted := cm.currentTerm
	cm.mu.Unlock()

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for {
		<-ticker.C
		cm.mu.Lock()
		if cm.state != Candidate && cm.state != Follower {
			cm.dlog("in election timer state: %s, bailing out", cm.state)
			cm.mu.Unlock()
			return
		}

		if termStarted != cm.currentTerm {
			cm.dlog("in election timer term changed for %d to %d, bailing out", termStarted, cm.currentTerm)
			cm.mu.Unlock()
			return
		}

		if elapsed := time.Since(cm.electionResetEvent); elapsed >= timeDuration {
			// start election
			cm.startElection()
			cm.mu.Unlock()
			return
		}
		cm.mu.Unlock()
	}

}

func (cm *ConsensusModule) startElection() {

}

func (cm *ConsensusModule) electionTimeout() time.Duration {
	min := 150
	max := 300

	t := int64(rand.Intn(max-min) + min)
	return time.Duration(t * int64(time.Millisecond))
}
