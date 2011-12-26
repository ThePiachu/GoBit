// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mymath
//Package for handling common math and conversion operations used in the rest of the program.


import(
	"big"
	"encoding/hex"
	"crypto/rand"
	"bytes"
	
	"log"
	"strings"
	"math"
)

func Hex2Uint64(b []byte) uint64{
	var answer uint64
	answer=0
	if len(b)>0{
		maxBytes:=8
		if len(b)<8{
			maxBytes=len(b)
		}
		for i:=0;i<maxBytes;i++{
			answer*=256
			answer+=uint64(b[i])
		}
	}
	return answer
}

func Hex2Uint32(b []byte) uint32{
	var answer uint32
	answer=0
	if len(b)>0{
		maxBytes:=4
		if len(b)<4{
			maxBytes=len(b)
		}
		for i:=0;i<maxBytes;i++{
			answer*=256
			answer+=uint32(b[i])
		}
	}
	return answer
}


func HexRev2Uint64(b []byte) uint64{
	return Hex2Uint64(Rev(b))
}

func HexRev2Uint32(b []byte) uint32{
	return Hex2Uint32(Rev(b))
}





//TODO: use this function to reverse some of the below functions?
func Rev(b []byte)[]byte{
	answer:=make([]byte, len(b))
	for i:=0;i<len(b);i++{
		answer[i]=b[len(b)-1-i]
	}
	return answer
}

func Hex2Big(b []byte) *big.Int{
	answer:=big.NewInt(0)
	
	for i:=0;i<len(b);i++{
		answer.Lsh(answer, 8)
		answer.Add(answer, big.NewInt(int64(b[i])))
	}
	
	return answer
}

func HexRev2Big(rev []byte) *big.Int{
	
	b:=make([]byte, len(rev))
	
	for i:=0;i<len(rev);i++{
		b[len(rev)-i-1]=rev[i]
	}
	
	answer:=big.NewInt(0)
	
	for i:=0;i<len(b);i++{
		answer.Lsh(answer, 8)
		answer.Add(answer, big.NewInt(int64(b[i])))
	}
	
	return answer
}

//TODO: check
func Big2Hex32(b *big.Int) [32]byte{
	var answer [32]byte
	
	tmp:=b.Bytes();
	
	if len(tmp)<=32{
		for i:=0;i<len(tmp);i++{
			answer[31-i]=tmp[i]
		}
	}
	
	
	return answer
}

func String2Hex32(s string) [32]byte{
	var answer [32]byte
	if len(s)==64{
		copy(answer[:], String2Hex(s))
	}
	return answer
}

func String2Hex(s string) []byte{
	answer, _:=hex.DecodeString(s)
	return answer
}

func Str2Hex(s string) []byte{
	return String2Hex(s)
}

func String2HexRev(s string) []byte{
	answer, _:=hex.DecodeString(s)
	return Rev(answer)
}

func Str2HexRev(s string) []byte{
	return Rev(String2Hex(s))
}

func Hex2String(b []byte)string{
	return strings.ToUpper(hex.EncodeToString(b))
}

func Hex2Str(b []byte)string{
	return Hex2String(b)
}

func HexRev2String(b []byte)string{
	return Hex2String(Rev(b))
}

func HexRev2Str(b []byte)string{
	return Hex2String(Rev(b))
}

func Uint322Hex(ui uint32) []byte{
	answer:=make([]byte, 4)
	answer[3]=uint8(ui%256)
	ui/=256
	answer[2]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	ui/=256
	answer[0]=uint8(ui%256)
	
	return answer
}

func Uint322HexRev(ui uint32) []byte{
	answer:=make([]byte, 4)
	answer[0]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	ui/=256
	answer[2]=uint8(ui%256)
	ui/=256
	answer[3]=uint8(ui%256)
	
	return answer
}

func Uint642Hex(ui uint64) []byte{
	answer:=make([]byte, 8)
	answer[7]=uint8(ui%256)
	ui/=256
	answer[6]=uint8(ui%256)
	ui/=256
	answer[5]=uint8(ui%256)
	ui/=256
	answer[4]=uint8(ui%256)
	ui/=256
	answer[3]=uint8(ui%256)
	ui/=256
	answer[2]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	ui/=256
	answer[0]=uint8(ui%256)
	
	return answer
}

func Uint642HexRev(ui uint64) []byte{
	answer:=make([]byte, 8)
	answer[0]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	ui/=256
	answer[2]=uint8(ui%256)
	ui/=256
	answer[3]=uint8(ui%256)
	ui/=256
	answer[4]=uint8(ui%256)
	ui/=256
	answer[5]=uint8(ui%256)
	ui/=256
	answer[6]=uint8(ui%256)
	ui/=256
	answer[7]=uint8(ui%256)
	
	return answer
}

func Uint162Hex(ui uint16) []byte{
	answer:=make([]byte, 2)
	answer[1]=uint8(ui%256)
	ui/=256
	answer[0]=uint8(ui%256)
	
	return answer
}

func Uint162HexRev(ui uint16) []byte{
	answer:=make([]byte, 2)
	answer[0]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	
	return answer
}


//TODO: test
func Uint2Hex(ui uint) []byte{
	length:=int(math.Ceil(math.Log2(float64(ui))/8.0))
	answer:=make([]byte, length)
	tmp:=ui
	for i:=0;i<length;i++ {
		answer[length-1-i]=uint8(tmp%256)
		tmp=tmp/256
	}
	return answer
}

//TODO: test
func Int2Hex(i int) []byte{//for the Bitcoin Script
	var ui uint
	if i<0{
		ui=uint(-i)
	}else{
		ui=uint(i)
	}
	answer:=Uint2Hex(ui)
	
	if i<0{
		if answer[0]>0x7F{
			answer=append([]byte{0x01}, answer[:]...)
		}else{
			answer[0]+=0x80
		}
	}
	return answer
}

func Randuint64() []byte{
	uint64max:=big.NewInt(1)
	uint64max.Lsh(uint64max, 64)
	
	randnum, _:=rand.Int(rand.Reader, uint64max)
	
	random:=randnum.Bytes()
	answer:=make([]byte, 8)
	
	for i:=0;i<len(random);i++{
		answer[i]=random[i]
	}
	
	return answer
}

func Randuint64Rev() []byte{
	uint64max:=big.NewInt(1)
	uint64max.Lsh(uint64max, 64)
	
	randnum, _:=rand.Int(rand.Reader, uint64max)
	
	random:=randnum.Bytes()
	answer:=make([]byte, 8)
	
	for i:=0;i<len(random);i++{
		answer[len(answer)-1-i]=random[i]
	}
	
	return answer
}

func ConcatBytes(list ...[]byte) []byte{
	size:=0//size of the resulting concatenated list
	for i:=0; i<len(list);i++{
		size+=len(list[i])//counting the sizes of individual parts of the list
	}
	answer:= make([]byte, size)//creates the array for the answer
	
	iterator:=0//iterator to count the position in the answer array
	for i:=0;i<len(list);i++{
		copy(answer[iterator:], list[i])//copies the data into the answer array
	 	iterator+=len(list[i])
	}
	return answer//returns the result
}

func AddByte(one []byte, two byte) []byte{
	size:=len(one)+1//size of the resulting concatenated list

	answer:= make([]byte, size)//creates the array for the answer
	
	copy(answer[0:], one)//copies the data into the answer array
	answer[len(one)]=two
	return answer//returns the result
}

func Byte2String(b []byte) string{
	return bytes.NewBuffer(b).String()
}









//Testing

//TODO: do
func TestEverything() bool{
	TestEverythingBitmath()
	
	if(RevTest()==false){
		return false
	}
	
	log.Print("All tests okay!")
	return true
}

func RevTest() bool{
	one:=make([]byte, 3)
	two:=make([]byte, 3)
	one[0]=0xFE
	one[1]=0xA9
	one[2]=0x01
	
	two[0]=0x01
	two[1]=0xA9
	two[2]=0xFE
	if (bytes.Compare(Rev(one), two)!=0){
		return false
	}
	return true
}


/*


func HexRev2Uint64(b []byte) uint64{
	return Hex2Uint64(Rev(b))
}

func HexRev2Uint32(b []byte) uint32{
	return Hex2Uint32(Rev(b))
}

//TODO: add to tests and test
func Hex2Uint64(b []byte) uint64{
	var answer uint64
	answer=0
	if len(b)>0{
		maxBytes:=8
		if len(b)<8{
			maxBytes=len(b)
		}
		for i:=0;i<maxBytes;i++{
			answer*=2
			answer+=uint64(b[i])
		}
	}
	return answer
}
//TODO: add to tests and test
func Hex2Uint32(b []byte) uint32{
	var answer uint32
	answer=0
	if len(b)>0{
		maxBytes:=4
		if len(b)<4{
			maxBytes=len(b)
		}
		for i:=0;i<maxBytes;i++{
			answer*=2
			answer+=uint32(b[i])
		}
	}
	return answer
}



func Hex2Big(b []byte) *big.Int{
	answer:=big.NewInt(0)
	
	for i:=0;i<len(b);i++{
		answer.Lsh(answer, 8)
		answer.Add(answer, big.NewInt(int64(b[i])))
	}
	
	return answer
}

func HexRev2Big(rev []byte) *big.Int{
	
	b:=make([]byte, len(rev))
	
	for i:=0;i<len(rev);i++{
		b[len(rev)-i-1]=rev[i]
	}
	
	answer:=big.NewInt(0)
	
	for i:=0;i<len(b);i++{
		answer.Lsh(answer, 8)
		answer.Add(answer, big.NewInt(int64(b[i])))
	}
	
	return answer
}

//TODO: check
func Big2Hex32(b *big.Int) [32]byte{
	var answer [32]byte
	
	tmp:=b.Bytes();
	
	if len(tmp)<=32{
		for i:=0;i<len(tmp);i++{
			answer[31-i]=tmp[i]
		}
	}
	
	
	return answer
}

func String2Hex(s string) []byte{
	answer, _:=hex.DecodeString(s)
	return answer
}

func Str2Hex(s string) []byte{
	return String2Hex(s)
}

func String2HexRev(s string) []byte{
	answer, _:=hex.DecodeString(s)
	return Rev(answer)
}

func Str2HexRev(s string) []byte{
	return Rev(String2Hex(s))
}

//TODO: add to tests and test
func Hex2String(b []byte)string{
	return strings.ToUpper(hex.EncodeToString(b))
}
//TODO: add to tests and test
func Hex2Str(b []byte)string{
	return Hex2String(b)
}
//TODO: add to tests and test
func HexRev2String(b []byte)string{
	return Hex2String(Rev(b))
}
//TODO: add to tests and test
func HexRev2Str(b []byte)string{
	return Hex2String(Rev(b))
}

func Uint322Hex(ui uint32) []byte{
	answer:=make([]byte, 4)
	answer[3]=uint8(ui%256)
	ui/=256
	answer[2]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	ui/=256
	answer[0]=uint8(ui%256)
	
	return answer
}

func Uint322HexRev(ui uint32) []byte{
	answer:=make([]byte, 4)
	answer[0]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	ui/=256
	answer[2]=uint8(ui%256)
	ui/=256
	answer[3]=uint8(ui%256)
	
	return answer
}

func Uint642Hex(ui uint64) []byte{
	answer:=make([]byte, 8)
	answer[7]=uint8(ui%256)
	ui/=256
	answer[6]=uint8(ui%256)
	ui/=256
	answer[5]=uint8(ui%256)
	ui/=256
	answer[4]=uint8(ui%256)
	ui/=256
	answer[3]=uint8(ui%256)
	ui/=256
	answer[2]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	ui/=256
	answer[0]=uint8(ui%256)
	
	return answer
}

func Uint642HexRev(ui uint64) []byte{
	answer:=make([]byte, 8)
	answer[0]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	ui/=256
	answer[2]=uint8(ui%256)
	ui/=256
	answer[3]=uint8(ui%256)
	ui/=256
	answer[4]=uint8(ui%256)
	ui/=256
	answer[5]=uint8(ui%256)
	ui/=256
	answer[6]=uint8(ui%256)
	ui/=256
	answer[7]=uint8(ui%256)
	
	return answer
}

func Uint162Hex(ui uint16) []byte{
	answer:=make([]byte, 2)
	answer[1]=uint8(ui%256)
	ui/=256
	answer[0]=uint8(ui%256)
	
	return answer
}

func Uint162HexRev(ui uint16) []byte{
	answer:=make([]byte, 2)
	answer[0]=uint8(ui%256)
	ui/=256
	answer[1]=uint8(ui%256)
	
	return answer
}

func Randuint64() []byte{
	uint64max:=big.NewInt(1)
	uint64max.Lsh(uint64max, 64)
	
	randnum, _:=rand.Int(rand.Reader, uint64max)
	
	random:=randnum.Bytes()
	answer:=make([]byte, 8)
	
	for i:=0;i<len(random);i++{
		answer[i]=random[i]
	}
	
	return answer
}

func Randuint64Rev() []byte{
	uint64max:=big.NewInt(1)
	uint64max.Lsh(uint64max, 64)
	
	randnum, _:=rand.Int(rand.Reader, uint64max)
	
	random:=randnum.Bytes()
	answer:=make([]byte, 8)
	
	for i:=0;i<len(random);i++{
		answer[len(answer)-1-i]=random[i]
	}
	
	return answer
}

func ConcatBytes(list ...[]byte) []byte{
	size:=0//size of the resulting concatenated list
	for i:=0; i<len(list);i++{
		size+=len(list[i])//counting the sizes of individual parts of the list
	}
	answer:= make([]byte, size)//creates the array for the answer
	
	iterator:=0//iterator to count the position in the answer array
	for i:=0;i<len(list);i++{
		copy(answer[iterator:], list[i])//copies the data into the answer array
	 	iterator+=len(list[i])
	}
	return answer//returns the result
}

func AddByte(one []byte, two byte) []byte{
	size:=len(one)+1//size of the resulting concatenated list

	answer:= make([]byte, size)//creates the array for the answer
	
	copy(answer[0:], one)//copies the data into the answer array
	answer[len(one)]=two
	return answer//returns the result
}

func Byte2String(b []byte) string{
	return bytes.NewBuffer(b).String()
}
*/