// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

import (
	"mymath"
)

//TODO:test
type Headers struct{
	Count mymath.VarInt
	Headers []*BlockHeader//block_header[]
}

//TODO:check if there is any limit on Count
func (hs *Headers)AddHeader(bh *BlockHeader){
	hs.Count++
	hs.Headers=append(hs.Headers, bh)
}

func (hs *Headers)Clear(){
	hs.Count=0
	hs.Headers=nil
}

func (hs *Headers)Compile()[]byte{

	vi:=mymath.VarInt2HexRev(hs.Count)//TODO: check if Rev or not

	answer:=make([]byte, len(vi)+81*len(hs.Headers))//TODO: double check if it is 81

	iterator:=0
	copy(answer[iterator:], vi)	
	iterator+=len(vi)
	for i:=0;i<len(hs.Headers);i++{
		copy(answer[iterator:], hs.Headers[i].Compile())	
		iterator+=81
	}

	return answer
}




//TODO: test
type BlockHeader struct{
	Version [4]byte
	PrevBlock [32]byte//char[32]
	MerkleRoot [32]byte//char[32]
	Timestamp [4]byte
	Bits [4]byte
	Nonce [4]byte
	TxnCount [1]byte
}


func (bh *BlockHeader)GetVersion()[]byte{
	return bh.Version[:]
}
func (bh *BlockHeader)GetPrevBlock()[]byte{
	return bh.PrevBlock[:]
}
func (bh *BlockHeader)GetMerkleRoot()[]byte{
	return bh.MerkleRoot[:]
}
func (bh *BlockHeader)GetTimestamp()[]byte{
	return bh.Timestamp[:]
}
func (bh *BlockHeader)GetBits()[]byte{
	return bh.Bits[:]
}
func (bh *BlockHeader)GetNonce()[]byte{
	return bh.Nonce[:]
}
func (bh *BlockHeader)GetTxnCount()[]byte{
	return bh.TxnCount[:]
}

func (bh *BlockHeader)SetVersion(ver uint32){
	answer:=mymath.Uint322HexRev(ver)
	copy(bh.Version[:], answer[:])
}

func (bh *BlockHeader)SetPrevBlockStr(prev string){
	if len(prev)==32{
		answer:=mymath.String2Hex(prev)
		copy(bh.PrevBlock[:], answer)
	}
}

func (bh *BlockHeader)SetPrevBlock(prev []byte){
	if len(prev)==32{
		copy(bh.PrevBlock[:], prev)
	}
}

func (bh *BlockHeader)SetMerkleRootStr(merkle string){
	if len(merkle)==32{
		answer:=mymath.String2Hex(merkle)
		copy(bh.MerkleRoot[:], answer)
	}
}

func (bh *BlockHeader)SetMerkleRoot(merkle []byte){
	if len(merkle)==32{
		copy(bh.MerkleRoot[:], merkle)
	}
}

func (bh *BlockHeader)SetTimestamp(time uint32){
	answer:=mymath.Uint322HexRev(time)
	copy(bh.Timestamp[:], answer)
}

func (bh *BlockHeader)SetBits(bi uint32){
	answer:=mymath.Uint322HexRev(bi)
	copy(bh.Bits[:], answer)
}

func (bh *BlockHeader)SetNonce(non uint32){
	answer:=mymath.Uint322HexRev(non)
	copy(bh.Nonce[:], answer)
}

func (bh *BlockHeader)Compile()[]byte{
	answer:=make([]byte, len(bh.Version)+len(bh.PrevBlock)+len(bh.MerkleRoot)+len(bh.Timestamp)+len(bh.Bits)+len(bh.Nonce)+len(bh.TxnCount))
	
	iterator:=0
	copy(answer[iterator:], bh.Version[:])
	iterator+=len(bh.Version)
	copy(answer[iterator:], bh.PrevBlock[:])
	iterator+=len(bh.PrevBlock)
	copy(answer[iterator:], bh.MerkleRoot[:])
	iterator+=len(bh.MerkleRoot)
	copy(answer[iterator:], bh.Timestamp[:])
	iterator+=len(bh.Timestamp)
	copy(answer[iterator:], bh.Bits[:])
	iterator+=len(bh.Bits)
	copy(answer[iterator:], bh.Nonce[:])
	iterator+=len(bh.Nonce)
	copy(answer[iterator:], bh.TxnCount[:])
	
	return answer
}

func (bh *BlockHeader)Hash()[]byte{
	answer:=make([]byte, len(bh.Version)+len(bh.PrevBlock)+len(bh.MerkleRoot)+len(bh.Timestamp)+len(bh.Bits)+len(bh.Nonce))
	
	iterator:=0
	copy(answer[iterator:], bh.Version[:])
	iterator+=len(bh.Version)
	copy(answer[iterator:], bh.PrevBlock[:])
	iterator+=len(bh.PrevBlock)
	copy(answer[iterator:], bh.MerkleRoot[:])
	iterator+=len(bh.MerkleRoot)
	copy(answer[iterator:], bh.Timestamp[:])
	iterator+=len(bh.Timestamp)
	copy(answer[iterator:], bh.Bits[:])
	iterator+=len(bh.Bits)
	copy(answer[iterator:], bh.Nonce[:])
	
	return mymath.Rev(mymath.DoubleSHA(answer))
}

func (bh *BlockHeader)Len() int{
	return len(bh.Version)+len(bh.PrevBlock)+len(bh.MerkleRoot)+len(bh.Timestamp)+len(bh.Bits)+len(bh.Nonce)+len(bh.TxnCount)
}
