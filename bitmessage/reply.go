// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

import(
	"mymath"
)

const (
	SUCCESS uint32 = iota //==0 //The IP Transaction can proceed (checkorder), or has been accepted (submitorder)
	WALLET_ERROR uint32 = iota //==1 //AcceptWalletTransaction() failed
	DENIED uint32 = iota //==2 //IP Transactions are not accepted by this node
	)

//TODO: test
type Reply struct{
	Reply [4]byte
}

func (r *Reply)SetReply(rep uint32){
	answer:=mymath.Uint322HexRev(rep)//TODO: check if it is rev or not
	copy(r.Reply[:], answer[:])
}

func (r *Reply)Compile()[]byte{
	return r.Reply[:]
}

func (r *Reply)Len() int{
	return len(r.Reply)
}