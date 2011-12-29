package mymath
//Subpackage used for all variations of hashing algorithms


import(
    "hash"
    "crypto/sha256"
    "crypto/sha1"
    "crypto/ripemd160"
)

//TODO: test and add to tests
//double SHA-256 hashing of a single byte array
func DoubleSHA(b []byte)([]byte){
   var h hash.Hash = sha256.New()
   h.Write(b)
   var h2 hash.Hash = sha256.New()
   h2.Write(h.Sum())
   
   return h2.Sum()
}

//TODO: test and add to tests
//reverse double SHA-256 hashing of a single byte array
func DoubleSHARev(b []byte)([]byte){   
   return Rev(DoubleSHA(Rev(b)))
}

//TODO: test and add to tests
//Single SHA-256 hashing of a single byte array
func SingleSHA(b []byte)([]byte){
   var h hash.Hash = sha256.New()
   h.Write(b)
   
   return h.Sum()
}

//TODO: test and add to tests
//Reversed single SHA-256 hashing of a single byte array
func SingleSHARev(b []byte)([]byte){   
   return Rev(SingleSHA(Rev(b)))
}

//TODO: test and add to tests
//Single SHA-1 hashing of a single byte array
func SingleSHA1(b []byte)([]byte){
   var h hash.Hash = sha1.New()//TODO: double check
   h.Write(b)
   
   return h.Sum()
}

//Reversed SHA-1 hashing of a single byte array
//TODO: test and add to tests
func SingleSHA1Rev(b []byte)([]byte){   
   return Rev(SingleSHA1(Rev(b)))//TODO: double check
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
func DoubleSHAPair(a []byte, b []byte)([]byte){

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
func DoubleSHAPairRev(a []byte, b []byte)([]byte){
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
func SHARipemd(b []byte)([]byte){
	
	//sha hashing of the input
    var h hash.Hash = sha256.New()
    h.Write(b)

	//ripemd hashing of the sha hash
    var h2 hash.Hash = ripemd160.New()
    h2.Write(h.Sum())

    return h2.Sum()//return
}

//TODO: test and add to tests
//reverse SHA-256 RIPEMD-160 hash
func SHARipemdRev(b []byte)([]byte){
   return Rev(SHARipemd(Rev(b)))
}

//TODO: test and add to tests
//RIPEMD-160 operation for bitcoin address hashing
func Ripemd(b []byte)([]byte){
	//ripemd hashing of the sha hash
    var h hash.Hash = ripemd160.New()
    h.Write(b)

    return h.Sum()//return
}

//TODO: test and add to tests
//reverse RIPEMD-160 hash
func RipemdRev(b []byte)([]byte){
   return Rev(Ripemd(Rev(b)))
}