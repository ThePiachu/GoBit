// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

//TODO: do
//TODO: figure out?
//https://en.bitcoin.it/wiki/Protocol_specification#checkorder
type CheckOrder struct{
/*
checkorder
This message is used for IP Transactions, to ask the peer if it accepts such transactions and allow it to look at the content of the order.
It contains a CWalletTx object
Payload:
Field Size	Description	Data type	Comments
Fields from CMerkleTx
?	hashBlock
?	vMerkleBranch
?	nIndex
Fields from CWalletTx
?	vtxPrev
?	mapValue
?	vOrderForm
?	fTimeReceivedIsTxTime
?	nTimeReceived
?	fFromMe
?	fSpent
*/
}