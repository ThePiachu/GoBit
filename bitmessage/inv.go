// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

import(
	"mymath"
)

//TODO: change inventory to be able to handle everything?
//TODO: test

const (
	ERROR uint32 = iota //==0
	MSG_TX uint32 = iota //==1
	MSG_BLOCK uint32 = iota //==2
	)

type Inv struct{
	Count mymath.VarInt//var_int
	Inventory []*InventoryVector
}

//getdata responce to inv
type GetData Inv

//TODO: delete
/*func (i *Inv)AddNetworkAddress(na *NetworkAddress){
	if i.Count<1388{ //so payload length would be less than 50000 bytes
		i.Count++
		i.Inventory=append(i.Inventory, NewInventoryVectorFromNetworkAddress(na))
	}
}*/

//TODO: support adding blocks and transactions

func (i *Inv)Clear(){
	i.Count=0
	i.Inventory=nil
}

//TODO: double check if 36 is really the answer
func (i *Inv)Compile()[]byte{
	vi:=mymath.VarInt2HexRev(i.Count)//TODO: check if Rev or not

	answer:=make([]byte, len(vi)+36*len(i.Inventory))

	iterator:=0
	copy(answer[iterator:], vi)	
	iterator+=len(vi)
	for j:=0;j<len(i.Inventory);j++{
		copy(answer[iterator:], i.Inventory[j].Compile())	
		iterator+=34
	}

	return answer
}



//TODO: test
type InventoryVector struct{
	Vectortype uint32
	Hash [32]byte
}

//TODO: support adding blocks and transactions

//TODO: delete
/*
func NewInventoryVectorFromNetworkAddress(na *NetworkAddress) *InventoryVector{
	iv:=new(InventoryVector)
	//TODO: set proper type
	copy(iv.Hash[:], na.GetHash())
	
	return iv
}*/

func (iv *InventoryVector)SetType(newtype uint32){
	iv.Vectortype=newtype
}

func (iv *InventoryVector)SetHash(newhash []byte){
	if len(newhash)==32{
		copy(iv.Hash[:], newhash)
	}
}

func (iv *InventoryVector)Compile()[]byte{
	answer:=make([]byte, 4+len(iv.Hash))
	
	iterator:=0
	copy(answer[iterator:], mymath.Uint322HexRev(iv.Vectortype))//TODO: check endianess
	iterator+=4
	copy(answer[iterator:], iv.Hash[:])
	
	return answer
}	