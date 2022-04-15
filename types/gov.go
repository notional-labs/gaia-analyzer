package types

type Vote struct {
	Code       int
	Option     string
	ProposalId int
	Height     int64
	TxHash     string
}
