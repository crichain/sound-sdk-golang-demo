/**
* @Author:wjp
* @Date:2022/7/29 14:00
 */
package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
)

func Sign(data []byte, prv *ecdsa.PrivateKey) string {
	//if len(data) <= 2 {
	//	return ""
	//}
	//if data[0:2] != "0x" {
	//	data = fmt.Sprintf("%s%s", "0x", data)
	//}
	//aaaa, err := hexutil.Decode(data)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//str := ""
	//for _, v := range aaaa {
	//	str = fmt.Sprintf("%s%s", str, string(v))
	//}
	fmt.Println(hexutil.Encode(data))
	sum256 := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, prv, sum256[:])
	if err != nil {
		return ""
	}
	pub := FromECDSAPub(&prv.PublicKey)
	addr := PubkeyToAddress(prv.PublicKey)
	pubstring := strings.ToLower(hex.EncodeToString(pub))
	rs := reverse(r.Bytes())

	ys := reverse(s.Bytes())

	return fmt.Sprintf("%s%s%s%s", pubstring, strings.ToLower(addr.Hex()[2:]), hex.EncodeToString(rs)[0:64], hex.EncodeToString(ys)[0:64])

}

//翻转byte
func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
