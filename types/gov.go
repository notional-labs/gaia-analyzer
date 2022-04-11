package types

type Vote struct {
	Option     string
	ProposalId int
	Height     int64
	TxHash     string
}
