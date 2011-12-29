// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

import (
	"mymath"
	"time"
	"net"
)

const (
	NOT_NODE_NETWORK uint64 = iota //==0
	NODE_NETWORK uint64 = iota //==1
	)

//TODO: double checks, but it's probably OK
type VersionMessage struct{
	Version [4]byte
	Services [8]byte
	Timestamp [8]byte
	AddrYou [26]byte
	AddrMe [26]byte
	Nonce [8]byte
	SubVersionNum []byte
	StartHeight [4]byte
}

func (vm *VersionMessage)SetVersion(ver uint32){
	answer:=mymath.Uint322HexRev(ver)
	for i:=0;i<4;i++{
		vm.Version[i]=answer[i]
	}
}

func (vm *VersionMessage)SetServices(ser uint64){
	answer:=mymath.Uint642HexRev(ser)
	for i:=0;i<8;i++{
		vm.Services[i]=answer[i]
	}
}

func (vm *VersionMessage)SetTimestamp(setTime uint64){
	answer:=mymath.Uint642HexRev(setTime)
	for i:=0;i<8;i++{
		vm.Timestamp[i]=answer[i]
	}
}

func (vm *VersionMessage)SetTimestampNow(){
	answer:=mymath.Uint642HexRev(uint64(time.Seconds()))
	for i:=0;i<8;i++{
		vm.Timestamp[i]=answer[i]
	}
}

func (vm *VersionMessage)SetAddrYou(ip net.IP, ser uint64, port uint16){
	na:=new (NetworkAddress)
	na.SetServices(ser)
	na.SetIP(ip)
	na.SetPort(port)
	
	answer:=na.CompileForVersion()
	
	for i:=0;i<26;i++{
		vm.AddrYou[i]=answer[i]
	}
}

func (vm *VersionMessage)SetAddrMe(ip net.IP, ser uint64, port uint16){
	na:=new (NetworkAddress)
	na.SetServices(ser)
	na.SetIP(ip)
	na.SetPort(port)
	
	answer:=na.CompileForVersion()
	
	for i:=0;i<26;i++{
		vm.AddrMe[i]=answer[i]
	}
}

func (vm *VersionMessage)SetNonce(n uint64){
	answer:=mymath.Uint642HexRev(n)
	for i:=0;i<8;i++{
		vm.Nonce[i]=answer[i]
	}
}

func (vm *VersionMessage)SetRandomNonce(){
	answer:=mymath.Randuint64()
	for i:=0;i<8;i++{
		vm.Nonce[i]=answer[i]
	}
}

func (vm *VersionMessage)SetSubVersionNull(){
	vm.SubVersionNum=make([]byte, 1)
}

func (vm *VersionMessage)SetStartHeight(block uint32){
	answer:=mymath.Uint322HexRev(block)
	for i:=0;i<4;i++{
		vm.StartHeight[i]=answer[i]
	}
}

func (vm *VersionMessage)Compile()[]byte{
	
	answer:=make([]byte, len(vm.Version)+len(vm.Services)+len(vm.Timestamp)+len(vm.AddrYou)+len(vm.AddrMe)+len(vm.Nonce)+len(vm.SubVersionNum)+len(vm.StartHeight))
	
	iterator:=0
	copy(answer[iterator:], vm.Version[:])
	iterator+=len(vm.Version)
	copy(answer[iterator:], vm.Services[:])
	iterator+=len(vm.Services)
	copy(answer[iterator:], vm.Timestamp[:])
	iterator+=len(vm.Timestamp)
	copy(answer[iterator:], vm.AddrYou[:])
	iterator+=len(vm.AddrYou)
	copy(answer[iterator:], vm.AddrMe[:])
	iterator+=len(vm.AddrMe)
	copy(answer[iterator:], vm.Nonce[:])
	iterator+=len(vm.Nonce)
	copy(answer[iterator:], vm.SubVersionNum[:])
	iterator+=len(vm.SubVersionNum)
	copy(answer[iterator:], vm.StartHeight[:])
	
	return answer
}	