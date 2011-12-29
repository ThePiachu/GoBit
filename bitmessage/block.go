// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

import(
	"mymath"
	"big"
)


//TODO: test
type Block struct{
	Version [4]byte
	PrevBlock [32]byte//char[32]
	MerkleRoot [32]byte//char[32]
	Timestamp [4]byte
	Bits [4]byte
	Nonce [4]byte
	TxnCount mymath.VarInt
	Txns []*Tx//tx[]
}

func (b *Block)GetVersion()[]byte{
	return b.Version[:]
}
func (b *Block)GetPrevBlock()[]byte{
	return b.PrevBlock[:]
}
func (b *Block)GetMerkleRoot()[]byte{
	return b.MerkleRoot[:]
}
func (b *Block)GetTimestamp()[]byte{
	return b.Timestamp[:]
}
func (b *Block)GetBits()[]byte{
	return b.Bits[:]
}
func (b *Block)GetNonce()[]byte{
	return b.Nonce[:]
}
func (b *Block)GetTxnCount()mymath.VarInt{
	return b.TxnCount
}
func (b *Block)GetTxns()[]*Tx{
	return b.Txns
}


func (b *Block)SetVersion(ver uint32){
	answer:=mymath.Uint322HexRev(ver)
	copy(b.Version[:], answer)
}

func (b *Block)SetPrevBlockStr(prev string){
	if len(prev)==32{
		answer:=mymath.String2Hex(prev)
		copy(b.PrevBlock[:], answer)
	}
}

func (b *Block)SetPrevBlockStrRev(prev string){
	if len(prev)==32{
		answer:=mymath.String2Hex(prev)
		copy(b.PrevBlock[:], mymath.Rev(answer))
	}
}

func (b *Block)SetPrevBlock(prev []byte){
	if len(prev)==32{
		copy(b.PrevBlock[:], prev)
	}
}

func (b *Block)SetMerkleRootStr(merkle string){
	if len(merkle)==64{
		answer:=mymath.String2Hex(merkle)
		copy(b.MerkleRoot[:], answer)
	}
}

func (b *Block)SetMerkleRootStrRev(merkle string){
	if len(merkle)==64{
		answer:=mymath.String2Hex(merkle)
		copy(b.MerkleRoot[:], mymath.Rev(answer))
	}
}

func (b *Block)SetMerkleRoot(merkle []byte){
	if len(merkle)==32{
		copy(b.MerkleRoot[:], merkle)
	}
}

func (b *Block)SetTimestamp(time uint32){
	answer:=mymath.Uint322HexRev(time)
	copy(b.Timestamp[:], answer)
}

func (b *Block)SetBits(bi uint32){
	answer:=mymath.Uint322HexRev(bi)
	copy(b.Bits[:], answer)
}

func (b *Block)SetNonce(non uint32){
	answer:=mymath.Uint322HexRev(non)
	copy(b.Nonce[:], answer)
}

func (b *Block)AddTx(t *Tx){
	b.TxnCount++
	b.Txns=append(b.Txns, t)
}

func (b *Block)Clear(){
	b.TxnCount=0
	b.Txns=nil
}

func (b *Block)GetHeader()*BlockHeader{
	header:=new(BlockHeader)
	
	copy(header.Version[:], b.Version[:])
	copy(header.PrevBlock[:], b.PrevBlock[:])
	copy(header.MerkleRoot[:], b.MerkleRoot[:])
	copy(header.Timestamp[:], b.Timestamp[:])
	copy(header.Bits[:], b.Bits[:])
	copy(header.Nonce[:], b.Nonce[:])
	//TxnCount is zero
	
	return header
}
func (b *Block)Hash()[]byte{
	answer:=make([]byte, len(b.Version)+len(b.PrevBlock)+len(b.MerkleRoot)+len(b.Timestamp)+len(b.Bits)+len(b.Nonce))
	
	iterator:=0
	copy(answer[iterator:], b.Version[:])
	iterator+=len(b.Version)
	copy(answer[iterator:], b.PrevBlock[:])
	iterator+=len(b.PrevBlock)
	copy(answer[iterator:], b.MerkleRoot[:])
	iterator+=len(b.MerkleRoot)
	copy(answer[iterator:], b.Timestamp[:])
	iterator+=len(b.Timestamp)
	copy(answer[iterator:], b.Bits[:])
	iterator+=len(b.Bits)
	copy(answer[iterator:], b.Nonce[:])
	
	return mymath.Rev(mymath.DoubleSHA(answer))
}

func (b *Block)Compile()[]byte{
	vi:=mymath.VarInt2HexRev(b.TxnCount)//TODO: check if Rev or not
	
	totalLen:=len(b.Version)+len(b.PrevBlock)+len(b.MerkleRoot)+len(b.Timestamp)+len(b.Bits)+len(b.Nonce)+len(vi)
	
	for i:=0;i<len(b.Txns);i++{//TODO:check if the pointer is working
		totalLen+=b.Txns[i].Len()
	}
	
	answer:=make([]byte, totalLen)
	
	iterator:=0
	copy(answer[iterator:], b.Version[:])
	iterator+=len(b.Version)
	copy(answer[iterator:], b.PrevBlock[:])
	iterator+=len(b.PrevBlock)
	copy(answer[iterator:], b.MerkleRoot[:])
	iterator+=len(b.MerkleRoot)
	copy(answer[iterator:], b.Timestamp[:])
	iterator+=len(b.Timestamp)
	copy(answer[iterator:], b.Bits[:])
	iterator+=len(b.Bits)
	copy(answer[iterator:], b.Nonce[:])
	iterator+=len(b.Nonce)
	copy(answer[iterator:], vi[:])
	iterator+=len(vi)
	
	for i:=0;i<len(b.Txns);i++{//TODO:check if the pointer is working
		tmp:=b.Txns[i].Compile()[:]
		//log.Printf("%X - tx", tmp)
		copy(answer[iterator:], tmp)
		iterator+=len(tmp)
	}
	
	return answer

}

func (b *Block)Len()int{
	totalLen:=len(b.Version)+len(b.PrevBlock)+len(b.MerkleRoot)+len(b.Timestamp)+len(b.Bits)+len(b.Nonce)+b.TxnCount.Len()
	
	for i:=0;i<len(b.Txns);i++{//TODO:check if the pointer is working
		totalLen+=b.Txns[i].Len()
	}
	return totalLen
}

//TODO: check
func GenerateGenesisBlock() *Block{
	
	b:=new(Block)
	
	b.SetVersion(1)
	b.SetPrevBlockStrRev("0000000000000000000000000000000000000000000000000000000000000000")
	b.SetMerkleRootStrRev("4A5E1E4BAAB89F3A32518A88C31BC87F618F76673E2CC77AB2127B7AFDEDA33B")
	b.SetTimestamp(1231006505)
	b.SetBits(0x1d00ffff)
	b.SetNonce(2083236893)
	
	t:=new(Tx)
	
	t.SetVersion(1)
	
	tin:=new(TxIn)
	tout:=new(TxOut)
	
	var op OutPoint
	
	op.SetHash(NewHashFromBig(big.NewInt(0)))
	op.SetIndex(4294967295)//-1
	
	tin.SetOutPoint(op)
	tin.SetSignature(mymath.String2Hex("04ffff001d0104455468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73"))
	tin.SetSequence(0xFFFFFFFF)
	
	tout.SetValue(5000000000)
	tout.SetScriptStr("4104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5f", 0xAC)
	
	t.AddIn(tin)
	t.AddOut(tout)
	t.SetLockTime(0)
	
	b.AddTx(t)
	
	return b
}

//TODO: do
func DecodeBlock(blck []byte) *Block{
	blk:=new(Block)
	//log.Print("DecodeBlock")
	
	if len(blck)<4+32+32+4+4+4{
		return nil
	}
	copy(blk.Version[:], blck[0:4])
	copy(blk.PrevBlock[:], blck[4:4+32])
	copy(blk.MerkleRoot[:], blck[4+32:4+32+32])
	copy(blk.Timestamp[:], blck[4+32+32:4+32+32+4])
	copy(blk.Bits[:], blck[4+32+32+4:4+32+32+4+4])
	copy(blk.Nonce[:], blck[4+32+32+4+4:4+32+32+4+4+4])
	
	/*for i:=0;i<4;i++{
		blk.Version[i]=block[i]
	}
	for i:=0;i<32;i++{
		blk.PrevBlock[i]=block[4+i]
	}
	for i:=0;i<32;i++{
		blk.MerkleRoot[i]=block[4+32+i]
	}
	for i:=0;i<4;i++{
		blk.Timestamp[i]=block[4+32+32+i]
	}
	for i:=0;i<4;i++{
		blk.Bits[i]=block[4+32+32+4+i]
	}
	for i:=0;i<4;i++{
		blk.Nonce[i]=block[4+32+32+4+4+i]
	}*/
	//4+32+32+4+4+4=80
	var remainingBits []byte
	blk.TxnCount, remainingBits=mymath.DecodeVarIntGiveRest(blck[80:])
	//log.Printf("decode block remaining bits - %X", remainingBits)
	
	blk.Txns, remainingBits=DecodeMultipleTxs(remainingBits, int(blk.TxnCount))
	//log.Printf("decode block remaining bits 2 - %X", remainingBits)
	
	return blk
}