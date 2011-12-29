// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

//TODO: do
//https://en.bitcoin.it/wiki/Protocol_specification#submitorder
type SubmitOrder struct{
	hash [32]byte
	//wallet_entry CWalletTx
}