/**
* @Author:wjp
* @Date:2022/7/29 10:11
 */
package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

//生成秘钥对
func TestGenerateKey(t *testing.T) {
	key, _ := GenerateKey(rand.Reader, DefaultCurve, nil)

	k := key.ExportECDSA()
	//生成公钥字符串
	pubBytes := hex.EncodeToString(FromECDSAPub(&k.PublicKey))
	//生成私钥字符串
	pri := hex.EncodeToString(FromECDSA(k))
	//生成地址
	address := PubkeyToAddress(k.PublicKey).Hex()

	fmt.Println(fmt.Sprintf("私钥:%s\n公钥:%s\n地址:%s\n", pri, pubBytes, strings.ToLower(address)))
}

//生成地址
func TestCreateAddress(t *testing.T) {
	//传入私钥字符串
	ecdsas, err := HexToECDSA("b046347a995c3131c99fc5cf9e29ee4f7721e9b5ff06397df4eab597d08a9ef1")
	if err != nil {
		return
	}
	address := PubkeyToAddress(ecdsas.PublicKey)
	fmt.Println(strings.ToLower(address.String()))
}

//生成秘钥及地址并进行签名
func TestNewGenerateKey(T *testing.T) {

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return
	}
	//生成公钥字符串
	pubBytes := hex.EncodeToString(FromECDSAPub(&key.PublicKey))

	//生成私钥字符串
	pri := hex.EncodeToString(FromECDSA(key))

	//生成地址
	address := PubkeyToAddress(key.PublicKey).Hex()

	fmt.Println(fmt.Sprintf("私钥:%s\n公钥:%s\n地址:%s\n", pri, pubBytes, strings.ToLower(address)))
	//sign := Sign("helloworld", key)
	//fmt.Println("生成的数据签名为:", sign)
}
