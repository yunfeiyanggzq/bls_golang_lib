package BLS
import (
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"github.com/Nik-U/pbc"
	"os"
)
var  BLS_sys   *BLS
type  BLS  struct{
	bls_pairing  *pbc.Pairing
	bls_g        *pbc.Element
}
type BLS_byte  struct{
	Params       string
	G           []byte
}
var blsfilePath="./BLS/bls.gob"

func main(){
	Save_bls_sys_into_lib()
	BLS_start()
	privKey,pubKey:=Generate_bls_keypair()
	signature :=Bls_signature([]byte("guzhiqiang") ,privKey)
	sibyte:=SetSIGIntoByte(signature)
	sign:=SetPubKeyFromByte(sibyte)
	Bls_verify([]byte("guzhiqiang") ,pubKey,sign)
}

func SetPriKeyIntoByte(privkey  *pbc.Element)[]byte{
	return privkey.Bytes()
}
func SetPriKeyFromByte(privkey  []byte)*pbc.Element{
	return   BLS_sys.bls_pairing.NewZr().SetBytes(privkey)
}

func SetPubKeyIntoByte(pubkey  *pbc.Element)[]byte{
	return  pubkey.Bytes()
}
func SetPubKeyFromByte(pubkey  []byte)*pbc.Element{
	return BLS_sys.bls_pairing.NewG2().SetBytes(pubkey)
}
func SetSIGIntoByte(sig  *pbc.Element)[]byte{
	return  sig.Bytes()
}
func SetSIGFromByte(sig  []byte)*pbc.Element{
	return BLS_sys.bls_pairing.NewG2().SetBytes(sig)
}

func Save_bls_sys_into_lib(){
	params := pbc.GenerateA(160, 512)
	para_byte:=params.String()
	g_byte:=params.NewPairing().NewG2().Rand().Bytes()
	bls_byte:=BLS_byte{para_byte,g_byte }
	blsSaveToFilebyte(bls_byte)
}
func BLS_start(){
	blsbyte_:=blsLoadFromFile()
	pairing, _ := pbc.NewPairingFromString(blsbyte_.Params)
	g := pairing.NewG2().SetBytes(blsbyte_.G)
	BLS_sys=&BLS{pairing,g}

}
func  blsSaveToFilebyte(bls BLS_byte) {
	pa := &bls
	file, _ := os.OpenFile(blsfilePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	enc.Encode(pa)
}
func blsLoadFromFile(  )*BLS_byte {
	file, _ := os.Open(blsfilePath)
	defer file.Close()
	var pa BLS_byte
	dec := gob.NewDecoder(file)
	dec.Decode(&pa)
	return &pa
}
func  Generate_bls_keypair()(*pbc.Element,*pbc.Element){
	privKey :=BLS_sys.bls_pairing.NewZr().Rand()
	//g := _pairing.NewG2().Rand()
	pubKey :=BLS_sys.bls_pairing.NewG2().PowZn(BLS_sys.bls_g, privKey)
	return  privKey,pubKey
}
func  Bls_signature(message  []byte,privkey *pbc.Element) *pbc.Element{
	h := BLS_sys.bls_pairing.NewG1().SetFromStringHash(string(message), sha256.New())
	signature := BLS_sys.bls_pairing.NewG2().PowZn(h, privkey)
	return  signature
}
func  Bls_verify(message  []byte,pubkey  *pbc.Element,signature  *pbc.Element)bool{
	h := BLS_sys.bls_pairing.NewG1().SetFromStringHash(string(message), sha256.New())
	temp1 := BLS_sys.bls_pairing.NewGT().Pair(h, pubkey)
	temp2 := BLS_sys.bls_pairing.NewGT().Pair(signature, BLS_sys.bls_g)
	if !temp1.Equals(temp2) {
		fmt.Println("*BUG* Signature check failed *BUG*")
		return false
	} else {
		fmt.Println("*Signature check  success*")
		return  true
	}
}




















