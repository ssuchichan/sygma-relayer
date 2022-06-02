package elector

import (
	"github.com/ChainSafe/chainbridge-core/tss/common"
	"github.com/libp2p/go-libp2p-core/peer"
)

type staticCoordinatorElector struct {
	sessionID string
}

func NewCoordinatorElector(sessionID string) CoordinatorElector {
	return &staticCoordinatorElector{sessionID: sessionID}
}

func (s *staticCoordinatorElector) Coordinator(peers peer.IDSlice) (peer.ID, error) {
	sortedPeers := common.SortPeersForSession(peers, s.sessionID)
	if len(sortedPeers) == 0 {
		return peer.ID(""), nil
	}
	return common.SortPeersForSession(peers, s.sessionID)[0].ID, nil
}
