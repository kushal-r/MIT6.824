package raft

type serverState int32

const (
	Leader serverState = iota
	Follower
	Candidate
)

func (state serverState) String() string {
	switch state {
	case Leader:
		return "Leader"
	case Follower:
		return "Follower"
	default:
		return "Candidate"
	}
}

type RequestVoteArgs struct {
	Term,
	CandidateId int
}

type RequestVoteReply struct {
	Server      int
	VoteGranted bool
	Term        int
	Err         bool
}
