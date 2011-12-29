// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

import (
	"mymath"
	"net"
	"time"
)

type NetworkAddress struct{
	Time [4]byte
	Services [8]byte
	IP [16]byte
	Port [2]byte
}

//TODO: do
func (na *NetworkAddress)GetHash()[]byte{

	return nil
}

func (na *NetworkAddress)SetTimestamp(setTime uint32){
	answer:=mymath.Uint322HexRev(setTime)
	for i:=0;i<4;i++{
		na.Time[i]=answer[i]
	}
}

func (na *NetworkAddress)SetTimestampNow(setTime uint32){
	answer:=mymath.Uint322HexRev(uint32(time.Seconds()))
	for i:=0;i<4;i++{
		na.Time[i]=answer[i]
	}
}

func (na *NetworkAddress)SetServices(ser uint64){
	answer:=mymath.Uint642HexRev(ser)
	for i:=0;i<8;i++{
		na.Services[i]=answer[i]
	}
}

func (na *NetworkAddress)SetIP(IP net.IP){
	answer:=IP.To16()
	for i:=0;i<16;i++{
		na.IP[i]=answer[i]
	}
}

func (na *NetworkAddress)SetPort(port uint16){
	answer:=mymath.Uint162Hex(port)
	for i:=0;i<2;i++{
		na.Port[i]=answer[i]
	}
}

func (na *NetworkAddress)Compile()[]byte{
	answer:=make([]byte, len(na.Time)+len(na.Services)+len(na.IP)+len(na.Port))
	
	iterator:=0
	copy(answer[iterator:], na.Time[:])
	iterator+=4
	copy(answer[iterator:], na.Services[:])
	iterator+=8
	copy(answer[iterator:], na.IP[:])
	iterator+=16
	copy(answer[iterator:], na.Port[:])
	
	return answer
}

func (na *NetworkAddress)CompileForVersion()[]byte{
	answer:=make([]byte, len(na.Services)+len(na.IP)+len(na.Port))
	
	iterator:=0
	copy(answer[iterator:], na.Services[:])
	iterator+=8
	copy(answer[iterator:], na.IP[:])
	iterator+=16
	copy(answer[iterator:], na.Port[:])
	
	return answer

}
