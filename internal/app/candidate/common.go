package candidate

// swagger:model candidate
type Candidate struct {
	CandidateID int `json:"candidate_id"`
	UserId      int `json:"user_id"`
	PartyId     int `json:"party_id"`
}

// swagger:model candidateList
type CandidateList struct {
	Candidates []Candidate `json:"candidates"`
}
