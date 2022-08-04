/**
* @Author:wjp
* @Date:2022/7/27 16:50
 */
package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

const PublicKeyLength = 64

// FromECDSA exports a private key into a binary dump.
func FromECDSA(priv *ecdsa.PrivateKey) []byte {
	if priv == nil {
		return nil
	}
	keybytes := math.PaddedBigBytes(priv.D, priv.Params().BitSize/8)

	reverseString(keybytes)

	return keybytes
}

//公钥生成地址
func PubkeyToAddress(p ecdsa.PublicKey) common.Address {
	pubBytes := FromECDSAPub(&p)
	digest := sha256.Sum256(pubBytes)
	return common.BytesToAddress(digest[:20])
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	pubBytes := make([]byte, PublicKeyLength)
	pubXBytes := pub.X.Bytes()
	pubYBytes := pub.Y.Bytes()
	reverseString(pubXBytes)
	reverseString(pubYBytes)
	copy(pubBytes[:], pubXBytes)
	copy(pubBytes[PublicKeyLength/2:], pubYBytes)
	return pubBytes
}

func reverseString(s []byte) {

	if len(s) == 0 {
		return
	}
	reverseString(s[1:])
	c := s[0]

	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s[len(s)-1] = c

}

func toECDSA(d []byte) (*ecdsa.PrivateKey, error) {
	var e ecdsa.PrivateKey
	reverseString(d)
	e.D = new(big.Int).SetBytes(d)
	e.PublicKey.Curve = elliptic.P256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return &e, nil
}

// ToECDSA creates a private key with the given D value.
func ToECDSA(d []byte) (*ecdsa.PrivateKey, error) {
	return toECDSA(d)
}

//私钥字符串转私钥
func HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error) {
	//key := fmt.Sprintf("%s%s", hexkey[len(hexkey)-2:], hexkey[:len(hexkey)-2])

	b, err := hex.DecodeString(hexkey)
	if byteErr, ok := err.(hex.InvalidByteError); ok {
		return nil, fmt.Errorf("invalid hex character %q in private key", byte(byteErr))
	} else if err != nil {
		return nil, errors.New("invalid hex data for private key")
	}
	return ToECDSA(b)
}
