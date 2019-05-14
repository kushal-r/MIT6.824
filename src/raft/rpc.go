package raft

func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	reply.Err, reply.Server = false, rf.me
	if rf.currentTerm == args.Term && rf.votedFor == args.CandidateId {
		reply.VoteGranted, reply.Term = true, rf.currentTerm
		return
	}

	if rf.currentTerm > args.Term || // valid candidate
		(rf.currentTerm == args.Term && rf.votedFor != -1) { // the server has voted in this term
		reply.Term, reply.VoteGranted = rf.currentTerm, false
		return
	}

	if args.Term > rf.currentTerm {
		rf.currentTerm, rf.votedFor = args.Term, -1
		if rf.state != Follower { // once server becomes follower, it has to reset electionTimer
			rf.resetElectionTimer(newRandDuration(ElectionTimeout))
			rf.state = Follower
		}
	}

	rf.leaderId = -1
	reply.Term = args.Term
	reply.VoteGranted = true
	rf.votedFor = args.CandidateId
	rf.resetElectionTimer(newRandDuration(ElectionTimeout))
}
