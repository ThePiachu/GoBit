// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitmessage

import(
	"mymath"
	"log"
)

//TODO: test
type Tx struct{
	Version [4]byte
	TxInCount mymath.VarInt
	TxIn []*TxIn//[]TxIn
	TxOutCount mymath.VarInt
	TxOut []*TxOut//[]TxOut
	LockTime [4]byte
}

//TODO: test
//TODO: double check if rev or not
func (t *Tx)Hash() []byte{
	return mymath.DoubleSHA(t.Compile())
}

func (t *Tx)GetVersion() []byte{
	return t.Version[:]
}

func (t *Tx)GetTxInCount() mymath.VarInt{
	return t.TxInCount
}

func (t *Tx)GetTxIn() []*TxIn{
	return t.TxIn
}

func (t *Tx)GetTxOutCount() mymath.VarInt{
	return t.TxOutCount
}

func (t *Tx)GetTxOut() []*TxOut{
	return t.TxOut
}

func (t *Tx)GetLockTime() []byte{
	return t.LockTime[:]
}

func (t *Tx)SetVersion(ver uint32){
	answer:=mymath.Uint322HexRev(ver)
	copy(t.Version[:], answer)
}

func (t *Tx)AddIn(txin *TxIn){
	t.TxInCount++
	t.TxIn=append(t.TxIn, txin)
}

func (t *Tx)ClearIn(){
	t.TxInCount=0
	t.TxIn=nil
}

func (t *Tx)AddOut(txout *TxOut){
	t.TxOutCount++
	t.TxOut=append(t.TxOut, txout)
}

func (t *Tx)ClearOut(){
	t.TxOutCount=0
	t.TxOut=nil
}

func (t *Tx)SetLockTime(time uint32){
	answer:=mymath.Uint322HexRev(time)
	copy(t.LockTime[:], answer)
}

//TODO:test
func (t *Tx)Compile()[]byte{

	tic:=mymath.VarInt2HexRev(t.TxInCount)//TODO: double check if it is Rev or not
	toc:=mymath.VarInt2HexRev(t.TxOutCount)//TODO: double check if it is Rev or not
	
	totalLen:=0
	totalLen+=len(t.Version)
	totalLen+=len(tic)
	for i:=0;i<len(t.TxIn);i++{//TODO:check if the pointer is working
		totalLen+=t.TxIn[i].Len()
	}
	totalLen+=len(toc)
	for i:=0;i<len(t.TxOut);i++{//TODO:check if the pointer is working
		totalLen+=t.TxOut[i].Len()
	}
	totalLen+=len(t.LockTime)

	answer:=make([]byte, totalLen)
	
	iterator:=0
	copy(answer[iterator:], t.Version[:])
	iterator+=len(t.Version)
	copy(answer[iterator:], tic)
	iterator+=len(tic)
	for i:=0;i<len(t.TxIn);i++{//TODO:check if the pointer is working
		tmp:=t.TxIn[i].Compile()[:]
		copy(answer[iterator:], tmp)
		iterator+=len(tmp)
	}
	copy(answer[iterator:], toc)
	iterator+=len(toc)
	for i:=0;i<len(t.TxOut);i++{//TODO:check if the pointer is working
		tmp:=t.TxOut[i].Compile()[:]
		copy(answer[iterator:], tmp)
		iterator+=len(tmp)
	}
	copy(answer[iterator:], t.LockTime[:])
	
	return answer
}

//TODO:test
func (t *Tx)Len()int{
	tic:=mymath.VarInt2HexRev(t.TxInCount)//TODO: double check if it is Rev or not
	toc:=mymath.VarInt2HexRev(t.TxOutCount)//TODO: double check if it is Rev or not
	
	totalLen:=0
	totalLen+=len(t.Version)
	totalLen+=len(tic)
	for i:=0;i<len(t.TxIn);i++{//TODO:check if the pointer is working
		totalLen+=t.TxIn[i].Len()
	}
	totalLen+=len(toc)
	for i:=0;i<len(t.TxOut);i++{//TODO:check if the pointer is working
		totalLen+=t.TxOut[i].Len()
	}
	totalLen+=len(t.LockTime)
	return totalLen
}

//TODO: test
func DecodeMultipleTxs(txs []byte, count int) ([]*Tx, []byte){
	//log.Print("DecodeMultipleTxs")
	var answer []*Tx
	var tmp *Tx
	tmpbytes:=txs
	for i:=0;i<count;i++{
		//log.Printf("loop b - %X", tmpbytes)
		tmp, tmpbytes=DecodeTx(tmpbytes)
		if tmp==nil{
			break
		}
		answer=append(answer, tmp)
	}
	
	return answer, tmpbytes
}

//TODO: test
func DecodeTx(tx []byte) (*Tx, []byte){
	//log.Printf("DecodeTx - %X", tx)
	//4+1+41+1+8+4=59
	if len(tx)<4+1+41+1+8+4{
		return nil, tx
	}
	
	answer:=new(Tx)
	/*
	Version [4]byte
	TxInCount mymath.VarInt
	TxIn []*TxIn//[]TxIn
	TxOutCount mymath.VarInt
	TxOut []*TxOut//[]TxOut
	LockTime [4]byte
	*/
	copy(answer.Version[:], tx[0:4])
	var tmpbytes []byte
	//log.Printf("Before ins: %X", tx[4:])
	answer.TxInCount, tmpbytes=mymath.DecodeVarIntGiveRest(tx[4:])
	answer.TxIn, tmpbytes=DecodeMultipleTxIns(tmpbytes, int(answer.TxInCount))
	//log.Printf("Before outs: %X", tmpbytes)
	answer.TxOutCount, tmpbytes=mymath.DecodeVarIntGiveRest(tmpbytes)
	answer.TxOut, tmpbytes=DecodeMultipleTxOuts(tmpbytes, int(answer.TxOutCount))
	//log.Printf("After outs: %X", tmpbytes)
	if len(tmpbytes)<4{
		return nil, tx
	}
	copy(answer.LockTime[:], tmpbytes[0:4])
	
	return answer, tmpbytes[4:]
}

//**************************************************************

//TODO:test
type TxIn struct{
	PreviousOutput OutPoint
	ScriptLength mymath.VarInt
	SignatureScript []byte//uchar[]
	Sequence [4]byte
}

func (ti *TxIn)GetPreviousOutput() OutPoint{
	return ti.PreviousOutput
}

func (ti *TxIn)GetScriptLength() mymath.VarInt{
	return ti.ScriptLength
}

func (ti *TxIn)GetSignatureScript() []byte{
	return ti.SignatureScript[:]
}

func (ti *TxIn)GetSequence() []byte{
	return ti.Sequence[:]
}

func (ti *TxIn)SetOutPoint(out OutPoint){
	ti.PreviousOutput=out
}

func (ti *TxIn)SetSignature(sig []byte){
	ti.SignatureScript=sig
	ti.ScriptLength=mymath.VarInt(len(ti.SignatureScript))
}

func (ti *TxIn)SetSequence(seq uint32){
	answer:=mymath.Uint322HexRev(seq)
	copy(ti.Sequence[:], answer)
}

//TODO:test
func (ti *TxIn)Compile()[]byte{
	sl:=mymath.VarInt2HexRev(ti.ScriptLength)//TODO: check if Rev or not
	po:=ti.PreviousOutput.Compile()

	answer:=make([]byte, len(po)+len(sl)+len(ti.SignatureScript)+len(ti.Sequence))

	iterator:=0
	copy(answer[iterator:], po)	
	iterator+=len(po)
	copy(answer[iterator:], sl)	
	iterator+=len(sl)
	copy(answer[iterator:], ti.SignatureScript[:])	
	iterator+=len(ti.SignatureScript)
	copy(answer[iterator:], ti.Sequence[:])

	return answer
}

//TODO:test
func (ti *TxIn)Len()int{
	return ti.PreviousOutput.Len()+ti.ScriptLength.Len()+len(ti.SignatureScript)+len(ti.Sequence)
}

//TODO: test
func DecodeMultipleTxIns(b []byte, count int) ([]*TxIn, []byte){
	var answer []*TxIn
	var tmp *TxIn
	tmpbytes:=b
	//log.Printf("b - %X", tmpbytes)
	for i:=0;i<count;i++{
		tmp, tmpbytes=DecodeTxIn(tmpbytes)
		if tmp==nil{
			break
		}
		answer=append(answer, tmp)
	}
	
	return answer, tmpbytes
}

//TODO: test
func DecodeTxIn(b []byte) (*TxIn, []byte){
	if len(b)<36+1+4{
		return nil, b
	}
	//log.Printf("txin b - %X", b)
	answer:=new(TxIn)
	answer.PreviousOutput.Hash=NewHash(b[0:32])
	copy(answer.PreviousOutput.Index[:], b[32:36])
	
	var tmpbytes []byte
	answer.ScriptLength, tmpbytes=mymath.DecodeVarIntGiveRest(b[36:])
	if len(tmpbytes)<int(answer.ScriptLength)+4{
		return nil, b
	}
	answer.SignatureScript=tmpbytes[0:answer.ScriptLength]
	copy(answer.Sequence[:], tmpbytes[answer.ScriptLength:answer.ScriptLength+4])

	//log.Printf("txin bytes - %X", tmpbytes[answer.ScriptLength+4:])
	return answer, tmpbytes[answer.ScriptLength+4:]
}

//**************************************************************

//TODO:test
type OutPoint struct{
	Hash Hash
	Index [4]byte
}

func (op *OutPoint)SetHash(newHash Hash){
	op.Hash=newHash
}

func (op *OutPoint)SetIndex(ind uint32){
	answer:=mymath.Uint322HexRev(ind)
	copy(op.Index[:], answer)
}

func (op* OutPoint)GetHash()[]byte{
	return op.Hash[:]
}

func (op* OutPoint)GetIndex()[]byte{
	return op.Index[:]
}

//TODO:test
func (op *OutPoint)Compile()[]byte{
	answer:=make([]byte, len(op.Hash)+len(op.Index))

	iterator:=0
	copy(answer[iterator:], op.Hash[:])	
	iterator+=len(op.Hash)
	copy(answer[iterator:], op.Index[:])

	return answer
}

func (op *OutPoint)Len()int{
	return op.Hash.Len()+len(op.Index)
}

//**************************************************************

//TODO:test
type TxOut struct{
	Value [8]byte
	PkScriptLength mymath.VarInt
	PkScript []byte//uchar[]
}

func (to *TxOut)GetValue() []byte{
	return to.Value[:]
}

func (to *TxOut)GetPkScriptLength() mymath.VarInt{
	return to.PkScriptLength
}

func (to *TxOut)GetPkScript() []byte{
	return to.PkScript[:]
}

func (to *TxOut)SetValue(val uint64){
	answer:=mymath.Uint642HexRev(val)
	copy(to.Value[:], answer)
}

func (to *TxOut)SetScript(script []byte){
	to.PkScript=script
	to.PkScriptLength=mymath.VarInt(len(to.PkScript))
}

func (to *TxOut)SetScriptStr(script string, opcode byte){
	to.PkScript=mymath.AddByte(mymath.String2Hex(script), opcode)
	to.PkScriptLength=mymath.VarInt(len(to.PkScript))
}
//TODO:test
func (to *TxOut)Compile()[]byte{
	vi:=mymath.VarInt2HexRev(to.PkScriptLength)//TODO: check if Rev or not

	answer:=make([]byte, len(to.Value)+len(vi)+len(to.PkScript))

	iterator:=0
	copy(answer[iterator:], to.Value[:])	
	iterator+=len(to.Value)
	copy(answer[iterator:], vi)	
	iterator+=len(vi)
	copy(answer[iterator:], to.PkScript)

	return answer
}

//TODO:test
func (to *TxOut)Len()int{
	return len(to.Value)+to.PkScriptLength.Len()+len(to.PkScript)
}


//TODO: test
func DecodeMultipleTxOuts(b []byte, count int) ([]*TxOut, []byte){
	var answer []*TxOut
	var tmp *TxOut
	tmpbytes:=b
	for i:=0;i<count;i++{
		//log.Printf("input b - %X", tmpbytes)
		tmp, tmpbytes=DecodeTxOut(tmpbytes)
		if tmp==nil{
			break
		}
		answer=append(answer, tmp)
	}
	
	return answer, tmpbytes
}

//TODO: test
func DecodeTxOut(b []byte) (*TxOut, []byte){
	if len(b)<8+1{
		return nil, b
	}
	/*
	
	Value [8]byte
	PkScriptLength mymath.VarInt
	PkScript []byte//uchar[]
	*/
	answer:=new(TxOut)
	copy(answer.Value[:], b[0:8])
	
	var tmpbytes []byte
	answer.PkScriptLength, tmpbytes=mymath.DecodeVarIntGiveRest(b[8:])
	if len(tmpbytes)<int(answer.PkScriptLength){
		log.Printf("b - %X", b)
		log.Printf("len(tmpbytes)=%d", len(tmpbytes))
		log.Printf("PkScriptLength=%d", answer.PkScriptLength)
		return nil, b
	}
	answer.PkScript=make([]byte, answer.PkScriptLength)
	copy(answer.PkScript, tmpbytes[0:answer.PkScriptLength])
	
	//log.Printf("PkScript\n%X", answer.PkScript)
	
	return answer, tmpbytes[answer.PkScriptLength:]
}