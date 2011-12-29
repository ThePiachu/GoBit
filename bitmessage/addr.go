// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage


import (

)

//checked, works as expected
type AddressList struct{
	Count uint8
	AddrList []*NetworkAddress
}

/*type Address struct{
	lastseen uint32
	address NetworkAddress
}*/

func (al *AddressList)AddAddress(na *NetworkAddress){
	if al.Count<29{ //so payload length would be less than 1000 bytes
		al.Count++
		al.AddrList=append(al.AddrList, na)
	}
}

func (al *AddressList)Clear(){
	al.Count=0
	al.AddrList=nil
}

//TODO: double check if 30 is really the answer
func (al *AddressList)Compile()[]byte{
	answer:=make([]byte, 1+30*len(al.AddrList))

	answer[0]=al.Count
	iterator:=1
	for i:=0;i<len(al.AddrList);i++{
		copy(answer[iterator:], al.AddrList[i].Compile())	
		iterator+=34
	}

	return answer
}
/*
func (a *Address)SetLastSeen(seen uint32){
	a.lastseen=seen
}

func (a *Address)SetLastSeenNow(){
	a.lastseen=uint32(time.Seconds())
}

func (a *Address)SetAddress(na NetworkAddress){
	a.address=na
}

func (a *Address)Compile()[]byte{
	addr:=a.address.Compile()
	answer:=make([]byte, 4+len(addr))
	copy(answer[0:], mymath.Uint322Hex(a.lastseen))
	copy(answer[4:], a.address.Compile())	
	return answer
}*/