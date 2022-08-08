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
	"strings"
)

//签名
func Sign(data []byte, prv *ecdsa.PrivateKey) string {

	sum256 := sha256.Sum256(data)

	r, s, err := ecdsa.Sign(rand.Reader, prv, sum256[:])

	if err != nil {
		return ""
	}
	pubkey := strings.ToLower(hex.EncodeToString(FromECDSAPub(&prv.PublicKey)))

	addr := PubkeyToAddress(prv.PublicKey)

	rs := reverse(r.Bytes())

	ys := reverse(s.Bytes())

	return fmt.Sprintf("%s%s%s%s", pubkey, strings.ToLower(addr.Hex()[2:]), hex.EncodeToString(rs)[0:64], hex.EncodeToString(ys)[0:64])

}

//翻转byte
func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
