// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

import(
	"mymath"
	"big"
)

type Hash [32]byte

func (h Hash)Len()int{
	return 32
}

func NewHashFromBig(hash *big.Int) Hash{
	var h Hash
	h=mymath.Big2Hex32(hash)
	return h
}

func NewHashFromString(hash string) Hash{
	var h Hash
	h=mymath.String2Hex32(hash)
	return h
}

func NewHash(b []byte) Hash{
	var h Hash
	if len(b)==32{
		copy(h[:], b)
	}
	return h
}
