// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mymath
//Subpackage used for all Bitcoin-specific math calculations, operations and conversions

import(
	"log"
	"bytes"
	
    "hash"
    "big"
    "crypto/sha256"
    "crypto/ripemd160"
    "crypto/elliptic"
)
//Bitcoin-related math

//TODO: test and add to tests
func Makesecp256k1(){

	var p256 = new(elliptic.Curve)
	 //p=FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2
	 //a=00000000 00000000 00000000 00000000 00000000 00000000 0000000000000000
	 //b=00000000 00000000 00000000 00000000 00000000 00000000 0000000000000007
	 //G, compressed = 02 79BE667E F9DCBBAC 55A06295 CE870B07 029BFCDB 2DCE28D9 59F2815B 16F81798
	 //G, uncompressed = 04 79BE667E F9DCBBAC 55A06295 CE870B07 029BFCDB 2DCE28D9 59F2815B 16F81798 483ADA77 26A3C465 5DA4FBFC 0E1108A8 FD17B448 A6855419 9C47D08F FB10D4B8
    //n=FFFFFFFF FFFFFFFF FFFFFFFF FFFFFFFE BAAEDCE6 AF48A03B BFD25E8C D0364141
    //h= 01
    
    
    p256.P, _ = new(big.Int).SetString("115792089210356248762697446949407573530086143415290314195533631308867097853951", 10)
	p256.N, _ = new(big.Int).SetString("115792089210356248762697446949407573529996955224135760342422259061068512044369", 10)
	p256.B, _ = new(big.Int).SetString("5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b", 16)
	p256.Gx, _ = new(big.Int).SetString("6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296", 16)
	p256.Gy, _ = new(big.Int).SetString("4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5", 16)
	p256.BitSize = 256

}

//TODO: test and add to tests
//double SHA-256 hashing of a single byte array
func DoubleSHA1(b []byte)([]byte){
   var h hash.Hash = sha256.New()
   h.Write(b)
   var h2 hash.Hash = sha256.New()
   h2.Write(h.Sum())
   
   return h2.Sum()
}

//TODO: test and add to tests
//reverse double SHA-256 hashing of a single byte array
func DoubleSHA1Rev(b []byte)([]byte){   
   return Rev(DoubleSHA1(Rev(b)))
}

//TODO: test and add to tests
//double hash input bytes, double hash their concatanation
func DoubleDoubleSHA(a []byte, b []byte)([]byte){
	var tmp []byte
	
	//hash first input
    var h1 hash.Hash = sha256.New()
    h1.Write(a)
   
   //hash the first input the second time
    tmp=h1.Sum()
    h1=sha256.New()
    h1.Write(tmp)
    
    //hash the second input
    var h2 hash.Hash = sha256.New()
    h2.Write(b)
   
	//hash the second input the second time
    tmp=h2.Sum()
    h2=sha256.New()
    h2.Write(tmp)
    
    //hash the concatenation of the double hashes of both inputs
    var answer hash.Hash=sha256.New()
    answer.Write(append(h1.Sum(), h2.Sum()...))
    
    //double hash the concatenation
    tmp=answer.Sum()
    answer=sha256.New()
    answer.Write(tmp)
   
    return answer.Sum()//return result
}

//TODO: test and add to tests
//double SHA-256 hash of a concatenation of the two inputs
func DoubleSHA(a []byte, b []byte)([]byte){

	var tmp []byte
   
    //hash the concatenation of the two inputs
    var answer hash.Hash=sha256.New()
    answer.Write(append(a, b...))
    
    //hash it the second time
    tmp=answer.Sum()
    answer=sha256.New()
    answer.Write(tmp)
   
    return answer.Sum()//return
}

//TODO: test and add to tests
//double SHA-256 hash of a concatenation of the reverse of two inputs
func DoubleSHARev(a []byte, b []byte)([]byte){
	var tmp []byte

	//hash the concatenation of the reverse of both inputs
    var answer hash.Hash=sha256.New()
    //answer.Write(append(a[:]+b[:]))
    answer.Write(append(Rev(a), Rev(b)...))
    
    //hash it again
    tmp=answer.Sum()
    answer=sha256.New()
    answer.Write(tmp)
   
    return Rev(answer.Sum())//return the reverse of the output
}

//TODO: test and add to tests
//SHA-256 RIPEMD-160 operation for bitcoin address hashing
func Ripemd(b []byte)([]byte){
	
	//sha hashing of the input
    var h hash.Hash = sha256.New()
    h.Write(b)

	//ripemd hashing of the sha hash
    var h2 hash.Hash = ripemd160.New()
    h2.Write(h.Sum())

    return h2.Sum()//return
}

//TODO: test and add to tests
//reverse RIPEMD-160 hash
func RipemdRev(b []byte)([]byte){
   return Rev(Ripemd(Rev(b)))
}

type VarInt uint64

func (vi VarInt)Len()int{
	if vi<0xfd{
		return 1
	}
	if vi<=0xffff{
		return 3
	}
	if vi<=0xffffffff{
		return 5
	}
	return 9
}

func (vi VarInt)Compile()[]byte{
	return VarInt2HexRev(vi)
}

//TODO: test
func DecodeVarInt(b []byte)VarInt{
	if len(b)==0{
		return 0
	}
	if b[0]<0xFD{
		return VarInt(b[0])
	}
	if b[0]==0xFD{
		return VarInt(Hex2Uint64(b[1:3]))
	}
	if b[0]==0xFE{
		return VarInt(Hex2Uint64(b[1:5]))
	}
	if b[0]==0xFF{
		return VarInt(Hex2Uint64(b[1:9]))
	}
	return 0
}
//TODO: check and add to tests
func DecodeVarIntGiveRest(b []byte) (VarInt, []byte){
	if len(b)==0{
		return 0, b
	}
	if b[0]<0xFD{
		return VarInt(b[0]), b[1:]
	}
	if b[0]==0xFD{
		return VarInt(Hex2Uint64(b[1:3])), b[4:]
	}
	if b[0]==0xFE{
		return VarInt(Hex2Uint64(b[1:5])), b[6:]
	}
	if b[0]==0xFF{
		return VarInt(Hex2Uint64(b[1:9])), b[10:]
	}
	return 0, b
}

//variable-length hex result
//https://en.bitcoin.it/wiki/Protocol_specification#Variable_length_integer
func VarInt2Hex(vi VarInt) []byte{
	answer:=make([]byte, 9)
	copy(answer[1:], Uint642Hex(uint64(vi)))
	if vi<0xfd {
		return answer[8:]
	}
	if vi<=0xffff {
		answer[6]=0xfd
		return answer[6:]
	}
	if vi<=0xffffffff {
		answer[4]=0xfe
		return answer[4:]
	}
	answer[0]=0xff
	return answer
}

//variable-length hex result
//https://en.bitcoin.it/wiki/Protocol_specification#Variable_length_integer
func VarInt2HexRev(vi VarInt) []byte{
	answer:=make([]byte, 9)
	copy(answer[1:], Uint642HexRev(uint64(vi)))
	
	if vi<0xfd {
		return answer[1:2]
	}
	if vi<=0xffff {
		answer[0]=0xfd
		return answer[0:3]
	}
	if vi<=0xffffffff {
		answer[0]=0xfe
		return answer[0:5]
	}
	answer[0]=0xff
	return answer
}

type VarStr struct{
	length VarInt
	str []byte
}

func (vs *VarStr)Len()int{
	return vs.length.Len()+len(vs.str)
}

func (vs *VarStr)Set(newString string){
	vs.str=make([]byte, len(newString))
	
	for i:=0;i<len(newString);i++{
		vs.str[i]=newString[i]
	}
	vs.length=VarInt(len(newString))
}

func (vs *VarStr)Compile()[]byte{
	answer:=make([]byte, vs.Len())
	copy(answer[:], vs.length.Compile())
	copy(answer[vs.length.Len():], vs.str)
	return answer
}

func (vs *VarStr)Read()string{
	return string(vs.str)
}
//TODO: test and add to tests
func GenerateMerkleTreeFromString(leafs []string)[]string{
	tmp:=make([][]byte, len(leafs))
	
	for i:=0;i<len(tmp);i++{
		tmp[i]=make([]byte, len(leafs[i])/2)
		copy(tmp[i][:], Str2Hex(leafs[i]))
	}
	merkletree:=GenerateMerkleTree(tmp)
	
	answer:=make([]string, len(merkletree))
	for i:=0;i<len(merkletree);i++{
		answer[i]=Hex2Str(merkletree[i])
	}
	return answer
}

//TODO: test and add to tests
func GenerateMerkleTree(leafs [][]byte)[][]byte{
	answer:=make([][]byte, len(leafs))
	for i:=0;i<len(answer);i++{
		answer[i]=make([]byte, len(leafs[i]))
		copy(answer[i][:], leafs[i][:])
	}
	level:=make([][]byte, len(leafs))
	for i:=0;i<len(answer);i++{
		level[i]=make([]byte, len(leafs[i]))
		copy(level[i][:], leafs[i][:])
	}
	//answer:=leafs[:]
	//level:=leafs[:]
	for;len(level)>1;{
		//currentlevel:=level[:int(math.Ceil(float64(len(level)/2.0)))]
		//currentlevel:=make([][]byte, int(math.Ceil(float64(1.0*len(level)/2.0))))
		currentlevel:=make([][]byte, len(level)/2+len(level)%2)
		log.Printf("len - %d", len(currentlevel))
		for i:=0;i<len(currentlevel)-1;i++{
			currentlevel[i]=DoubleSHARev(level[2*i], level[2*i+1])
		}
		if(len(level)%2==1){
			currentlevel[len(currentlevel)-1]=DoubleSHARev(level[2*(len(currentlevel)-1)], level[2*(len(currentlevel)-1)])
		}else{
			currentlevel[len(currentlevel)-1]=DoubleSHARev(level[2*(len(currentlevel)-1)], level[2*(len(currentlevel)-1)+1])
		}
		//answer=append(answer, currentlevel)
		tmp:=make([][]byte, len(answer)+len(currentlevel))
		for i:=0;i<len(answer);i++{
			tmp[i]=make([]byte, len(answer[i]))
			copy(tmp[i][:], answer[i][:])
		}
		level=make([][]byte, len(currentlevel))
		for i:=0;i<len(currentlevel);i++{
			tmp[len(answer)+i]=make([]byte, len(currentlevel[i]))
			copy(tmp[len(answer)+i][:], currentlevel[i][:])
			level[i]=make([]byte, len(currentlevel[i]))
			copy(level[i][:], currentlevel[i][:])
		}
		//copy(tmp[:], answer)
		//copy(tmp[len(answer):], currentlevel)
		answer=tmp
		//level=currentlevel
	}
	return answer
}

func TestGenerateMerkleTree()bool{
	/* //http://blockexplorer.com/rawblock/0000000000000ae4e079775dfb20e0571ca5e24d7fc73489a8a48bb9758e17d6
	"b30c84a7e9a68873e40f807557b9af8d364e66f0cbfaa64e684a5ed662e5a197",
    "74182c272bd3b201a6afc1fcaf5fa60c603dbd52f2d3150ce7aa593d697fd01f",
    "48ba51a6a9a7ede4618e5cdeb8cd3a000ca9e369df96a2bdd9cd36563781f072",
    "665960a9ee63b2ab7c27eaff6f8e1cccbb35ef8b8a20a82f3bfd6ce3445057e6",
    "2ba0808f3cbbf21b45916585bc8b2c93c131abb79425fae92e3e0cb6968808ca",
    "58e13ec3fc70f745a2e6345f799f030ee6f26499eb4a59959b92155857325241",
    "96ec719655fa0a59a1570eab9172efb67c8de29d091486a10f0a1514dadc5e3a",
    "ee96e27544effdb27f19d90b1938b2e60c84f4091521a4082b0d91f16f27fdae",
    "92ea9d03ae8e51874030c69958734607246d645c1f047b11edf6be2d595e80b9",
    "6250804544c2010a86b147f7e00f1943794396f009ba2239728497c74cc2aeb8",
    
    "c84f67766b3cccca370a317d44cd042861629434f127715065e4a35d6cda32e9",
    "533638ea7a72c2ef1cf43cbd56861c7329c064f2830e380f04a686cd5a91db28",
    "af7aac7595207ad1cd5e80d1cb2c875dab6674b363cea87482912d514345dad4",
    "a3c61f0642c833f9a804bc92f824feed1927e903103df2d34a8d3987d52b2838",
    "dcff5d86211ea488879b9d495126371fa0eca322d8e14cbf09a46bf2ce7ba045",
    
    "25ee9f74d7e462e3432981ba231fbb630944575704dbfbf61a03f21371f04a80",
    "ea7894cfeb08bcd7a7de6d731cf036087a335221b0b4e3bf123252ffaff4b5f8",
    "e2d0b6819628e8f64d2e6ac632285c7cad9d5f2d78524317a108f26b66aebb0d",
    
    "fd918efdc92eac85717d23083c5d749d5cd2c43c153cb11d6299100fb743788d",
    "f47bf9fe04c7051d09461f0930c695413c51f2e9831e8c5efa6700e2f22161ab",
    
    "d6402a834394e961a8d07c1cd4d949badeefd371230f0f49fda9cb96158dab92"
    */
	merkletree:=make([][]byte, 10)
	merkletree[0]=String2Hex("b30c84a7e9a68873e40f807557b9af8d364e66f0cbfaa64e684a5ed662e5a197")
	merkletree[1]=String2Hex("74182c272bd3b201a6afc1fcaf5fa60c603dbd52f2d3150ce7aa593d697fd01f")
	merkletree[2]=String2Hex("48ba51a6a9a7ede4618e5cdeb8cd3a000ca9e369df96a2bdd9cd36563781f072")
	merkletree[3]=String2Hex("665960a9ee63b2ab7c27eaff6f8e1cccbb35ef8b8a20a82f3bfd6ce3445057e6")
	merkletree[4]=String2Hex("2ba0808f3cbbf21b45916585bc8b2c93c131abb79425fae92e3e0cb6968808ca")
	merkletree[5]=String2Hex("58e13ec3fc70f745a2e6345f799f030ee6f26499eb4a59959b92155857325241")
	merkletree[6]=String2Hex("96ec719655fa0a59a1570eab9172efb67c8de29d091486a10f0a1514dadc5e3a")
	merkletree[7]=String2Hex("ee96e27544effdb27f19d90b1938b2e60c84f4091521a4082b0d91f16f27fdae")
	merkletree[8]=String2Hex("92ea9d03ae8e51874030c69958734607246d645c1f047b11edf6be2d595e80b9")
	merkletree[9]=String2Hex("6250804544c2010a86b147f7e00f1943794396f009ba2239728497c74cc2aeb8")
	
	//log.Printf("merkletree - %dx%d", len(merkletree), len(merkletree[0]))
	//log.Printf("b30c84a7e9a68873e40f807557b9af8d364e66f0cbfaa64e684a5ed662e5a197 = %X", String2Hex("b30c84a7e9a68873e40f807557b9af8d364e66f0cbfaa64e684a5ed662e5a197"))
	
	/*one:=make([]byte, 2)
	two:=make([]byte, 2)
	one[0]=0x00
	one[1]=0x01
	two[0]=0x02
	two[1]=0x03
	
	log.Printf("%X", append(one[:], two[:]))
	
	three:=[]byte{0, 1}
	four:=[]byte{2, 3}
	
	five:=append(three, four)*/
	
	
	
	//log.Printf("aaaa+bbbb=%X", append(String2Hex("aaaa"), String2Hex("bbbb")))
	
	/*for i:=0;i<len(merkletree);i++{
		log.Printf("merkletree %d - %X", i, merkletree[i])
	}*/
	answer:=GenerateMerkleTree(merkletree)
	
	if(len(answer)==21){
		return true
	}
	
	/*for i:=0;i<len(merkletree);i++{
		log.Printf("merkletree %d - %X", i, merkletree[i])
	}
	for i:=0;i<len(answer);i++{
		log.Printf("answer %d - %X", i, answer[i])
	}*/
	return false
}




//Testing

func TestEverythingBitmath()bool{
	if(TestVarInt()==false){
		log.Print("TestVarInt()==false")
		return false
	}
	if(TestCompile()==false){
		log.Print("TestCompile()==false")
		return false
	}
	if(TestLen()==false){
		log.Print("TestLen()==false")
		return false
	}
	if(TestVarInt2Hex()==false){
		log.Print("TestVarInt2Hex()==false")
		return false
	}
	if(TestVarInt2HexRev()==false){
		log.Print("TestVarInt2HexRev()==false")
		return false
	}
	if(TestVarStr()==false){
		log.Print("TestVarStr()==false")
		return false
	}
	//log.Print("All Bitmath tests okay!")
	TestGenerateMerkleTree()
	return true
}



func TestVarInt()bool{
	var vi1 VarInt
	vi1=0
	
	var vi2 VarInt
	vi2=1
	
	var vi3 VarInt
	vi3=0xffffffff
	
	if(vi1!=0){
		return false
	}
	if(vi2!=1){
		return false
	}
	if(vi3!=0xffffffff){
		return false
	}
	
	return true
}

func TestCompile()bool{
	if(bytes.Compare(VarInt(0).Compile(), VarInt2HexRev(VarInt(0)))!=0){
		log.Printf("%X - %X", VarInt(0).Compile(), VarInt2Hex(VarInt(0)))
		return false
	}
	if(bytes.Compare(VarInt(1).Compile(), VarInt2HexRev(VarInt(1)))!=0){
		log.Printf("%X - %X", VarInt(1).Compile(), VarInt2Hex(VarInt(1)))
		return false
	}
	if(bytes.Compare(VarInt(0x010000).Compile(), VarInt2HexRev(VarInt(0x010000)))!=0){
		log.Printf("%X - %X", VarInt(0x010000).Compile(), VarInt2Hex(VarInt(0x010000)))
		return false
	}
	if(bytes.Compare(VarInt(0xffffffff).Compile(), VarInt2HexRev(VarInt(0xffffffff)))!=0){
		log.Printf("%X - %X", VarInt(0xffffffff).Compile(), VarInt2Hex(VarInt(0xffffffff)))
		return false
	}
	return true
}

func TestLen()bool{
	if(VarInt(0).Len()!=1){
		log.Print("TestLen(0)==false")
		return false
	}
	if(VarInt(1).Len()!=1){
		log.Print("TestLen(1)==false")
		return false
	}
	if(VarInt(0xFC).Len()!=1){
		log.Print("TestLen(fc)==false")
		return false
	}
	if(VarInt(0xFD).Len()!=3){
		log.Print("TestLen(fd)==false")
		return false
	}
	if(VarInt(0xFFFE).Len()!=3){
		log.Print("TestLen(fffe)==false")
		return false
	}
	if(VarInt(0xFFFF).Len()!=3){
		log.Print("TestLen(ffff)==false")
		return false
	}
	if(VarInt(0x010000).Len()!=5){
		log.Print("TestLen(0x010000)==false")
		return false
	}
	if(VarInt(0xFFFFFFFE).Len()!=5){
		log.Print("TestLen(fffffffe)==false")
		return false
	}
	if(VarInt(0xFFFFFFFF).Len()!=5){
		log.Print("TestLen(ffffffff)==false")
		return false
	}
	if(VarInt(0x0100000000).Len()!=9){
		log.Print("TestLen(0x0100000000)==false")
		return false
	}
	return true
}

func TestVarInt2Hex() bool{
	var byte1 []byte
	var byte3 []byte
	var byte5 []byte
	var byte9 []byte
	
	byte1=make([]byte, 1)
	byte3=make([]byte, 3)
	byte5=make([]byte, 5)
	byte9=make([]byte, 9)
	
	byte1[0]=0
	byte3[0]=0
	byte5[0]=0
	byte9[0]=0
	
	if(bytes.Compare(byte1, VarInt2Hex(VarInt(0)))!=0){
		log.Print("0")
		return false
	}
	byte1[0]=0xF0
	if(bytes.Compare(byte1, VarInt2Hex(VarInt(0xF0)))!=0){
		log.Print("F0")
		return false
	}
	byte1[0]=0xFC
	if(bytes.Compare(byte1, VarInt2Hex(VarInt(0xFC)))!=0){
		log.Print("FC")
		return false
	}
	byte3[0]=0xFD
	byte3[1]=0
	byte3[2]=0xFD
	if(bytes.Compare(byte3, VarInt2Hex(VarInt(0xFD)))!=0){
		log.Printf("FD - %X", VarInt2Hex(VarInt(0xFD)))
		return false
	}
	byte3[2]=0xFF
	if(bytes.Compare(byte3, VarInt2Hex(VarInt(0xFF)))!=0){
		log.Print("FF")
		return false
	}
	byte3[1]=0xFF
	byte3[2]=0xFE
	if(bytes.Compare(byte3, VarInt2Hex(VarInt(0xFFFE)))!=0){
		log.Print("FFFE")
		return false
	}
	byte3[2]=0xFF
	if(bytes.Compare(byte3, VarInt2Hex(VarInt(0xFFFF)))!=0){
		log.Print("FFFF")
		return false
	}
	
	byte5[0]=0xFE
	byte5[1]=0
	byte5[2]=0x01
	byte5[3]=0
	byte5[4]=0
	if(bytes.Compare(byte5, VarInt2Hex(VarInt(0x010000)))!=0){
		log.Print("0x010000")
		return false
	}
	byte5[1]=0xFF
	byte5[2]=0xFF
	byte5[3]=0xFF
	byte5[4]=0xFE
	if(bytes.Compare(byte5, VarInt2Hex(VarInt(0xFFFFFFFE)))!=0){
		log.Print("0xFFFFFFFE")
		return false
	}
	byte5[4]=0xFF
	if(bytes.Compare(byte5, VarInt2Hex(VarInt(0xFFFFFFFF)))!=0){
		log.Print("0xFFFFFFFF")
		return false
	}
	
	byte9[0]=0xFF
	byte9[1]=0
	byte9[2]=0
	byte9[3]=0
	byte9[4]=0x01
	byte9[5]=0
	byte9[6]=0
	byte9[7]=0
	byte9[8]=0
	
	if(bytes.Compare(byte9, VarInt2Hex(VarInt(0x0100000000)))!=0){
		log.Print("0x0100000000")
		return false
	}
	
	return true
}

func TestVarInt2HexRev() bool{
	var byte1 []byte
	var byte3 []byte
	var byte5 []byte
	var byte9 []byte
	
	byte1=make([]byte, 1)
	byte3=make([]byte, 3)
	byte5=make([]byte, 5)
	byte9=make([]byte, 9)
	
	byte1[0]=0
	byte3[0]=0
	byte5[0]=0
	byte9[0]=0
	
	if(bytes.Compare(byte1, VarInt2HexRev(VarInt(0)))!=0){
		log.Print("0")
		return false
	}
	byte1[0]=0xF0
	if(bytes.Compare(byte1, VarInt2HexRev(VarInt(0xF0)))!=0){
		log.Print("F0")
		return false
	}
	byte1[0]=0xFC
	if(bytes.Compare(byte1, VarInt2HexRev(VarInt(0xFC)))!=0){
		log.Print("FC")
		return false
	}
	byte3[0]=0xFD
	byte3[2]=0
	byte3[1]=0xFD
	if(bytes.Compare(byte3, VarInt2HexRev(VarInt(0xFD)))!=0){
		log.Printf("FD - %X", VarInt2HexRev(VarInt(0xFD)))
		return false
	}
	byte3[1]=0xFF
	if(bytes.Compare(byte3, VarInt2HexRev(VarInt(0xFF)))!=0){
		log.Print("FF")
		return false
	}
	byte3[2]=0xFF
	byte3[1]=0xFE
	if(bytes.Compare(byte3, VarInt2HexRev(VarInt(0xFFFE)))!=0){
		log.Print("FFFE")
		return false
	}
	byte3[1]=0xFF
	if(bytes.Compare(byte3, VarInt2HexRev(VarInt(0xFFFF)))!=0){
		log.Print("FFFF")
		return false
	}
	
	byte5[0]=0xFE
	byte5[4]=0
	byte5[3]=0x01
	byte5[2]=0
	byte5[1]=0
	if(bytes.Compare(byte5, VarInt2HexRev(VarInt(0x010000)))!=0){
		log.Print("0x010000")
		return false
	}
	byte5[4]=0xFF
	byte5[3]=0xFF
	byte5[2]=0xFF
	byte5[1]=0xFE
	if(bytes.Compare(byte5, VarInt2HexRev(VarInt(0xFFFFFFFE)))!=0){
		log.Print("0xFFFFFFFE")
		return false
	}
	byte5[1]=0xFF
	if(bytes.Compare(byte5, VarInt2HexRev(VarInt(0xFFFFFFFF)))!=0){
		log.Print("0xFFFFFFFF")
		return false
	}
	
	byte9[0]=0xFF
	byte9[8]=0
	byte9[7]=0
	byte9[6]=0
	byte9[5]=0x01
	byte9[4]=0
	byte9[3]=0
	byte9[2]=0
	byte9[1]=0
	
	if(bytes.Compare(byte9, VarInt2HexRev(VarInt(0x0100000000)))!=0){
		log.Print("0x0100000000")
		return false
	}
	return true
}

func TestVarStr()bool{
	vs1:=new(VarStr)
	vs1.Set("Hello!")
	
	if(vs1.Len()!=7){
		log.Print("s1.Len()")
		return false
	}
	if(vs1.Read()!="Hello!"){
		log.Print("s1.Read()")
		return false
	}
	
	compiled:=make([]byte, 7)
	compiled[0]=0x06
	compiled[1]=0x48
	compiled[2]=0x65
	compiled[3]=0x6C
	compiled[4]=0x6C
	compiled[5]=0x6F
	compiled[6]=0x21
	if(bytes.Compare(compiled, vs1.Compile())!=0){
		log.Print("vs1.Compile()")
		return false
	}
	
	return true
}