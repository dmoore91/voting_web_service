package candidate

// swagger:model candidate
type Candidate struct {
	CandidateID int    `json:"candidate_id"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
}

// swagger:model candidateList
type CandidateList struct {
	Candidates []Candidate `json:"candidates"`
}
