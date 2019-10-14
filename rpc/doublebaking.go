// Copyright (c) 2019 KIDTSUNAMI
// Author: alex@kidtsunami.com

package rpc

// DoubleBakingOp represents a double_baking_evidence operation
type DoubleBakingOp struct {
	GenericOp
	BH1      BlockHeader             `json:"bh1"`
	BH2      BlockHeader             `json:"bh2"`
	Metadata *DoubleBakingOpMetadata `json:"metadata"`
}

// DoubleBakingOpMetadata represents a double_baking_evidence operation metadata
type DoubleBakingOpMetadata struct {
	BalanceUpdates BalanceUpdates `json:"balance_updates"`
}
