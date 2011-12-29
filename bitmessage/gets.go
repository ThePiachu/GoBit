// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

import(
	"mymath"
)

//TODO: test

type GetBlocks struct{
	Version [4]byte
	StartCount mymath.VarInt//var_int
	BlockLocatorHashes []Hash
	HashStop Hash
}

func (gb *GetBlocks)SetVersion(newVersion uint32){
	verchars:=mymath.Uint322HexRev(newVersion)
	
	copy(gb.Version[:], verchars)
}

func (gb *GetBlocks)AddHash(newhash Hash){
	gb.StartCount++
	gb.BlockLocatorHashes = append(gb.BlockLocatorHashes, newhash)
}

func (gb *GetBlocks)Clear(){
	gb.StartCount=0
	gb.BlockLocatorHashes=nil
}

func (gb *GetBlocks)SetStopHash(stop Hash){
	gb.HashStop=stop
}

//TODO: test
func (gb *GetBlocks)Compile()[]byte{

	sc:=mymath.VarInt2HexRev(gb.StartCount)//TODO: double check if it is Rev or not
	answer:=make([]byte, len(gb.Version)+len(sc)+32*len(gb.BlockLocatorHashes)+len(gb.HashStop))
	
	iterator:=0
	copy(answer[iterator:], gb.Version[:])
	iterator+=len(gb.Version)
	copy(answer[iterator:], sc)
	iterator+=len(sc)
	for i:=0;i<len(gb.BlockLocatorHashes);i++{
		tmp:=gb.BlockLocatorHashes[i]
		copy(answer[iterator:], tmp[:])
		iterator+=len(tmp)
	}
	
	copy(answer[iterator:], gb.HashStop[:])
	
	return answer
}







type GetHeaders struct{
	StartCount mymath.VarInt//var_int
	HashStart Hash
	HashStop Hash
}

func (gh *GetHeaders)SetStartCount(start mymath.VarInt){
	gh.StartCount=start
}

func (gh *GetHeaders)SetStartHash(start Hash){
	gh.HashStart=start
}

func (gh *GetHeaders)SetStopHash(stop Hash){
	gh.HashStop=stop
}

//TODO: test
func (gh *GetHeaders)Compile()[]byte{
	sc:=mymath.VarInt2HexRev(gh.StartCount)//TODO: double check if it is Rev or not
	answer:=make([]byte, len(sc)+len(gh.HashStart)+len(gh.HashStop))
	
	iterator:=0
	copy(answer[iterator:], sc)
	iterator+=len(sc)
	copy(answer[iterator:], gh.HashStart[:])
	iterator+=len(gh.HashStart)
	copy(answer[iterator:], gh.HashStop[:])
	
	return answer
}