package qrtc

import (
	"github.com/pion/sdp"
)

// SDPType for sdp type
type SDPType int

const (
	// SDPTypeOffer indicates that a description MUST be treated as an SDP
	// offer.
	SDPTypeOffer SDPType = iota + 1

	// SDPTypePranswer indicates that a description MUST be treated as an
	// SDP answer, but not a final answer. A description used as an SDP
	// pranswer may be applied as a response to an SDP offer, or an update to
	// a previously sent SDP pranswer.
	SDPTypePranswer

	// SDPTypeAnswer indicates that a description MUST be treated as an SDP
	// final answer, and the offer-answer exchange MUST be considered complete.
	// A description used as an SDP answer may be applied as a response to an
	// SDP offer or as an update to a previously sent SDP pranswer.
	SDPTypeAnswer

	// SDPTypeRollback indicates that a description MUST be treated as
	// canceling the current SDP negotiation and moving the SDP offer and
	// answer back to what it was in the previous stable state. Note the
	// local or remote SDP descriptions in the previous stable state could be
	// null if there has not yet been a successful offer-answer negotiation.
	SDPTypeRollback
)

// SessionDescription for session description
type SessionDescription struct {
	Type SDPType `json:"type"`
	SDP  string  `json:"sdp"`

	// This will never be initialized by callers, internal use only
	parsed *sdp.SessionDescription
}
