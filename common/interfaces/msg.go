// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package interfaces

import (
	log "github.com/sirupsen/logrus"
)

/**************************
 * IRCD  Interface for Redeem Condition Datastructures (RCD)
 *
 * https://github.com/FactomProject/FactomDocs/blob/master/factomDataStructureDetails.md#factoid-transaction
 **************************/

type IMsg interface {
	Printable
	BinaryMarshallable

	// AckMatch() Returns true if this message must match an Ack to be processed.  Or if it is an Ack.
	AckMatch() bool

	// Type() Returns a byte indicating the type of message.
	Type() byte

	// IsLocal() A local message is never broadcast to the greater network.
	IsLocal() bool
	SetLocal(bool)

	// IsFullBroadcast() FullBroadcast means send to every node
	IsFullBroadcast() bool
	SetFullBroadcast(bool)

	// GetOrigin()
	// Returns the origin of this message; used to track
	// where a message came from. If int == -1, then this
	// FactomNode generated the message.
	GetOrigin() int
	SetOrigin(int)

	GetNetworkOrigin() string
	SetNetworkOrigin(string)

	// GetTimestamp() Timestamp
	// Returns the timestamp for a message
	GetTimestamp() Timestamp

	// GetRepeatHash() IHash
	// This is the hash used to check for repeated messages.  Almost always this
	// is the MsgHash, however for Chain Commits, Entry Commits, and Factoid Transactions,
	// this is the GetHash().
	GetRepeatHash() IHash

	//GetHash() IHash
	// Hash for this message as used by Consensus (i.e. what we match). Does not include
	// signatures to avoid Signature Maliation attacks.
	GetHash() IHash

	//GetMsgHash() IHash
	// Hash of this message.  Each message must be unique includes signatures
	GetMsgHash() IHash

	//GetFullMsgHash() IHash
	// Returns the full message hash of a message (includes signatures)
	GetFullMsgHash() IHash

	//IsPeer2Peer() bool
	// If this message should only reply to a peer, this is true.  If to
	// be broadcast, this should be false.  If the Origin is 0, then the
	// network can pick a peer to try.  If Origin is > 0, then the message
	// must go back to that peer (this message is a reply).
	IsPeer2Peer() bool
	SetPeer2Peer(bool)

	// Validate the message, given the state.  Three possible results:
	//  < 0 -- Message is invalid.  Discard
	//  0   -- Cannot tell if message is Valid
	//  1   -- Message is valid
	Validate(IState) int

	//ComputeVMIndex(IState)
	//Set the VMIndex for a message
	ComputeVMIndex(IState)

	//LeaderExecute(IState)
	// Call here if the server is a leader
	LeaderExecute(IState)

	//GetLeaderChainID() IHash
	// Debugging thing to track the leader responsible for a message ack.
	GetLeaderChainID() IHash
	SetLeaderChainID(IHash)

	//FollowerExecute(IState)
	// Call here if the server is a follower
	FollowerExecute(IState)

	//SendOut(IState, IMsg)
	// Send this message out over the NetworkOutQueue.  This is done with a method
	// to allow easier debugging and simulation.
	SendOut(IState, IMsg)

	//GetNoResend() bool
	// Some messages (DBState messages, missing data messages) must be explicitly sent.
	// We won't resend them or pass them on.
	GetNoResend() bool
	SetNoResend(bool)
	GetResendCnt() int

	// Process.  When we get a sequence of acknowledgements that we trust, we process.
	// A message will only be processed once, and in order, guaranteed.
	// Returns true if able to process, false if process is waiting on something.
	Process(dbheight uint32, state IState) bool

	//GetVMIndex() int
	// Some Messages need to be processed on certain VMs.  We set this and querry
	// the indexes of these machines here.
	GetVMIndex() int
	SetVMIndex(int)
	GetVMHash() []byte
	SetVMHash([]byte)
	GetMinute() byte
	SetMinute(byte)

	//MarkSentInvalid(bool)
	// Stall handling
	MarkSentInvalid(bool)
	SentInvalid() bool

	IsStalled() bool
	SetStall(bool)
	Expire(IState) bool

	//LogFields() log.Fields
	// Equivalent to String() for logging
	LogFields() log.Fields
}

// IMsgInternal
// Internal Messaging supporting Elections
type IMsgInternal interface {
	IMsg
	ProcessElections(IState, IElectionMsg)
}

type IMsgAck interface {
	IMsg
	GetDBHeight() uint32
}
