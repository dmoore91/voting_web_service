package candidate

// swagger:model candidate
type Candidate struct {
	CandidateID int    `json:"candidate_id"`
	Username    string `json:"username"`
	Party       string `json:"party"`
}

// swagger:model candidateList
type CandidateList struct {
	Candidates []Candidate `json:"candidates"`
}
