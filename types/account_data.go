package types

type AccountData struct {
	Address         string
	AccountMetadata AccountMetadata
	Actions         Actions
	AtomBalance     int
	Score           int
}

// action that an account executed via a tx
type Actions struct {
	Votes           []Vote
	Delegates       []Delegate
	Sends           []Send
	SubmitProposals []SubmitProposal
}

type Send struct {
	Receiver string
	Amount   int
}

type Vote struct {
	Option     string
	ProposalId int
}

type Delegate struct {
	DelegatedValidator string
	Amount             string
}

type SubmitProposal struct {
	ProposalId int
}

// metadata of an account
type AccountMetadata struct {
	NumTxs               int
	NumVotes             int
	AmountDelegated      int
	NumProposalsSubmited int
}
